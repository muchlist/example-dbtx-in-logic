package dbtx

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TxManager interface {
	WithAtomic(ctx context.Context, tFunc func(ctx context.Context) error) error
}

type txManager struct {
	db  *pgxpool.Pool
	log *slog.Logger
}

func NewTxManager(sqlDB *pgxpool.Pool, log *slog.Logger) TxManager {
	return &txManager{
		db:  sqlDB,
		log: log,
	}
}

// =========================================================================
// TRANSACTION

// WithAtomic runs function within transaction
// The transaction commits when function were finished without error
func (r *txManager) WithAtomic(ctx context.Context, tFunc func(ctx context.Context) error) error {

	// begin transaction
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}

	// run callback
	err = tFunc(injectTx(ctx, tx))
	if err != nil {
		// if error, rollback
		if errRollback := tx.Rollback(ctx); errRollback != nil {
			r.log.Error("rollback transaction", slog.String("error", errRollback.Error()))
		}
		return err
	}
	// if no error, commit
	if errCommit := tx.Commit(ctx); errCommit != nil {
		return fmt.Errorf("failed to commit transaction: %w", errCommit)
	}
	return nil
}
