package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/muchlist/example-dbtx-in-logic/internal/model"
	"github.com/muchlist/example-dbtx-in-logic/internal/repository"
	"github.com/muchlist/example-dbtx-in-logic/internal/service"
	"github.com/muchlist/example-dbtx-in-logic/pkg/dbtx"
	dbtxgorm "github.com/muchlist/example-dbtx-in-logic/pkg/dbtx_gorm"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// ==================================
	// example depency injection with pgx
	pgxpool, _ := dbtx.OpenDB(dbtx.Config{
		DSN:          "",
		MaxOpenConns: 0,
		MinOpenConns: 0,
	})

	repo := repository.NewRepo(pgxpool)
	txManager := dbtx.NewTxManager(pgxpool, logger)
	usecase := service.NewService(repo, txManager)

	// test transfer
	_ = usecase.TransferMoney(context.Background(), model.TransferDTO{
		AccountA: "1",
		AccountB: "2",
		Amount:   5000,
	})

	// ==================================
	// example depency injection with gorm
	gormDB, _ := dbtxgorm.OpenDB(dbtxgorm.Config{
		DSN:          "",
		MaxOpenConns: 0,
		MinOpenConns: 0,
	})

	repoWithGorm := repository.NewRepoWithGorm(gormDB)
	txManagerGorm := dbtxgorm.NewTxManager(gormDB, logger)
	usecaseGorm := service.NewService(repoWithGorm, txManagerGorm)

	// test transfer
	_ = usecaseGorm.TransferMoney(context.Background(), model.TransferDTO{
		AccountA: "1",
		AccountB: "2",
		Amount:   5000,
	})

	// perhatikan service.NewService() tidak terpengaruh pada perubahan apapun
	// artinya kita telah berhasil menjaga core logic semurni mungkin
}
