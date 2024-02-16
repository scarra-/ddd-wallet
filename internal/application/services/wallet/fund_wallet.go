package wallet

import (
	"github.com/aadejanovs/wallet/internal/application/services/wallet/dtos"
	domain "github.com/aadejanovs/wallet/internal/domain/model"
	txRepo "github.com/aadejanovs/wallet/internal/infrastructure/repositories/transaction"
	walletRepo "github.com/aadejanovs/wallet/internal/infrastructure/repositories/wallet"
)

type FundWalletService struct {
	walletRepo *walletRepo.WalletRepository
	txRepo     *txRepo.TransactionRepository
}

func NewFundWalletService(
	walletRepo *walletRepo.WalletRepository,
	txRepo *txRepo.TransactionRepository,
) *FundWalletService {
	return &FundWalletService{
		walletRepo: walletRepo,
		txRepo:     txRepo,
	}
}

func (s *FundWalletService) Fund(dto *dtos.FundWalletRequestDto) (*dtos.TxResponseDto, error) {
	wallet, err := s.walletRepo.OfIdForUpdate(dto.WalletId)
	if err != nil {
		return nil, err
	}

	fundSource := domain.NewFundSource(dto.Type, dto.OriginId)

	tx, err := wallet.Fund(
		domain.NewTxChecker(s.txRepo),
		dto.Amount,
		fundSource,
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
