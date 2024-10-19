package dbtxgorm

// jika menggunakan gorm, kita tidak memerlukan interface DBTX ini, karena gorm sudah menggabungkan
// antara kemampuan db transaction dan operasi db biasa ke *gorm.DB
// jadi kita tinggal menggunakan *gorm.DB saja
type DBTX interface {
	// ... cek file pkg/dbtx.db_tx.go untuk melihat perbedaannya
}
