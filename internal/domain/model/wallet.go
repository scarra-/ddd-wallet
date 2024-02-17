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
	IncrementalId int    `gorm:"column:incremental_id;primarykey"`
	Id            string `gorm:"column:id"`
	OwnerId       string `gorm:"column:owner_id"`

	Balance  int    `gorm:"column:balance;not null"`
	Currency string `gorm:"column:currency;not null"`

	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index"`
}

func NewWallet(ownerId, currency string) *Wallet {
	return &Wallet{
		Id:       "wal-" + utils.RandomKey(28),
		OwnerId:  ownerId,
		Balance:  0,
		Currency: currency,
	}
}

func (w *Wallet) Fund(txChecker *TxChecker, amount int, source *FundSource) (*Transaction, error) {
	if 0 >= amount {
		return nil, InvalidAmount
	}

	if txChecker.TxOfOriginIdExists(source.OriginId) {
		return nil, OriginIdAlreadyUsed
	}

	w.Balance += amount
	w.UpdatedAt = time.Now()

	tx := newFundTransaction(w, amount, source)

	return tx, nil
}

func (w *Wallet) Spend(txChecker *TxChecker, amount int, source *SpendSource) (*Transaction, error) {
	if 0 >= amount {
		return nil, InvalidAmount
	}

	if 0 > (w.Balance - amount) {
		return nil, InsufficientFunds
	}

	if txChecker.TxOfOriginIdExists(source.OriginId) {
		return nil, OriginIdAlreadyUsed
	}

	w.Balance -= amount
	w.UpdatedAt = time.Now()

	tx := newSpendTransaction(w, amount, source)

	return tx, nil
}
