package tasks

import (
	"encoding/json"

	"github.com/alex-appy-love-story/backend/model"
	"github.com/hibiken/asynq"
)

const (
	TypeStartOrder = "task:perform"
)

func NewStartOrderTask(orderInfo model.OrderInfo) (*asynq.Task, error) {
	payload, err := json.Marshal(orderInfo)
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeStartOrder, payload), nil
}
