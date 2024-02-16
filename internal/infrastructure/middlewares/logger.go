package middlewares

import (
	customErr "github.com/aadejanovs/wallet/internal/infrastructure/errors"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func setupLogger() (*zap.SugaredLogger, error) {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	logger, err := config.Build()
	if err != nil {
		return nil, err
	}

	return logger.Sugar(), nil
}

func LoggingMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		logger, err := setupLogger()
		if err != nil {
			return c.Status(500).JSON(customErr.NewErrorResponse((500)))
		}
		defer logger.Sync()

		c.Locals("logger", logger)

		return c.Next()
	}
}
