package models

type Wallet struct {
	ID      string  `json:"walletId"`
	UserID  string  `json:"userId"`
	Balance float64 `json:"balance"`
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
