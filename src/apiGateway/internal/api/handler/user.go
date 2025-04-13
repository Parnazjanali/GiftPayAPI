package handler

import (
	"apiGateway/internal/model"
	Service "apiGateway/internal/service"
	"apiGateway/utils"

	"github.com/gofiber/fiber/v2"
)

func LoginUser(c *fiber.Ctx) error {

	var req model.User
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Error: struct {
				Code    string "json:\"code\""
				Message string "json:\"message\""
			}{
				Code:    "400",
				Message: "Username or Password not provided.",
			},
		})
	}
	user, err := Service.AuthenticateUser(req.Username, req.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(model.ErrorResponse{
			Error: struct {
				Code    string "json:\"code\""
				Message string "json:\"message\""
			}{
				Code:    "401",
				Message: "Incorrect username or password.",
			},
		})
	}
	token, err := utils.GenerateToken(*user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.ErrorResponse{
			Error: struct {
				Code    string "json:\"code\""
				Message string "json:\"message\""
			}{
				Code:    "500",
				Message: "Internal server error.",
			},
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token":      token,
		"expires_in": 3600,
	})
}
func ActivateGiftCard(c *fiber.Ctx) error {
	userId := c.Locals("userId")

	if userId == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(model.ErrorResponse{
			Error: struct {
				Code    string "json:\"code\""
				Message string "json:\"message\""
			}{
				Code:    "401",
				Message: "Unauthorized",
			},
		})
	}

	giftCardId := c.Params("giftCardId")
	if giftCardId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Error: struct {
				Code    string "json:\"code\""
				Message string "json:\"message\""
			}{
				Code:    "400",
				Message: "Gift card ID not provided.",
			},
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"userId":     userId,
		"giftCardId": giftCardId,
		"message":    "Gift Card will be activated soon",
	})
}
