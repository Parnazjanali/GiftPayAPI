package server

import (
	"GiftPayAPI/src/walletManager/internal/api/handler"
	"GiftPayAPI/src/walletManager/internal/api/middleware"
	sqliteDb "GiftPayAPI/src/walletManager/internal/repository/Db/sqlite"

	"github.com/gofiber/fiber/v2"
)

func Run() {

	sqliteDb.InitDb()

	app := fiber.New()

	app.Get("/wallet/{userId}", middleware.JWTAuthMiddleWare, handler.GetWalletId)
	app.Post("/wallet/{walletId}/set-balance", middleware.JWTAuthMiddleWare, handler.SetWalletBalance)

	app.Listen("0.0.0.0:8081")

}
