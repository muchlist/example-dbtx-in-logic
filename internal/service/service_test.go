package service

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/muchlist/example-dbtx-in-logic/internal/model"
	"github.com/muchlist/example-dbtx-in-logic/internal/port/mockport"
)

func Test_service_TransferMoney(t *testing.T) {
	tests := []struct {
		name    string
		mock    func(as *mockport.MockAccountStorer, tx *mockport.MockTxManager)
		input   model.TransferDTO
		wantErr bool
	}{
		{
			name: "Should be success",
			mock: func(as *mockport.MockAccountStorer, tx *mockport.MockTxManager) {
				// menjalankan semua proses yang ada didalam WithAtomic
				tx.EXPECT().WithAtomic(gomock.Any(), gomock.Any()).
					DoAndReturn(func(x any, tFunc func(ctx context.Context) error) error {
						return tFunc(context.Background())
					})
				as.EXPECT().GetAccountByID(gomock.Any(), "1").Return(
					model.AccountEntity{
						ID:      "1",
						Balance: 10000,
					}, nil,
				)
				as.EXPECT().GetAccountByID(gomock.Any(), "2").Return(
					model.AccountEntity{
						ID:      "2",
						Balance: 20000,
					}, nil,
				)
				as.EXPECT().UpdateAccount(gomock.Any(), model.AccountEntity{
					ID:      "1",
					Balance: 5000,
				}).Return(
					nil,
				)
				as.EXPECT().UpdateAccount(gomock.Any(), model.AccountEntity{
					ID:      "2",
					Balance: 25000,
				}).Return(
					nil,
				)

			},
			input: model.TransferDTO{
				AccountA: "1",
				AccountB: "2",
				Amount:   5000,
			},
			wantErr: false,
		},
		{
			name: "Should be failed, balance not enough",
			mock: func(as *mockport.MockAccountStorer, tx *mockport.MockTxManager) {
				// menjalankan semua proses yang ada didalam WithAtomic
				tx.EXPECT().WithAtomic(gomock.Any(), gomock.Any()).
					DoAndReturn(func(x any, tFunc func(ctx context.Context) error) error {
						return tFunc(context.Background())
					})
				as.EXPECT().GetAccountByID(gomock.Any(), "1").Return(
					model.AccountEntity{
						ID:      "1",
						Balance: 4000,
					}, nil,
				)
				as.EXPECT().GetAccountByID(gomock.Any(), "2").Return(
					model.AccountEntity{
						ID:      "2",
						Balance: 20000,
					}, nil,
				)
			},
			input: model.TransferDTO{
				AccountA: "1",
				AccountB: "2",
				Amount:   5000,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mockport.NewMockAccountStorer(ctrl)
			txMan := mockport.NewMockTxManager(ctrl)

			tt.mock(repo, txMan)

			service := NewService(repo, txMan)

			if err := service.TransferMoney(ctx, tt.input); (err != nil) != tt.wantErr {
				t.Errorf("service.TransferMoney() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
