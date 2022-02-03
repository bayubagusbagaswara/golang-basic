# Defer

- `Defer Function` adalah function yang bisa kita jadwalkan untuk dieksekusi `setelah` sebuah function selesai di eksekusi
- Defer Function akan `selalu dieksekusi` walaupun terjadi error di function yang dieksekusi
- Ini mirip seperti `FINAL` di bahasa pemrograman lain
- Jadi tidak peduli sebuah function ketika dieksekusi itu sukses/gagal, tapi pada akhirnya tetap akan menjalankan `DEFER`

# Panic

- `Panic function` adalah function yang bisa kita gunakan untuk menghentikan program
- Panic function biasanya dipanggil ketika terjadi `error pada saat program kita berjalan`
- Saat panic function dipanggil, program akan terhenti. Namun defer function tetap akan dieksekusi
- Jadi, saat kita memanggil panic, maka otomatis program akan berhenti
- Ini mirip seperti `TRY` di bahasa pemrograman lain

# Recover

- Recover adalah function yang bisa kita gunakan untuk menangkap data panic
- Dengan recover proses panic akan terhenti, sehingga program akan tetap berjalan
- Ini mirip seperti `CATCH` di bahasa pemrograman lain, karena fungsinya juga menangkap error yang ada
