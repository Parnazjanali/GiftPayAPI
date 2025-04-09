package middleware

import (
	"strings"
	"walletManager/internal/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

var JwtSeceret = []byte("secret")

func JWTAuthMiddleWare(c *fiber.Ctx) error {

	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ErrorResponse{
			Error: struct {
				Code    string "json: \"code\""
				Message string "json: \"message\""
			}{
				Code:    "401",
				Message: "Missing or invalid authorization header",
			},
		})
	}

	tokenParts := strings.Split(authHeader, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ErrorResponse{
			Error: struct {
				Code    string "json: \"code\""
				Message string "json: \"message\""
			}{
				Code:    "401",
				Message: "Invalid token format",
			},
		})
	}
	tokenStr := tokenParts[1]

	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {

			return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid signing method")
		}
		return JwtSeceret, nil

	})
	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ErrorResponse{
			Error: struct {
				Code    string "json: \"code\""
				Message string "json: \"message\""
			}{
				Code:    "401",
				Message: "Invalid or Expired token",
			},
		})
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ErrorResponse{
			Error: struct {
				Code    string "json: \"code\""
				Message string "json: \"message\""
			}{
				Code:    "401",
				Message: "Invalid token claims",
			},
		})
	}

	userID, exists := claims["userId"].(string)
	if !exists {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ErrorResponse{
			Error: struct {
				Code    string "json: \"code\""
				Message string "json: \"message\""
			}{
				Code:    "401",
				Message: "User ID not found in token",
			},
		})
	}

	c.Locals("userId", userID)

	return c.Next()

}
