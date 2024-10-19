# example-dbtx-in-logic
Database transaction in logic layer with golang

Implementasi `database transaction` adalah aspek krusial dalam pengembangan aplikasi, terutama pada proyek yang menuntut konsistensi data yang tinggi. Artikel ini akan membahas bagaimana cara melakukan transaksi-database pada service layer (logic), dengan tetap mempertahankan prinsip-prinsip clean architecture dan separation of concerns.

Dalam arsitektur populer seperti Clean Architecture, Hexagonal Architecture, maupun pendekatan Domain-Driven Design (DDD), pemisahan tanggung jawab menjadi kunci utama. Kita umumnya membagi kode menjadi beberapa lapisan, misalnya Handler -> Service -> Repository. Lapisan service idealnya berisi logika bisnis murni tanpa bergantung pada library eksternal, sementara repository bertanggung jawab atas interaksi dengan database.

Namun, ketika berhadapan dengan transaksi ACID (Atomicity, Consistency, Isolation, Durability), muncul pertanyaan: di mana sebaiknya logika `database transaction` ditempatkan? Di lapisan logika atau di lapisan repository?

> note : Atomicity artinya Menjamin bahwa serangkaian operasi dalam satu transaksi harus sepenuhnya berhasil atau sepenuhnya gagal. 

Sebagai ilustrasi, mari kita tinjau kasus transfer uang antar rekening: "Transfer uang dari rekening A ke rekening B, perbarui semua data terkait, dan jika gagal, batalkan seluruh proses." Terdapat dua pendekatan umum:

Baca selanjutnya di : [Database Transaction in Service Layer](https://blog.muchlis.dev/post/db-transaction/) 
