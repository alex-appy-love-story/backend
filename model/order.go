package model

type OrderInfo struct {
	UserID  uint `json:"user_id"`
	TokenID uint `json:"token_id"`
	Amount  uint `json:"amount"`
}
