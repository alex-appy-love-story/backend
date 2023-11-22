package entity

import (
    "time"
)

type Order struct {
	OrderID     uint64      `json:"order_id"`
	UserID      uint64      `json:"user_id"`
    ItemID      uint64      `json:"item_id"`
    Quantity    uint        `json:"quantity"`
	CreatedAt   *time.Time  `json:"created_at"`
	ShippedAt   *time.Time  `json:"shipped_at"`
	CompletedAt *time.Time  `json:"completed_at"`
}
