package handler

import (
	"GiftCardManager/internal/models"
	sqlitedb "GiftCardManager/internal/repository/db/sqlite"
	"GiftCardManager/internal/service"

	"github.com/gofiber/fiber/v2"
)

func LinkGiftCardId(c *fiber.Ctx) error {
	userId := c.Locals("userId").(string)

	if userId == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ErrorResponse{
			Error: struct {
				Code    string "json:\"code\""
				Message string "json:\"message\""
			}{
				Code:    "unauthorized",
				Message: "Invalid or expired token",
			},
		})
	}
	var req models.GIftCard
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error: struct {
				Code    string "json:\"code\""
				Message string "json:\"message\""
			}{
				Code:    "bad_request",
				Message: "Invalid request body",
			},
		})
	}

	walletID, err := service.GetWalletId(userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Error: struct {
				Code    string "json:\"code\""
				Message string "json:\"message\""
			}{
				Code:    "internal_server_error",
				Message: "Failed to retrieve wallet ID",
			},
		})
	}

	giftCard, err := sqlitedb.GetGiftCardId(req.ID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(models.ErrorResponse{
			Error: struct {
				Code    string "json:\"code\""
				Message string "json:\"message\""
			}{
				Code:    "not_found",
				Message: "Gift card not found",
			},
		})
	}
	if giftCard.IsUsed {
		return c.Status(fiber.StatusConflict).JSON(models.ErrorResponse{
			Error: struct {
				Code    string "json:\"code\""
				Message string "json:\"message\""
			}{
				Code:    "conflict",
				Message: "This gift card has already been linked to a wallet",
			},
		})
	}

	err = sqlitedb.LinkGIftCardToWallet(giftCard, walletID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Error: struct {
				Code    string "json:\"code\""
				Message string "json:\"message\""
			}{
				Code:    "internal_server_error",
				Message: "Failed to link gift card to wallet",
			},
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":     "success",
		"message":    "Gift card linked to wallet successfully",
		"walletId":   walletID,
		"giftCardId": giftCard.ID,
		"amount":     giftCard.Amount,
	})
}
