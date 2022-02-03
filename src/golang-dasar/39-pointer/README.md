# Pass by Value

- Secara default di Go-Lang, `semua variable` itu di passing by value, bukan by reference
- Artinya, jika kita mengirim sebuah variable ke dalam sebuah function, method, atau variable lain, maka sebenarnya yang dikirim itu adalah duplikasi value nya
- Jadi data di variable aslinya itu tidak berubah

# Pointer

- Pointer adalah kemampuan membuat reference ke lokasi data di memory yang sama, `tanpa menduplikasi data` yang sudah ada
- Sederhananya, dengan kemampuan pointer ini, kita bisa membuat `pass by reference`
- Jadi sumber data nya tetap satu, dan akan mereference/merujuk ke data utamanya

# Bagaimana caranya untuk melakukan Pointer?

- Menggunakan operator `&`
- Untuk membuat sebuah variable dengan nilai pointer ke variable yang lain, maka kita bisa menggunakan `operator &`, lalu diikuti nama variable nya
- Misal person2 dibuat dari person1, maka cukup tulis `&person1`. Artinya jika kita mengubah data di person2, maka akan mengubah juga di person1

# Operator bintang `*`

- Saat kita mengubah variable pointer, maka yang berubah hanya variable tersebut
- Semua variable yang mengacu/merujuk ke data yang sama tidak akan berubah
- Jika kita ingin mengubah seluruh variable yang mengacu ke data tersebut, maka kita bisa menggunakan operator `*`

# Function New

- Sebelumnya untuk membuat pointer dengan menggunakan `operator &`
- Go-Lang juga memiliki function new yang bisa digunakan untuk membuat pointer
- Namun function new hanya mengembalikan pointer ke data kosong, artinya tidak ada data awal
