package server

import (
	"apiGateway/config"
	"apiGateway/internal/api/handler"
	"apiGateway/internal/api/middleware"
	sqliteDB "apiGateway/repository/db/sqlite"

	"github.com/gofiber/fiber/v2"
)

func Run() {

	config.LoadConfig()

	sqliteDB.InitDb()

	app := fiber.New()

	app.Post("/api/Login", handler.LoginUser)
	app.Post("/api/active/:giftCardId", middleware.JWTAuthMiddleware, handler.ActivateGiftCard)

	app.Listen("0.0.0.0:8082")

}
