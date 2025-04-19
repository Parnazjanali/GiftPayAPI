package models

type GIftCard struct {
	ID       string  `gorm:"primaryKey" json:"giftcardId" `
	UserId   string  `json:"userId"`
	WalletId string  `json:"walletId"`
	Amount   float64 `json:"amount"`
	IsUsed   bool    `json:"isUsed"`
}

type LinkGiftCardRequest struct {
	GiftCardID string `json:"giftcardId"`
}

type LinkGiftCardResponse struct {
	Status     string `json:"status"`
	Message    string `json:"message"`
	WalletId   string `json:"walletId"`
	GiftCardId string `json:"giftcardId"`
	Amount     string `json:"amount"`
}
type Wallet struct {
	ID string `json:"walletId"`
}

type ErrorResponse struct {
	Error struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:error`
}
