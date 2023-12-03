package model

type OrderInfo struct {
	Username string `json:"username"`
	TokenID  uint   `json:"token_id"`
	Amount   uint   `json:"amount"`
}
