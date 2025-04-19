package server

import (
	"walletManager/internal/api/handler"
	sqliteDb "walletManager/internal/repository/Db/sqlite"

	"github.com/gofiber/fiber/v2"
)

func Run() {

	sqliteDb.InitDb()

	app := fiber.New()

	app.Get("/wallet/:userId", handler.GetWalletId)
	app.Post("/wallet/:walletId/set-balance", handler.SetWalletBalance)
	app.Get("/wallet/:walletId/get-balance", handler.GetWalletBalance)
	app.Get("/wallet/:walletId/get-transactions", handler.GetWalletTransactions)

	app.Listen("0.0.0.0:8081")

}
