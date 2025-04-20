package handler

import (
	"fmt"
	"walletManager/internal/models"
	sqliteDb "walletManager/internal/repository/Db/sqlite"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetWalletId(c *fiber.Ctx) error {
	userId := c.Params("userId")
	if userId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error: struct {
				Code    string "json: \"code\""
				Message string "json: \"message\""
			}{
				Code:    "400",
				Message: "userId is required in the URL path",
			},
		})
	}
	var wallet models.Wallet
	if err := sqliteDb.DB.First(&wallet, "user_id = ?", userId).Error; err != nil {
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
func GetWalletBalance(c *fiber.Ctx) error {
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

	var wallet models.Wallet
	if err := sqliteDb.DB.First(&wallet, "id = ?", walletId).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(models.ErrorResponse{
			Error: struct {
				Code    string "json: \"code\""
				Message string "json: \"message\""
			}{
				Code:    "404",
				Message: "Wallet not found",
			},
		})
	}
	if wallet.Balance <= 0 {
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
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"balance": wallet.Balance,
	})
}
func GetWalletTransactions(c *fiber.Ctx) error {

	walletID := c.Params("walletId")
	fmt.Println("Wallet ID:", walletID)
	if walletID == "" {
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

	var transactions []models.Transaction
	result := sqliteDb.DB.Where("wallet_id = ?", walletID).Order("timestamp DESC").Limit(10).Preload("Wallet").Find(&transactions)
	if result.Error != nil {
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

	return c.Status(fiber.StatusOK).JSON(transactions)
}
