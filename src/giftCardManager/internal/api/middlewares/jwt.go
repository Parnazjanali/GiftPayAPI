package middlewares

import (
	"GiftCardManager/internal/models"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("your-secret-key") 

func JwtAuthMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ErrorResponse{
			Error: struct {
				Code    string `json:"code"`
				Message string `json:"message"`
			}{
				Code:    "unauthorized",
				Message: "Missing or invalid authorization header",
			},
		})
	}

	tokenParts := strings.Split(authHeader, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ErrorResponse{
			Error: struct {
				Code    string `json:"code"`
				Message string `json:"message"`
			}{
				Code:    "unauthorized",
				Message: "Invalid token format",
			},
		})
	}

	tokenStr := tokenParts[1] 

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid signing method")
		}
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ErrorResponse{
			Error: struct {
				Code    string `json:"code"`
				Message string `json:"message"`
			}{
				Code:    "unauthorized",
				Message: "Invalid or expired token",
			},
		})
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ErrorResponse{
			Error: struct {
				Code    string `json:"code"`
				Message string `json:"message"`
			}{
				Code:    "unauthorized",
				Message: "Invalid token claims",
			},
		})
	}

	if userID, exists := claims["user_id"].(string); exists {
		c.Locals("userId", userID)
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ErrorResponse{
			Error: struct {
				Code    string `json:"code"`
				Message string `json:"message"`
			}{
				Code:    "unauthorized",
				Message: "User ID not found in token",
			},
		})
	}

	return c.Next()
}
