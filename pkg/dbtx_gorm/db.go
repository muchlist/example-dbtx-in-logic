package dbtxgorm

import (
	"database/sql"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	DSN          string
	MaxOpenConns int
	MinOpenConns int
}

func OpenDB(cfg Config) (*gorm.DB, error) {

	pgxDB, err := sql.Open("pgx", cfg.DSN)
	if err != nil {
		return nil, err
	}
	pgxDB.SetMaxIdleConns(cfg.MinOpenConns)
	pgxDB.SetMaxOpenConns(cfg.MaxOpenConns)
	// pgxDB.SetConnMaxLifetime(option.ConnMaxLifetime)
	// pgxDB.SetConnMaxIdleTime(option.ConnMaxIdleTime)

	db, err := gorm.Open(
		postgres.New(postgres.Config{Conn: pgxDB}),
		&gorm.Config{
			SkipDefaultTransaction: true,
		})
	if err != nil {
		return nil, err
	}

	return db, nil
}
