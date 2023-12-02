package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/alex-appy-love-story/backend/model"
	"github.com/alex-appy-love-story/backend/tasks"
	"github.com/hibiken/asynq"
)

type OrderInfo struct {
    Username string `json:"username"`
	TokenID uint `json:"token_id"`
	Amount  uint `json:"amount"`
}

func Create(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Got request")
	ctx := tasks.GetContext(r.Context())
    orderInfo := &model.OrderInfo{}

    if err := json.NewDecoder(r.Body).Decode(&orderInfo); err != nil {
        return
    }

    task, err := tasks.NewStartOrderTask(*orderInfo)
    if err != nil {
        log.Fatalf("could not create task: %v", err)
    }
    fmt.Println("ENQUEING")
    info, err := ctx.AsynqClient.Enqueue(task, asynq.Queue("order"), asynq.MaxRetry(0))
    if err != nil {
        log.Fatalf("could not enqueue task: %v", err)
    }
    log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)   

}
