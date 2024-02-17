package wallet

import (
	"errors"

	domain "github.com/aadejanovs/wallet/internal/domain/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type WalletRepository struct {
	db *gorm.DB
	tx *gorm.DB
}

func NewWalletRepository(db *gorm.DB) *WalletRepository {
	return &WalletRepository{db: db}
}

func (repo *WalletRepository) Save(wallet *domain.Wallet) error {
	repo.db.Save(wallet)
	return repo.db.Error
}

func (repo *WalletRepository) OfId(id string) (*domain.Wallet, error) {
	var w domain.Wallet
	if err := repo.db.Where("id = ?", id).First(&w).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.WalletNotFound
		}
		return nil, err
	}
	return &w, nil
}

// Returns wallet from DB.
// Uses pessimistic locking which prevents simultaneous updates from different processes.
func (repo *WalletRepository) OfIdForUpdate(id string) (*domain.Wallet, error) {
	var w domain.Wallet

	repo.tx = repo.db.Begin()
	if repo.tx.Error != nil {
		return nil, repo.tx.Error
	}

	err := repo.tx.
		Clauses(clause.Locking{Strength: "UPDATE"}).
		Model(&domain.Wallet{}).
		Where("id = ?", id).
		First(&w).Error

	if err != nil {
		repo.tx.Rollback()

		if err == gorm.ErrRecordNotFound {
			return nil, domain.WalletNotFound
		}

		return nil, err
	}

	return &w, nil
}

func (repo *WalletRepository) SaveWalletAndTx(wallet *domain.Wallet, tx *domain.Transaction) error {
	if repo.tx == nil {
		return errors.New("db tx not initiated")
	}

	if err := repo.tx.Save(wallet).Error; err != nil {
		repo.tx.Rollback()
		return err
	}

	if err := repo.tx.Create(tx).Error; err != nil {
		repo.tx.Rollback()
		return err
	}

	if err := repo.tx.Commit().Error; err != nil {
		repo.tx.Rollback()
		return err
	}

	return nil
}

func (r *WalletRepository) RollbackTx() error {
	if r.tx == nil {
		return errors.New("db tx not initiated")
	}

	r.tx.Rollback()
	return nil
}
