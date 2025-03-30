package service

import (
	"errors"
)

func GetWalletId(userId string) (string, error) {
	if userId == "test_user" {
		return "test_wallet", nil

	}
	return "", errors.New("wallet not found")
}
