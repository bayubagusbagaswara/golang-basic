# Package Initialization

- Saat kita membuat package, kita bisa membuat sebuah function yang akan diakses ketika package kita diakses
- Ini sangat dicocok, apabila package kita berisi function-function untuk berkomunikasi dengan database, maka kita membuat function inisialisasi untuk membuka koneksi ke database
- Untuk membuat function yang diakses secara otomatis ketika package diakses, kita cukup membuat function dengan nama `init`

# Blank Identifier

- Kadang kita hanya ingin menjalankan init function di package tanpa harus mengeksekusi salah satu function yang ada di package
- Secara default, Go-Lang akan complain ketika ada package yang di import tetapi tidak digunakan
- Untuk menangani hal tersebut, kita bisa menggunakan blank identifier `(_)` sebelum nama package ketika melakukan import
