package middlewares

import (
	"github.com/aadejanovs/wallet/database"
	"github.com/gofiber/fiber/v2"
)

func DbMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Locals("db", database.DBConn)
		return c.Next()
	}
}
