package domain

import (
	"errors"
	"time"

	"github.com/aadejanovs/wallet/internal/infrastructure/utils"
	"gorm.io/gorm"
)

var (
	TxNotFound = errors.New("tx not found")
)

type Transaction struct {
	IncrementalId int    `gorm:"column:incremental_id;primarykey"`
	Id            string `gorm:"column:id"`
	WalletId      string `gorm:"column:wallet_id"`
	Amount        int    `gorm:"column:amount;not null"`
	TxType        string `gorm:"column:type;not null"`
	OriginId      string `gorm:"column:origin_id"`

	CreatedAt time.Time      `gorm:"column:created_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index"`
}

func newFundTransaction(wallet *Wallet, amount int, source *FundSource) *Transaction {
	return &Transaction{
		Id:       "tx-" + utils.RandomKey(29),
		WalletId: wallet.Id,
		Amount:   amount,
		TxType:   source.Type,
		OriginId: source.OriginId,
	}
}

func newSpendTransaction(wallet *Wallet, amount int, source *SpendSource) *Transaction {
	return &Transaction{
		Id:       "tx-" + utils.RandomKey(29),
		WalletId: wallet.Id,
		Amount:   amount * (-1), // Tx will be negative.
		TxType:   source.Type,
		OriginId: source.OriginId,
	}
}

type TxRepo interface {
	OfOriginId(id string) (*Transaction, error)
}

// FundSource and SpendSource look same but with time as service matures
// and functionality grows - they will change and have different fields.
type FundSource struct {
	Type     string
	OriginId string
}

func NewFundSource(txType, originId string) *FundSource {
	return &FundSource{
		Type:     txType,
		OriginId: originId,
	}
}

type SpendSource struct {
	Type     string
	OriginId string
}

func NewSpendSource(txType, originId string) *SpendSource {
	return &SpendSource{
		Type:     txType,
		OriginId: originId,
	}
}

// TxChecker acts as a domain service which could be injected
// into domain model to check if transaction with specified
// origin id already exists. (To avoid double spend).
type TxChecker struct {
	txRepo TxRepo
}

func NewTxChecker(txRepo TxRepo) *TxChecker {
	return &TxChecker{
		txRepo: txRepo,
	}
}

func (t *TxChecker) TxOfOriginIdExists(originId string) bool {
	_, err := t.txRepo.OfOriginId(originId)

	if err != nil {
		return false
	}

	return true
}
