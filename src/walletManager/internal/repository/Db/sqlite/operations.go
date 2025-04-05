package sqliteDb

import (
	"GiftPayAPI/src/walletManager/internal/models"
	"errors"

	"gorm.io/gorm"
)

func UpdateWalletBalance(walletId string, newBalance float64) error {
	var wallet models.Wallet

	if err := DB.First(&wallet, "id=?", walletId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return gorm.ErrRecordNotFound
		}
		return err
	}
	wallet.Balance = newBalance
	if err := DB.Save(&wallet).Error; err != nil {
		return err
	}
	return nil
}
