package sqlitedb

import (
	"GiftCardManager/internal/models"
	"fmt"

	"gorm.io/gorm"
)

func GetGiftCardByID(db *gorm.DB, giftCardId string) (*models.GIftCard, error) {
	var giftCard models.GIftCard
	result := db.Where("id = ?", giftCardId).First(&giftCard)
	if result.Error != nil {
		return nil, result.Error
	}
	return &giftCard, nil
}
func LinkGiftCardToWallet(db *gorm.DB, giftCard *models.GIftCard, walletId string) error {
	// بررسی اینکه کارت هدیه قبلاً استفاده شده است یا نه
	if giftCard.IsUsed {
		return fmt.Errorf("این کارت هدیه قبلاً استفاده شده است")
	}

	// به‌روزرسانی اطلاعات کارت هدیه
	giftCard.WalletId = walletId
	giftCard.IsUsed = true
	result := db.Save(&giftCard)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func GetWalletByUserId(db *gorm.DB, userId string) (*models.Wallet, error) {
	var wallet models.Wallet
	result := db.Where("user_id = ?", userId).First(&wallet)
	if result.Error != nil {
		return nil, result.Error
	}
	return &wallet, nil
}
