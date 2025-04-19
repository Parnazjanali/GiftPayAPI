package handler

import (
	"GiftCardManager/internal/models"
	sqlitedb "GiftCardManager/internal/repository/db/sqlite"
	"GiftCardManager/internal/service"
	"log"

	"github.com/gofiber/fiber/v2"
)

func LinkGiftCardId(c *fiber.Ctx) error {
	var req models.LinkGiftCardRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    "400",
			"message": "Invalid request body",
		})
	}

	userId := c.Get("userId")
	if userId == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"code":    "401",
			"message": "Missing user ID",
		})
	}

	walletId, err := service.GetWalletId(userId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"code":    "404",
			"message": "Wallet not found",
		})
	}

	giftCard, err := sqlitedb.GetGiftCardByID(sqlitedb.Db, req.GiftCardID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"code":    "404",
			"message": "No gift card found with the provided ID",
		})
	}

	if giftCard.IsUsed {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"code":    "409",
			"message": "This gift card has already been linked to a wallet",
		})
	}

	err = sqlitedb.LinkGiftCardToWallet(sqlitedb.Db, giftCard, walletId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    "500",
			"message": "Internal server error while linking gift card",
		})
	}

	log.Println("🎉 Gift card linked successfully:", giftCard.ID, "->", walletId)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":     "success",
		"message":    "Gift card linked successfully, check your Wallet manager",
		"walletId":   walletId,
		"giftCardId": giftCard.ID,
		"amount":     giftCard.Amount,
	})
}
