package handler

import (
	"apiGateway/internal/model"
	Service "apiGateway/internal/service"
	"apiGateway/utils"
	"strings"

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
func ValidateToken(c *fiber.Ctx) error {

	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(model.ErrorResponse{
			Error: struct {
				Code    string `json:"code"`
				Message string `json:"message"`
			}{
				Code:    "401",
				Message: "Authorization header is missing",
			},
		})
	}

	tokenParts := strings.Split(authHeader, "Bearer ")
	if len(tokenParts) != 2 || tokenParts[0] != "" {
		return c.Status(fiber.StatusUnauthorized).JSON(model.ErrorResponse{
			Error: struct {
				Code    string `json:"code"`
				Message string `json:"message"`
			}{
				Code:    "401",
				Message: "Invalid or expired token",
			},
		})
	}

	tokenString := tokenParts[1]

	claims, err := utils.VerifyToken(tokenString)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(model.ErrorResponse{
			Error: struct {
				Code    string `json:"code"`
				Message string `json:"message"`
			}{
				Code:    "401",
				Message: "Invalid or expired token",
			},
		})
	}

	userId, ok := claims["username"].(string)
	if !ok || userId == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(model.ErrorResponse{
			Error: struct {
				Code    string `json:"code"`
				Message string `json:"message"`
			}{
				Code:    "401",
				Message: "Invalid or expired token",
			},
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"user_id": userId,
		"status":  "valid",
	})
}
