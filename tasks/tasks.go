package tasks

import (
	"context"
	"fmt"

	"github.com/hibiken/asynq"
)

func HandlePerformStepTask(ctx context.Context, t *asynq.Task) error {
	//appCtx := GetContext(ctx)

	fmt.Println("Got response")
	return nil
}
