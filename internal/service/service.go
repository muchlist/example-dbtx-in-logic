package service

import (
	"context"
	"errors"

	"github.com/muchlist/example-dbtx-in-logic/internal/model"
	"github.com/muchlist/example-dbtx-in-logic/internal/port"
)

// make sure the implementation satisfies the interface
var _ port.TransferServiceAssumer = (*service)(nil)

type service struct {
	Repo      port.AccountStorer
	TxManager port.TxManager // helper untuk transaction menjadi dependecy tambahan atau bisa digabung ke repo
}

func NewService(
	repo port.AccountStorer,
	txManager port.TxManager,
) *service {
	return &service{
		Repo:      repo,
		TxManager: txManager,
	}
}

func (s *service) TransferMoney(ctx context.Context, input model.TransferDTO) error {

	// shared variable untuk menampung hasil didalam WithAtomic jika ada
	// result := ...

	// Membungkus prosesnya dengan database transaction
	txErr := s.TxManager.WithAtomic(ctx, func(ctx context.Context) error {
		// Mengambil account A
		accountA, err := s.Repo.GetAccountByID(ctx, input.AccountA)
		if err != nil {
			return err // Gagal mengambil account A
		}

		// Mengambil account B
		accountB, err := s.Repo.GetAccountByID(ctx, input.AccountB)
		if err != nil {
			return err // Gagal mengambil account B
		}

		// Memeriksa apakah saldo account A cukup
		if accountA.Balance < input.Amount {
			return errors.New("saldo tidak cukup") // Gagal karena saldo tidak cukup
		}

		// Mengurangi saldo account A
		accountA.Balance -= input.Amount
		if err := s.Repo.UpdateAccount(ctx, accountA); err != nil {
			return err // Gagal update saldo account A
		}

		// Menambahkan jumlah ke saldo account B
		accountB.Balance += input.Amount
		if err := s.Repo.UpdateAccount(ctx, accountB); err != nil {
			return err // Gagal update saldo account B
		}

		return nil
	})

	if txErr != nil {
		return txErr
	}

	return nil
}
