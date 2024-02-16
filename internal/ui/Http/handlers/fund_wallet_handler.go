package handlers

import (
	"errors"

	"github.com/aadejanovs/wallet/internal/application/services/wallet"
	"github.com/aadejanovs/wallet/internal/application/services/wallet/dtos"
	domain "github.com/aadejanovs/wallet/internal/domain/model"
	httpErrors "github.com/aadejanovs/wallet/internal/infrastructure/errors"
	txRepo "github.com/aadejanovs/wallet/internal/infrastructure/repositories/transaction"
	walletRepo "github.com/aadejanovs/wallet/internal/infrastructure/repositories/wallet"
	validation "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func FundWallet(c *fiber.Ctx) error {
	dto := &dtos.FundWalletRequestDto{}

	if err := c.BodyParser(dto); err != nil {
		return c.Status(400).JSON(httpErrors.NewErrorResponse(400))
	}

	dto.WalletId = c.Params("id")

	if err := validation.New().Struct(dto); err != nil {
		vErr, ok := err.(validation.ValidationErrors)
		if !ok {
			return err
		}

		return c.Status(400).JSON(httpErrors.NewValidationErrorResponse(vErr))
	}

	db := c.Locals("db").(*gorm.DB)

	service := wallet.NewFundWalletService(
		walletRepo.NewWalletRepository(db),
		txRepo.NewTransactionRepository(db),
	)

	txDto, err := service.Fund(dto)
	if err != nil {
		if errors.Is(err, domain.WalletNotFound) {
			return c.Status(404).JSON(httpErrors.NewCustomErrorResponse(404, err.Error()))
		}

		if errors.Is(err, domain.InvalidAmount) || errors.Is(err, domain.OriginIdAlreadyUsed) {
			return c.Status(400).JSON(httpErrors.NewCustomErrorResponse(400, err.Error()))
		}

		logger := c.Locals("logger").(*zap.SugaredLogger)
		logger.Errorw("error_while_funding_wallet",
			"message", err,
			"wallet_id", dto.WalletId,
			"origin_id", dto.OriginId,
			"amount", dto.Amount,
		)

		return c.Status(503).JSON(httpErrors.NewErrorResponse(503))
	}

	return c.Status(201).JSON(txDto)
}
