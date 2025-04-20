package models

type Wallet struct {
	ID      string  `json:"walletId"`
	UserID  string  `json:"userId"`
	Balance float64 `json:"balance"`
}

type Transaction struct {
	ID              uint    `gorm:"primaryKey"` 
	WalletID        string  `gorm:"not null"`
	Amount          float64 `gorm:"not null"`
	TransactionType string  `gorm:"not null"`
	Timestamp       int64
	Wallet          Wallet `gorm:"foreignKey:WalletID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type SetBalanceRequest struct {
	Balance float64 `json:"balance"`
}

type ErrorResponse struct {
	Error struct {
		Code    string `json: "code"`
		Message string `json: "message"`
	} `json: "error"`
}
type Config struct {
	Port       string
	JwtSeceret string
}
