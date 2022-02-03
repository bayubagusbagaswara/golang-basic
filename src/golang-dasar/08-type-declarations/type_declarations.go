package main

import "fmt"

func main() {

	// cukup gunakan kata kunci type
	// artinya NoKTP ini adalah type declarations dari tipe data string
	type NoKTP string
	type Married bool

	// kita bisa menggunakan NoKTP sebagai tipe data baru, dimana artinya adalah alias untuk tipe data string
	var ktpBayu NoKTP = "12345678"
	var marriedStatus Married = true

	// ktpBayu bernilai string
	fmt.Println(ktpBayu)
	// NoKTP juga bernilai string
	fmt.Println(NoKTP("2222222222222"))
	fmt.Println(marriedStatus)
}
