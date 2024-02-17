package dtos

import (
	"time"

	domain "github.com/aadejanovs/wallet/internal/domain/model"
)

type CreateWalletRequestDto struct {
	OwnerId  string `json:"owner_id" validate:"required"`
	Currency string `json:"currency" validate:"required"`
}

type FundWalletRequestDto struct {
	WalletId string
	Amount   int    `json:"amount" validate:"required"`
	Type     string `json:"type" validate:"required"`
	OriginId string `json:"origin_id" validate:"required"`
}

type SpendWalletFundsRequestDto struct {
	WalletId string
	Amount   int    `json:"amount" validate:"required"`
	Type     string `json:"type" validate:"required"`
	OriginId string `json:"origin_id" validate:"required"`
}

type WalletResponseDto struct {
	Id        string    `json:"id"`
	OwnerId   string    `json:"owner_id"`
	Balance   int       `json:"balance"`
	Currency  string    `json:"currency"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

func NewWalletResponseDto(w *domain.Wallet) *WalletResponseDto {
	return &WalletResponseDto{
		Id:        w.Id,
		OwnerId:   w.OwnerId,
		Balance:   w.Balance,
		Currency:  w.Currency,
		UpdatedAt: w.UpdatedAt,
		CreatedAt: w.CreatedAt,
	}
}
