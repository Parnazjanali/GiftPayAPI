package models

type Wallet struct {
	ID      string  `json:"walletId"`
	UserID  string  `json:"userId"`
	Balance float64 `json:"balance"`
}

type Transaction struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	WalletID  string    `json:"wallet_id" gorm:"index"`
	Amount    float64   `json:"amount"`
	Timestamp int64     `json:"timestamp"` 
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
