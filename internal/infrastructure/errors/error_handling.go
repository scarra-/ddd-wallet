package errors

import (
	stdErr "errors"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func HandleHTTPErrors(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	// Retrieve the custom status code if it's a *fiber.Error
	var e *fiber.Error
	if stdErr.As(err, &e) {
		code = e.Code
	}

	if code == fiber.StatusInternalServerError {
		logger := ctx.Locals("logger").(*zap.SugaredLogger)
		logger.Errorw("internal_server_error",
			"message", err,
		)
	}

	ctx.Status(404).JSON(NewErrorResponse((code)))
	return nil
}
