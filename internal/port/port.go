package port

import (
	"context"

	"github.com/muchlist/example-dbtx-in-logic/internal/model"
)

//go:generate mockgen -source port.go -destination mockport/mock_port.go -package mockport
type TransferServiceAssumer interface {
	TransferMoney(ctx context.Context, input model.TransferDTO) error
}

type AccountStorer interface {
	GetAccountByID(ctx context.Context, id string) (model.AccountEntity, error)
	UpdateAccount(ctx context.Context, account model.AccountEntity) error
}

type TxManager interface {
	WithAtomic(ctx context.Context, tFunc func(ctx context.Context) error) error
}
