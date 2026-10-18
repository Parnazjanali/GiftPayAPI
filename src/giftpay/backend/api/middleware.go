package api

import (
	"log/slog"

	"giftpay/backend/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// RequestLogger middleware adds a request ID to the context and logs requests
func RequestLogger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		reqID := c.Get("X-Request-Id")
		if reqID == "" {
			reqID = uuid.New().String()
		}
		c.Set("X-Request-Id", reqID)

		// Create a logger instance with request ID
		log := slog.With("request_id", reqID)

		// Set logger in user context for downstream use
		ctx := logger.WithLogger(c.UserContext(), log)
		c.SetUserContext(ctx)

		log.Info("Incoming request",
			"method", c.Method(),
			"path", c.Path(),
			"ip", c.IP(),
		)

		err := c.Next()

		log.Info("Request completed",
			"status", c.Response().StatusCode(),
			"method", c.Method(),
			"path", c.Path(),
		)

		return err
	}
}

// ErrorHandler is a central error handling middleware
func ErrorHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := c.Next()
		if err != nil {
			log := logger.GetLogger(c.UserContext())
			log.Error("Unhandled error in request", "error", err)

			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}

			return c.Status(code).JSON(fiber.Map{
				"error": fiber.Map{
					"code":    code,
					"message": "Internal Server Error",
				},
			})
		}
		return nil
	}
}
