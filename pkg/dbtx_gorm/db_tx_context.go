package dbtxgorm

import (
	"context"

	"gorm.io/gorm"
)

type KeyTransaction string

const TXKey KeyTransaction = "project-name"

// ExtractTx extract transaction from context and transform database into db.DBTX
func ExtractTx(ctx context.Context, defaultTX *gorm.DB) *gorm.DB {
	if tx, ok := ctx.Value(TXKey).(*gorm.DB); ok {
		return tx
	}
	return defaultTX
}

// injectTx injects transaction to context
func injectTx(ctx context.Context, tx *gorm.DB) context.Context {
	return context.WithValue(ctx, TXKey, tx)
}
