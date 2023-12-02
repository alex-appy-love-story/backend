package app

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/hibiken/asynq"
)

type TaskHandler struct {
	app *App
}

type Task func(ctx context.Context, t *asynq.Task) error

// func (th *TaskHandler) asynqContextMiddleware(h asynq.Handler) asynq.Handler {
// 	return asynq.HandlerFunc(func(ctx context.Context, t *asynq.Task) error {
// 		ctx = context.WithValue(ctx, "asynq_client", th.app.AsynqClient)
// 		ctx = context.WithValue(ctx, "db_client", th.app.DBClient)
// 		return h.ProcessTask(ctx, t)
// 	})
// }

func loggingMiddleware(h asynq.Handler) asynq.Handler {
	return asynq.HandlerFunc(func(ctx context.Context, t *asynq.Task) error {
		start := time.Now()
		log.Printf("Start processing %q", t.Type())
		err := h.ProcessTask(ctx, t)
		if err != nil {
			return err
		}
		log.Printf("Finished processing %q: Elapsed Time = %v", t.Type(), time.Since(start))
		return nil
	})
}

 func (th *TaskHandler) routerContextMiddleware(next http.Handler) http.Handler {
 	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
 		ctx := r.Context()
 		ctx = context.WithValue(ctx, "asynq_client", th.app.AsynqClient)
 		ctx = context.WithValue(ctx, "db_client", th.app.DBClient)
 		next.ServeHTTP(rw, r.WithContext(ctx))
 	})
 }
