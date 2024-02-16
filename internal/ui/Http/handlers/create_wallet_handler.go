package handlers

import (
	"github.com/aadejanovs/wallet/internal/application/services/wallet"
	"github.com/aadejanovs/wallet/internal/application/services/wallet/dtos"
	httpErrors "github.com/aadejanovs/wallet/internal/infrastructure/errors"
	repo "github.com/aadejanovs/wallet/internal/infrastructure/repositories/wallet"
	validation "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func CreateWallet(c *fiber.Ctx) error {
	dto := &dtos.CreateWalletRequestDto{}

	if err := c.BodyParser(dto); err != nil {
		return c.Status(400).JSON(httpErrors.NewErrorResponse(400))
	}

	if err := validation.New().Struct(dto); err != nil {
		vErr, ok := err.(validation.ValidationErrors)
		if !ok {
			return err
		}

		return c.Status(400).JSON(httpErrors.NewValidationErrorResponse(vErr))
	}

	service := wallet.NewCreateWalletService(
		repo.NewWalletRepository(c.Locals("db").(*gorm.DB)),
	)

	responseDto, err := service.Create(dto)
	if err != nil {
		logger := c.Locals("logger").(*zap.SugaredLogger)
		logger.Errorw("error_while_creating_wallet", "message", err)

		return c.Status(503).JSON(httpErrors.NewErrorResponse(503))
	}

	return c.Status(201).JSON(responseDto)
}
