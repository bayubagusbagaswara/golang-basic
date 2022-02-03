package main

import "fmt"

func main() {

	name := "Bayu"
	counter := 0

	// buat anonymous function
	increment := func() {

		// deklarasi ulang variable name
		name := "Aan"

		fmt.Println("Increment")

		// ini dinamakan Closure, karena mampu mengakses variable counter, dan bahayanya ketika bisa mengubah nilai dari counter, karena akan berpengaruh ke semua scope counter dibawah
		counter++

		fmt.Println(name) // Aan
	}

	// eksekusi function increment
	increment() // counter++
	increment() // counter++

	// lalu cetak nilai counter saat ini
	fmt.Println(counter)

	fmt.Println(name) // Bayu

	// SEBAIKNYA KALO CLOSURE HANYA UNTUK AKSES DATA DARI SCOPE YANG BERBEDA
	// TAPI JANGAN SAMPAI MENGUBAH DATA DARI SCOPE YANG BERBEDA
	// BISA TERJADI INKONSISTENSI DATA, KARENA TIAP FUNCTION DIEKSEKUSI, MAKA AKAN MENGUBAH NILAI TERHADAP DATA DI SCOPE LUAR
	// DAN LEBIH BAIK BUAT VARIABLE BERBEDA TIAP SCOPE KODE PROGRAM
	// ATAU DEKLARASI ULANG VARIABLENYA, ATAU OVVERRIDE
}
