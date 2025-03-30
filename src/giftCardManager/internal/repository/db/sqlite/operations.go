package sqlitedb

import "GiftCardManager/internal/models"

func GetGiftCardId(giftCardId string) (*models.GIftCard, error) {
	var giftCard models.GIftCard
	result := Db.Where("id=?", giftCardId).First(&giftCard)
	if result.Error != nil {
		return nil, result.Error
	}
	return &giftCard, nil
}

func LinkGIftCardToWallet(giftCard *models.GIftCard, walletId string) error {
	giftCard.WalletId = &walletId
	giftCard.IsUsed = true

	result := Db.Save(&giftCard)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
