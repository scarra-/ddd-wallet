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
	incrementalId int    `gorm:"column:incremental_id;primarykey"`
	id            string `gorm:"column:id"`
	walletId      string `gorm:"column:wallet_id"`
	amount        int    `gorm:"column:amount;not null"`
	txType        string `gorm:"column:type;not null"`
	originId      string `gorm:"column:origin_id"`

	createdAt time.Time      `gorm:"column:created_at"`
	deletedAt gorm.DeletedAt `gorm:"column:deleted_at;index"`
}

func newFundTransaction(wallet *Wallet, amount int, source *FundSource) *Transaction {
	return &Transaction{
		id:       "tx-" + utils.RandomKey(29),
		walletId: wallet.Id(),
		amount:   amount,
		txType:   source.Type,
		originId: source.OriginId,
	}
}

func newSpendTransaction(wallet *Wallet, amount int, source *SpendSource) *Transaction {
	return &Transaction{
		id:       "tx-" + utils.RandomKey(29),
		walletId: wallet.Id(),
		amount:   amount * (-1), // Tx will be negative.
		txType:   source.Type,
		originId: source.OriginId,
	}
}

func (t *Transaction) Id() string {
	return t.id
}

func (t *Transaction) WalletId() string {
	return t.walletId
}

func (t *Transaction) Amount() int {
	return t.amount
}

func (t *Transaction) Type() string {
	return t.txType
}

func (t *Transaction) OriginId() string {
	return t.originId
}

func (t *Transaction) CreatedAt() time.Time {
	return t.createdAt
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
