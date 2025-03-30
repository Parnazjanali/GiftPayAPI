package apiserver

import (
	"GiftCardManager/internal/api/handler"
	"GiftCardManager/internal/api/middlewares"
	sqlitedb "GiftCardManager/internal/repository/db/sqlite"

	"github.com/gofiber/fiber/v2"
)

func Run() {
	sqlitedb.InitDb()

	app := fiber.New()

	app.Post("/api/gift-cards/link", middlewares.JwtAuthMiddleware, handler.LinkGiftCardId)
	app.Listen(":8080")

}
