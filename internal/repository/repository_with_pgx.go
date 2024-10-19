package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/muchlist/example-dbtx-in-logic/internal/model"
	"github.com/muchlist/example-dbtx-in-logic/internal/port"
	"github.com/muchlist/example-dbtx-in-logic/pkg/dbtx"
)

// make sure the implementation satisfies the interface
var _ port.AccountStorer = (*repo)(nil)

type repo struct {
	db *pgxpool.Pool
}

// NewRepo constructs a data for api access.
func NewRepo(sqlDB *pgxpool.Pool) *repo {
	return &repo{
		db: sqlDB,
	}
}

// GetAccountByID retrieves an account by its ID
func (r *repo) GetAccountByID(ctx context.Context, id string) (model.AccountEntity, error) {
	dbtx := dbtx.ExtractTx(ctx, r.db) // Extracting context and transforming standard db to DBTX interface

	var account model.AccountEntity
	err := dbtx.QueryRow(ctx, "SELECT id, balance FROM accounts WHERE id = $1", id).Scan(
		&account.ID, &account.Balance)

	return account, err
}

// UpdateAccount updates the details of an account
func (r *repo) UpdateAccount(ctx context.Context, account model.AccountEntity) error {
	dbtx := dbtx.ExtractTx(ctx, r.db) // Extracting context and transforming standard db to DBTX interface

	_, err := dbtx.Exec(ctx, `
        UPDATE accounts 
        SET balance = $1
        WHERE id = $2`,
		account.Balance, account.ID)

	return err
}
