package model

/*
agar simple saya menggabungkan entity, model, dto kedalam 1 package
idealnya dipisahkan ke beberapa package
*/

// ========================= ENTITY
type AccountEntity struct {
	ID      string
	Balance float64
}

// ========================= DTO
type TransferDTO struct {
	AccountA string
	AccountB string
	Amount   float64
}

// ========================= MODEL
type AccountModel struct {
	ID      string  `gorm:"column:id"`
	Balance float64 `gorm:"column:balance"`
}

func (AccountModel) TableName() string {
	return "accounts"
}

func (a *AccountModel) ToEntity() AccountEntity {
	return AccountEntity{
		ID:      a.ID,
		Balance: a.Balance,
	}
}
