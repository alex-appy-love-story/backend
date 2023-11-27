package tasks

import (
	"context"

	"github.com/hibiken/asynq"
	"gorm.io/gorm"
)

type AppContext struct {
	GormClient     *gorm.DB
	AsynqClient    *asynq.Client
	AsynqInspector *asynq.Inspector
	NextQueue      string
	ServerQueue    string
}


func GetContext(ctx context.Context) (appCtx AppContext) {

    appCtx.GormClient = nil
    appCtx.AsynqClient = nil

	if val := ctx.Value("asynq_client"); val != nil {
		appCtx.AsynqClient = val.(*asynq.Client)
	}

	if val := ctx.Value("db_client"); val != nil {
		appCtx.GormClient = val.(*gorm.DB)
	}

	if val := ctx.Value("next_queue"); val != nil {
		appCtx.NextQueue = val.(string)
	}

	if val := ctx.Value("server_queue"); val != nil {
		appCtx.ServerQueue = val.(string)
    }

	return
}
