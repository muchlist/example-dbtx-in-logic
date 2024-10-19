package repository

import (
	"context"

	"github.com/muchlist/example-dbtx-in-logic/internal/model"
	"github.com/muchlist/example-dbtx-in-logic/internal/port"
	dbtxgorm "github.com/muchlist/example-dbtx-in-logic/pkg/dbtx_gorm"
	"gorm.io/gorm"
)

// make sure the implementation satisfies the interface
var _ port.AccountStorer = (*repoWithGorm)(nil)

type repoWithGorm struct {
	db *gorm.DB
}

// NewRepo constructs a data for api access.
func NewRepoWithGorm(sqlDB *gorm.DB) *repoWithGorm {
	return &repoWithGorm{
		db: sqlDB,
	}
}

// GetAccountByID retrieves an account by its ID
func (r *repoWithGorm) GetAccountByID(ctx context.Context, id string) (model.AccountEntity, error) {
	dbtx := dbtxgorm.ExtractTx(ctx, r.db) // Extracting context and transforming standard db to DBTX interface

	var account model.AccountModel
	db := dbtx.WithContext(ctx).Model(&account).Where("id = ?", id).First(&account)
	if db.Error != nil {
		return model.AccountEntity{}, db.Error
	}

	return account.ToEntity(), nil
}

// UpdateAccount updates the details of an account
func (r *repoWithGorm) UpdateAccount(ctx context.Context, account model.AccountEntity) error {
	dbtx := dbtxgorm.ExtractTx(ctx, r.db) // Extracting context and transforming standard db to DBTX interface

	// Using GORM to update the account
	db := dbtx.WithContext(ctx).
		Model(&model.AccountModel{}).
		Where("id = ?", account.ID).
		Update("balance", account.Balance)
	return db.Error
}
