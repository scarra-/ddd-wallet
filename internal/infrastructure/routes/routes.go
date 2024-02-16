package routes

import (
	"github.com/aadejanovs/wallet/internal/ui/http/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/wallets/:id", handlers.GetWallet)
	app.Post("/wallets", handlers.CreateWallet)
	app.Post("/wallets/:id/fund", handlers.FundWallet)
	app.Post("/wallets/:id/spend", handlers.SpendWalletFunds)
}
