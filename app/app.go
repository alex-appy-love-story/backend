package app

import (
	"context"
	"fmt"

	"github.com/alex-appy-love-story/worker-template/tasks"
	"github.com/hibiken/asynq"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type App struct {
	Config      Config
	AsynqClient *asynq.Client
	DBClient    *gorm.DB
}

func New(config Config) *App {
	app := &App{
		Config: config,
		AsynqClient: asynq.NewClient(asynq.RedisClientOpt{
			Addr: config.RedisAddress,
		}),
	}

	return app
}

func (a *App) connectDB(ctx context.Context) error {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/db?charset=utf8mb4&parseTime=True&loc=Local",
		a.Config.DatabaseInfo.User,
		a.Config.DatabaseInfo.Password,
		a.Config.DatabaseInfo.Address,
	)

	gorm, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	a.DBClient = gorm
	return nil
}

func (a *App) Start(ctx context.Context) error {

	server := asynq.NewServer(
		asynq.RedisClientOpt{Addr: a.Config.RedisAddress},
		asynq.Config{
			Concurrency: 5,
		},
	)

	fmt.Println("Successfully connected to redis!")

	if err := a.connectDB(ctx); err != nil {
		return err
	}

	fmt.Println("Successfully connected to db!")

	defer func() {
		if err := a.AsynqClient.Close(); err != nil {
			fmt.Println("Failed to close redis", err)
		}
	}()

	fmt.Println("Starting server...")

	ch := make(chan error, 1)

	mux := asynq.NewServeMux()

	taskHanlder := TaskHandler{app: a}
	mux.Use(loggingMiddleware)
	mux.Use(taskHanlder.ContextMiddleware)

	tasks.RegisterTopic(mux)

	go func() {
		err := server.Run(mux)
		if err != nil {
			ch <- fmt.Errorf("Failed to start server: %w", err)
		}
		close(ch)
	}()

	select {
	case err := <-ch:
		return err
	case <-ctx.Done():
		server.Shutdown()
		return nil
	}
}
