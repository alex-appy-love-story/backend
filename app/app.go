package app

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/alex-appy-love-story/backend/tasks"
	"github.com/hibiken/asynq"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type App struct {
	Config      Config
	AsynqClient *asynq.Client
	DBClient    *gorm.DB
	router      http.Handler
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
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		a.Config.DatabaseConfig.User,
		a.Config.DatabaseConfig.Password,
		a.Config.DatabaseConfig.Address,
		a.Config.DatabaseConfig.DatabaseName,
	)

	gorm, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	a.DBClient = gorm
	return nil
}

func (a *App) getApp() *App {
	return a
}

func (a *App) Start(ctx context.Context) error {

	asynq_server := asynq.NewServer(
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

	a.router = loadRoutes(a)
	http_server := &http.Server{
		Addr:    ":3000",
		Handler: a.router,
	}

	fmt.Println("Server started!")

	mux := asynq.NewServeMux()

	taskHanlder := TaskHandler{app: a}
	mux.Use(loggingMiddleware)
	mux.Use(taskHanlder.asynqContextMiddleware)

	tasks.RegisterTopic(mux)

	asynq_ch := make(chan error, 1)
	http_ch := make(chan error, 1)

	go func() {
		err := http_server.ListenAndServe()
		if err != nil {
			http_ch <- fmt.Errorf("failed to start http server: %w", err)
		}
		close(http_ch)
	}()

	go func() {
		err := asynq_server.Run(mux)
		if err != nil {
			asynq_ch <- fmt.Errorf("Failed to start asynq server: %w", err)
		}
		close(asynq_ch)
	}()

	select {
	case err := <-asynq_ch:
		return err
	case err := <-http_ch:
		return err
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		asynq_server.Shutdown()
		return http_server.Shutdown(timeout)
	}
}
