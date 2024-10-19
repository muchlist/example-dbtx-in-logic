package dbtxgorm

import (
	"context"
	"fmt"

	"log/slog"

	"gorm.io/gorm"
)

type TxManager interface {
	WithAtomic(ctx context.Context, tFunc func(ctx context.Context) error) error
}

type txManager struct {
	db  *gorm.DB
	log *slog.Logger
}

func NewTxManager(sqlDB *gorm.DB, log *slog.Logger) TxManager {
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
	tx := r.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return fmt.Errorf("begin transaction: %w", tx.Error)
	}

	// run callback
	err := tFunc(injectTx(ctx, tx))
	if err != nil {
		// if error, rollback
		if txRoleback := tx.WithContext(ctx).Rollback(); txRoleback.Error != nil {
			r.log.Warn("rollback transaction", slog.String("error", txRoleback.Error.Error()))
		}
		return err
	}
	// if no error, commit
	if txCommit := tx.WithContext(ctx).Commit(); txCommit.Error != nil {
		return fmt.Errorf("failed to commit transaction: %w", txCommit.Error)
	}
	return nil
}
