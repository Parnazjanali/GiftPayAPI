package models

type GIftCard struct {
	ID       string  `gorm:"primaryKey" json:"giftcardId" `
	WalletId *string `json:"walletId"`
	Amount   float64 `json:"amount"`
	IsUsed   bool    `json:"isUsed"`
}
type ErrorResponse struct {
	Error struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:error`
}
