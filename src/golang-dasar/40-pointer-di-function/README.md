# Pointer di Function

- Saat kita membuat parameter di function, maka secara default adalah `pass by value`. Artinya data akan di copy lalu dikirim ke function tersebut
- Oleh karena itu, jika kita mengubah data di dalam function, maka data yang aslinya tidak akan pernah berubah
- Hal ini membuat variable (discope luar) menjadi aman, karena tidak akan bisa diubah
- Namun terkadang kita ingin membuat function yang bisa mengubah data asli parameter function tersebut
- Untuk melakukan hal tersebut, kita juga bisa menggunkan pointer di function
- Untuk menjadikan sebuah parameter sebagai function, maka kita bisa menggunakan `operator *` di parameter function nya
