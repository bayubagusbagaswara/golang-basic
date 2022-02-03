package main

import "fmt"

func main() {

	var nilai32 int32 = 32768

	// hasil konversi int32 menjadi int64
	var nilai64 int64 = int64(nilai32)

	// kita hati-hati kalau konversi ke tipe data yang lebih rendah
	// konversi int32 menjadi int16
	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	// cara konversi dari byte menjadi string
	var name = "Bayu Bagaswara"

	// ambil karakter index 0
	var e = name[0]

	// konversi byte ke string
	var eString = string(e)

	fmt.Println(name)
	fmt.Println(eString)
}
