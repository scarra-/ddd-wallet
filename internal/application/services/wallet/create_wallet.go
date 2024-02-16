package wallet

import (
	"github.com/aadejanovs/wallet/internal/application/services/wallet/dtos"
	domain "github.com/aadejanovs/wallet/internal/domain/model"
	repo "github.com/aadejanovs/wallet/internal/infrastructure/repositories/wallet"
)

type CreateWalletService struct {
	repo *repo.WalletRepository
}

func NewCreateWalletService(repo *repo.WalletRepository) *CreateWalletService {
	return &CreateWalletService{repo: repo}
}

func (s *CreateWalletService) Create(dto *dtos.CreateWalletRequestDto) (*dtos.WalletResponseDto, error) {
	w := domain.NewWallet(dto.OwnerId, dto.Currency)
	err := s.repo.Save(w)

	return dtos.NewWalletResponseDto(w), err
}
