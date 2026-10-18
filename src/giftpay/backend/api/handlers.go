package api

import (
	"errors"

	"giftpay/backend/models"
	"giftpay/backend/repository"
	"giftpay/backend/service"

	"github.com/gofiber/fiber/v2"
)

type GiftPayHandler struct {
	svc service.GiftPayService
}

func NewGiftPayHandler(svc service.GiftPayService) *GiftPayHandler {
	return &GiftPayHandler{svc: svc}
}

func (h *GiftPayHandler) ApplyDiscountCode(c *fiber.Ctx) error {
	var req models.ApplyGiftCodeRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error: struct {
				Code    string `json:"code"`
				Message string `json:"message"`
			}{Code: "400", Message: "Invalid request body"},
		})
	}

	if req.Code == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error: struct {
				Code    string `json:"code"`
				Message string `json:"message"`
			}{Code: "400", Message: "Code is required"},
		})
	}

	// Assuming user ID is passed in context by auth middleware, or we take it from body
	userID := req.UserID
	if userID == "" {
		userID = c.Get("userId") // fallback to header if empty
	}
	if userID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ErrorResponse{
			Error: struct {
				Code    string `json:"code"`
				Message string `json:"message"`
			}{Code: "401", Message: "Missing user ID"},
		})
	}

	ctx := c.UserContext()
	resp, err := h.svc.ApplyDiscountCode(ctx, req.Code, userID)
	if err != nil {
		status := fiber.StatusInternalServerError
		code := "500"
		msg := "Internal server error"

		switch {
		case errors.Is(err, repository.ErrCodeNotFound):
			status = fiber.StatusNotFound
			code = "404"
			msg = "Invalid or non-existent discount code"
		case errors.Is(err, repository.ErrCodeInactive):
			status = fiber.StatusBadRequest
			code = "400"
			msg = "Discount code is no longer active"
		case errors.Is(err, repository.ErrCodeExpired):
			status = fiber.StatusBadRequest
			code = "400"
			msg = "Discount code has expired"
		case errors.Is(err, repository.ErrCodeUsageLimit):
			status = fiber.StatusBadRequest
			code = "400"
			msg = "Discount code usage limit reached"
		case errors.Is(err, repository.ErrCodeAlreadyUsed):
			status = fiber.StatusConflict
			code = "409"
			msg = "You have already used this discount code"
		}

		return c.Status(status).JSON(models.ErrorResponse{
			Error: struct {
				Code    string `json:"code"`
				Message string `json:"message"`
			}{Code: code, Message: msg},
		})
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}

func RegisterRoutes(app *fiber.App, handler *GiftPayHandler) {
	api := app.Group("/api/v1/giftpay")
	api.Post("/apply", handler.ApplyDiscountCode)
}
