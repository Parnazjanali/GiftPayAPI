package sqliteDb

import (
	"walletManager/internal/models"
)

func GetWalletBalance(walletId string) (float64, error) {
	var wallet models.Wallet
	if err := DB.First(&wallet, "id = ?", walletId).Error; err != nil {
		return 0, err
	}
	return wallet.Balance, nil
}

func UpdateWalletBalance(walletId string, newBalance float64) error {
	var wallet models.Wallet
	if err := DB.First(&wallet, "id = ?", walletId).Error; err != nil {
		return err
	}
	wallet.Balance = newBalance
	return DB.Save(&wallet).Error
}
