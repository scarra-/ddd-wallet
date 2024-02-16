package dtos

import (
	"time"

	domain "github.com/aadejanovs/wallet/internal/domain/model"
)

type TxResponseDto struct {
	Id        string    `json:"id"`
	WalletId  string    `json:"owner_id"`
	Amount    int       `json:"amount"`
	Type      string    `json:"type"`
	OriginId  string    `json:"origin_id"`
	CreatedAt time.Time `json:"created_at"`
}

func NewTxResponseDto(tx *domain.Transaction) *TxResponseDto {
	return &TxResponseDto{
		Id:        tx.Id(),
		WalletId:  tx.WalletId(),
		Amount:    tx.Amount(),
		Type:      tx.Type(),
		OriginId:  tx.OriginId(),
		CreatedAt: tx.CreatedAt(),
	}
}
