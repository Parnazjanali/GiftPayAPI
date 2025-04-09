package middleware

import (
	"apiGateway/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func JWTAuthMiddleware(c *fiber.Ctx) error {

	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"Error": "Authorization header is missing",
		})
	}

	tokenParts := strings.Split(authHeader, "Bearer ")
	if len(tokenParts) != 2 || tokenParts[0] != "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"Error": "Invalid Authentication token",
		})
	}
	tokenString := tokenParts[1]

	claims, err := utils.VerifyToken(tokenString)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"Error": "Invalid Authentication token",
		})
	}
	c.Locals("userId", claims["username"])
	return c.Next()
}
