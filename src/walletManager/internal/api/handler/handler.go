package handler

import (
	"walletManager/internal/models"
	sqliteDb "walletManager/internal/repository/Db/sqlite"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetWalletId(c *fiber.Ctx) error {
	userId, ok := c.Locals("userId").(string)
	if !ok || userId == "" {
		c.Status(fiber.StatusUnauthorized).JSON(models.ErrorResponse{
			Error: struct {
				Code    string "json: \"code\""
				Message string "json: \"message\""
			}{
				Code:    "401",
				Message: "Invalid or missing userId",
			},
		})
	}
	var wallet models.Wallet
	result := sqliteDb.DB.Where("user_id=?", userId).First(&wallet)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(models.ErrorResponse{
			Error: struct {
				Code    string "json: \"code\""
				Message string "json: \"message\""
			}{
				Code:    "404",
				Message: "No wallet associated with the provided user ID",
			},
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":   "success",
		"walletId": wallet.ID,
		"userId":   userId,
	})
}
func SetWalletBalance(c *fiber.Ctx) error {
	walletId := c.Params("walletId")
	if walletId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error: struct {
				Code    string "json: \"code\""
				Message string "json: \"message\""
			}{
				Code:    "400",
				Message: "walletId is required in the URL path",
			},
		})
	}

	var Body models.SetBalanceRequest

	if err := c.BodyParser(&Body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error: struct {
				Code    string "json: \"code\""
				Message string "json: \"message\""
			}{
				Code:    "400",
				Message: "Invalid request body",
			},
		})
	}
	if Body.Balance <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error: struct {
				Code    string "json: \"code\""
				Message string "json: \"message\""
			}{
				Code:    "400",
				Message: "Balance must be greater than 0",
			},
		})
	}

	err := sqliteDb.UpdateWalletBalance(walletId, Body.Balance)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(models.ErrorResponse{
				Error: struct {
					Code    string "json: \"code\""
					Message string "json: \"message\""
				}{
					Code:    "404",
					Message: "Wallet not found",
				},
			})
			return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
				Error: struct {
					Code    string "json: \"code\""
					Message string "json: \"message\""
				}{
					Code:    "500",
					Message: "Internal server error",
				},
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":     "success",
		"message":    "Wallet balance updated successfully",
		"newBalance": Body.Balance,
	})

}
