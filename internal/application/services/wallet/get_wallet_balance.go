package wallet

import (
	"github.com/aadejanovs/wallet/internal/application/services/wallet/dtos"
	repo "github.com/aadejanovs/wallet/internal/infrastructure/repositories/wallet"
)

type GetWalletService struct {
	repo *repo.WalletRepository
}

func NewGetWalletService(repo *repo.WalletRepository) *GetWalletService {
	return &GetWalletService{repo: repo}
}

func (s *GetWalletService) Get(id string) (*dtos.WalletResponseDto, error) {
	w, err := s.repo.OfId(id)
	if err != nil {
		return nil, err
	}

	return dtos.NewWalletResponseDto(w), nil
}
