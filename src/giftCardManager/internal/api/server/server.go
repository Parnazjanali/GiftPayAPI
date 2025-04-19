package apiserver

import (
	"GiftCardManager/internal/api/handler"
	sqlitedb "GiftCardManager/internal/repository/db/sqlite"

	"github.com/gofiber/fiber/v2"
)

func Run() {
	sqlitedb.InitDb()

	app := fiber.New()

	app.Post("/api/gift-cards/link", handler.LinkGiftCardId)
	app.Listen("0.0.0.0:8080")

}
