package infrastructure

import (
	"log"
	"os"

	"github.com/aadejanovs/wallet/database"
	"github.com/aadejanovs/wallet/internal/infrastructure/errors"
	"github.com/aadejanovs/wallet/internal/infrastructure/middlewares"
	"github.com/aadejanovs/wallet/internal/infrastructure/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/joho/godotenv"
)

func Setup() *fiber.App {
	err := godotenv.Load()
	if err != nil && !os.IsNotExist(err) {
		log.Fatalf("Error loading .env file: %v", err)
	}

	if err := database.Connect(); err != nil {
		log.Fatal("Mysql connection failed", err)
	}

	app := fiber.New(fiber.Config{
		ErrorHandler: errors.HandleHTTPErrors,
	})

	app.Use(healthcheck.New())
	app.Use(middlewares.LoggingMiddleware())
	app.Use(middlewares.DbMiddleware())

	routes.SetupRoutes(app)

	return app
}
