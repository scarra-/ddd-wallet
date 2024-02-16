package domain

import (
	"errors"
	"time"

	"github.com/aadejanovs/wallet/internal/infrastructure/utils"
	"gorm.io/gorm"
)

var (
	InvalidAmount       = errors.New("invalid amount")
	InsufficientFunds   = errors.New("insufficient funds")
	OriginIdAlreadyUsed = errors.New("origin id already used")
	WalletNotFound      = errors.New("wallet not found")
)

type Wallet struct {
	incrementalId int    `gorm:"column:incremental_id;primarykey"`
	id            string `gorm:"column:id"`
	ownerId       string `gorm:"column:owner_id"`

	balance  int    `gorm:"column:balance;not null"`
	currency string `gorm:"column:currency;not null"`

	createdAt time.Time      `gorm:"column:created_at"`
	updatedAt time.Time      `gorm:"column:updated_at"`
	deletedAt gorm.DeletedAt `gorm:"column:deleted_at;index"`
}

func NewWallet(ownerId, currency string) *Wallet {
	return &Wallet{
		id:       "wal-" + utils.RandomKey(28),
		ownerId:  ownerId,
		balance:  0,
		currency: currency,
	}
}

func (w *Wallet) Fund(txChecker *TxChecker, amount int, source *FundSource) (*Transaction, error) {
	if 0 >= amount {
		return nil, InvalidAmount
	}

	if txChecker.TxOfOriginIdExists(source.OriginId) {
		return nil, OriginIdAlreadyUsed
	}

	w.balance += amount
	w.updatedAt = time.Now()

	tx := newFundTransaction(w, amount, source)

	return tx, nil
}

func (w *Wallet) Spend(txChecker *TxChecker, amount int, source *SpendSource) (*Transaction, error) {
	if 0 >= amount {
		return nil, InvalidAmount
	}

	if 0 > (w.balance - amount) {
		return nil, InsufficientFunds
	}

	if txChecker.TxOfOriginIdExists(source.OriginId) {
		return nil, OriginIdAlreadyUsed
	}

	w.balance -= amount
	w.updatedAt = time.Now()

	tx := newSpendTransaction(w, amount, source)

	return tx, nil
}

func (w *Wallet) Id() string {
	return w.id
}

func (w *Wallet) OwnerId() string {
	return w.ownerId
}

func (w *Wallet) Balance() int {
	return w.balance
}

func (w *Wallet) Currency() string {
	return w.currency
}

func (w *Wallet) CreatedAt() time.Time {
	return w.createdAt
}

func (w *Wallet) UpdatedAt() time.Time {
	return w.updatedAt
}
