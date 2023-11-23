package tasks

import (
	"context"

	"github.com/hibiken/asynq"
	"gorm.io/gorm"
)

func getClient(ctx context.Context) (gorm_client *gorm.DB, asynq_client *asynq.Client) {

	gorm_client = nil
	asynq_client = nil

	if val := ctx.Value("asynq_client"); val != nil {
		asynq_client = val.(*asynq.Client)
	}

	if val := ctx.Value("db_client"); val != nil {
		gorm_client = val.(*gorm.DB)
	}

	return
}
