package wallet

import (
	"github.com/aadejanovs/wallet/internal/application/services/wallet/dtos"
	domain "github.com/aadejanovs/wallet/internal/domain/model"
	txRepo "github.com/aadejanovs/wallet/internal/infrastructure/repositories/transaction"
	walletRepo "github.com/aadejanovs/wallet/internal/infrastructure/repositories/wallet"
)

type SpendWalletFundsService struct {
	walletRepo *walletRepo.WalletRepository
	txRepo     *txRepo.TransactionRepository
}

func NewSpendWalletFundsService(
	walletRepo *walletRepo.WalletRepository,
	txRepo *txRepo.TransactionRepository,
) *FundWalletService {
	return &FundWalletService{
		walletRepo: walletRepo,
		txRepo:     txRepo,
	}
}

func (s *FundWalletService) Spend(dto *dtos.SpendWalletFundsRequestDto) (*dtos.TxResponseDto, error) {
	wallet, err := s.walletRepo.OfIdForUpdate(dto.WalletId)
	if err != nil {
		return nil, err
	}

	spendSource := domain.NewSpendSource(dto.Type, dto.OriginId)

	tx, err := wallet.Spend(
		domain.NewTxChecker(s.txRepo),
		dto.Amount,
		spendSource,
	)
	if err != nil {
		s.walletRepo.RollbackTx()
		return nil, err
	}

	err = s.walletRepo.SaveWalletAndTx(wallet, tx)
	if err != nil {
		return nil, err
	}

	return dtos.NewTxResponseDto(tx), nil
}
