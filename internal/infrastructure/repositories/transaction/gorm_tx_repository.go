package transaction

import (
	domain "github.com/aadejanovs/wallet/internal/domain/model"
	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (repo *TransactionRepository) OfId(id string) (*domain.Transaction, error) {
	var tx domain.Transaction
	err := repo.db.Where("id = ?", id).First(&tx).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, domain.TxNotFound
		}

		return nil, err
	}

	return &tx, nil
}

func (repo *TransactionRepository) OfOriginId(id string) (*domain.Transaction, error) {
	var tx domain.Transaction
	err := repo.db.Where("origin_id = ?", id).First(&tx).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, domain.TxNotFound
		}

		return nil, err
	}

	return &tx, nil
}
