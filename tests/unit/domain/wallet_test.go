package domain

import (
	"errors"
	"testing"

	d "github.com/aadejanovs/wallet/internal/domain/model"
	"github.com/stretchr/testify/assert"
)

type txRepoStub struct {
	recordFound bool
}

func (r *txRepoStub) OfOriginId(id string) (*d.Transaction, error) {
	if r.recordFound {
		return &d.Transaction{}, nil
	}

	return nil, errors.New("record not found")
}

func TestFundWallet(t *testing.T) {
	t.Run("create and fund wallet", func(t *testing.T) {
		wallet := d.NewWallet("usr-AAaPeAqXYSTsuHqbOAyoLagvYxij", "eur")
		assert.Equal(t, 0, wallet.Balance)

		source := d.NewFundSource("card_deposit", "ch-aAAwVnWmJYQVpeCkPGMdOubktYVA6")

		txChecker := d.NewTxChecker(&txRepoStub{recordFound: false})
		tx, err := wallet.Fund(txChecker, 100, source)
		assert.Nil(t, err)

		assert.Equal(t, wallet.Id, tx.WalletId)
		assert.Equal(t, 100, tx.Amount)
		assert.Equal(t, "card_deposit", tx.TxType)
		assert.Equal(t, source.OriginId, tx.OriginId)
		assert.Equal(t, 100, wallet.Balance)

		txChecker = d.NewTxChecker(&txRepoStub{recordFound: false})
		wallet.Fund(txChecker, 200, source)
		assert.Equal(t, 300, wallet.Balance)
	})

	t.Run("double fund wallet with same origin", func(t *testing.T) {
		wallet := d.NewWallet("usr-AAaPeAqXYSTsuHqbOAyoLagvYxij", "eur")
		source := d.NewFundSource("card_deposit", "ch-aAAwVnWmJYQVpeCkPGMdOubktYVA6")

		txChecker := d.NewTxChecker(&txRepoStub{recordFound: false})
		wallet.Fund(txChecker, 100, source)

		txChecker = d.NewTxChecker(&txRepoStub{recordFound: true})
		_, err := wallet.Fund(txChecker, 100, source)

		assert.ErrorIs(t, err, d.OriginIdAlreadyUsed)
	})

	t.Run("fund wallet with invalid amount", func(t *testing.T) {
		wallet := d.NewWallet("usr-AAaPeAqXYSTsuHqbOAyoLagvYxij", "eur")
		source := d.NewFundSource("card_deposit", "ch-aAAwVnWmJYQVpeCkPGMdOubktYVA6")

		txChecker := d.NewTxChecker(&txRepoStub{recordFound: false})
		tx, err := wallet.Fund(txChecker, -100, source)
		assert.Nil(t, tx)
		assert.ErrorIs(t, err, d.InvalidAmount)
	})
}

func TestSpendWalletFunds(t *testing.T) {
	t.Run("spend wallet funds", func(t *testing.T) {
		wallet := d.NewWallet("usr-AAaPeAqXYSTsuHqbOAyoLagvYxij", "eur")
		fundSource := d.NewFundSource("card_deposit", "ch-aAAwVnWmJYQVpeCkPGMdOubktYVA6")
		spendSource := d.NewSpendSource("card_deposit", "ch-aAAwVnWmJYQVpeCkPGMdOubktYVA1")

		txChecker := d.NewTxChecker(&txRepoStub{recordFound: false})
		wallet.Fund(txChecker, 100, fundSource)

		tx, err := wallet.Spend(txChecker, 50, spendSource)
		assert.Nil(t, err)

		assert.Equal(t, wallet.Id, tx.WalletId)
		assert.Equal(t, -50, tx.Amount)
		assert.Equal(t, "card_deposit", tx.TxType)
		assert.Equal(t, spendSource.OriginId, tx.OriginId)
		assert.Equal(t, 50, wallet.Balance)

		tx, err = wallet.Spend(txChecker, 40, spendSource)
		assert.Nil(t, err)
		assert.Equal(t, -40, tx.Amount)
		assert.Equal(t, 10, wallet.Balance)

		tx, err = wallet.Spend(txChecker, 10, spendSource)
		assert.Nil(t, err)
		assert.Equal(t, -10, tx.Amount)
		assert.Equal(t, 0, wallet.Balance)
	})

	t.Run("double spend wallet funds", func(t *testing.T) {
		wallet := d.NewWallet("usr-AAaPeAqXYSTsuHqbOAyoLagvYxij", "eur")
		fundSource := d.NewFundSource("card_deposit", "ch-aAAwVnWmJYQVpeCkPGMdOubktYVA6")
		spendSource := d.NewSpendSource("card_deposit", "ch-aAAwVnWmJYQVpeCkPGMdOubktYVA1")

		txChecker := d.NewTxChecker(&txRepoStub{recordFound: false})
		wallet.Fund(txChecker, 100, fundSource)

		_, err := wallet.Spend(txChecker, 50, spendSource)
		assert.Nil(t, err)

		txChecker = d.NewTxChecker(&txRepoStub{recordFound: true})
		tx, err := wallet.Spend(txChecker, 50, spendSource)

		assert.Nil(t, tx)
		assert.ErrorIs(t, err, d.OriginIdAlreadyUsed)
	})

	t.Run("spend wallet funds with invalid amount", func(t *testing.T) {
		wallet := d.NewWallet("usr-AAaPeAqXYSTsuHqbOAyoLagvYxij", "eur")
		fundSource := d.NewFundSource("card_deposit", "ch-aAAwVnWmJYQVpeCkPGMdOubktYVA6")
		spendSource := d.NewSpendSource("card_deposit", "ch-aAAwVnWmJYQVpeCkPGMdOubktYVA1")

		txChecker := d.NewTxChecker(&txRepoStub{recordFound: false})
		wallet.Fund(txChecker, 100, fundSource)

		tx, err := wallet.Spend(txChecker, -50, spendSource)
		assert.Nil(t, tx)
		assert.ErrorIs(t, err, d.InvalidAmount)
	})

	t.Run("insufficient wallet funds", func(t *testing.T) {
		wallet := d.NewWallet("usr-AAaPeAqXYSTsuHqbOAyoLagvYxij", "eur")
		fundSource := d.NewFundSource("card_deposit", "ch-aAAwVnWmJYQVpeCkPGMdOubktYVA6")
		spendSource := d.NewSpendSource("card_deposit", "ch-aAAwVnWmJYQVpeCkPGMdOubktYVA1")

		txChecker := d.NewTxChecker(&txRepoStub{recordFound: false})

		tx, err := wallet.Spend(txChecker, 50, spendSource)
		assert.Nil(t, tx)
		assert.ErrorIs(t, err, d.InsufficientFunds)

		wallet.Fund(txChecker, 100, fundSource)

		tx, err = wallet.Spend(txChecker, 101, spendSource)
		assert.Nil(t, tx)
		assert.ErrorIs(t, err, d.InsufficientFunds)
	})
}
