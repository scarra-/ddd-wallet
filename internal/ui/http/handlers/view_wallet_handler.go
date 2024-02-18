package handlers

import (
	"errors"

	"github.com/aadejanovs/wallet/internal/application/services/wallet"
	domain "github.com/aadejanovs/wallet/internal/domain/model"
	httpErrors "github.com/aadejanovs/wallet/internal/infrastructure/errors"
	repo "github.com/aadejanovs/wallet/internal/infrastructure/repositories/wallet"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func GetWallet(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	service := wallet.NewGetWalletService(repo.NewWalletRepository(db))

	response, err := service.Get(c.Params("id"))

	if err != nil {
		if errors.Is(err, domain.WalletNotFound) {
			return c.Status(404).JSON(httpErrors.NewCustomErrorResponse(404, err.Error()))
		}

		logger := c.Locals("logger").(*zap.SugaredLogger)
		logger.Errorw("error_while_viewing_wallet",
			"message", err,
			"wallet_id", c.Params("id"),
		)

		return c.Status(503).JSON(httpErrors.NewErrorResponse(503))
	}

	return c.JSON(response)
}
