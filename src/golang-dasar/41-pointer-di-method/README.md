# Pointer di Method

- Walaupun sebuah method itu akan menempel di struct, tetapi sebenarnya data struct yang diakses di dalam method adalah `pass by value`
- Sangat direkomendasikan menggunakan pointer di method, sehingga tidak boros memory karena harus selalu diduplikasi ketika memanggil method
- Jadi, data struct nya itu tidak berubah, karena yang berubah adalah data struct yang di passing ke method aja
- Jika ingin data struct nya juga berubah, maka gunakan pointer di method struct nya
