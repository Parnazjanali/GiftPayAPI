package server

import (
	"apiGateway/internal/api/handler"
	"apiGateway/internal/middleware"
	sqliteDB "apiGateway/repository/db/sqlite"

	"github.com/gofiber/fiber/v2"
)

func Run() {

	sqliteDB.InitDb()

	app := fiber.New()

	app.Post("/api/Login", handler.LoginUser)
	app.Post("/api/Verify-Token", middleware.JWTAuthMiddleware, handler.ValidateToken)

	app.Listen("0.0.0.0:8082")

}
