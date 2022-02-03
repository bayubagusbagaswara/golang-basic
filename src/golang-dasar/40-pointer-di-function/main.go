package main

import "fmt"

// bikin struct
type Address struct {
	City, Province, Country string
}

// kita bikin function, dimana parameternya adalah berupa Struct
func ChangeAddressToIndonesia(address *Address) {

	// kita ubah data address
	address.Country = "Indonesia"

}

func main() {

	// buat data struct
	address := Address{
		City:     "Kediri",
		Province: "Jawa Timur",
	}

	// eksekusi function ChangeAddressToIndonesia, dan masukkan parameter address nya
	ChangeAddressToIndonesia(&address)

	fmt.Println(address) // ini tidak berubah, tetap [Kediri, Jawa Timur]

	// KARENA SECARA DEFAULT NILAI YANG DIKIRIM KE PARAMETER FUNCTION ITU ADALAH PASS BY VALUE (DUPLIKAT DATA)

	// UNTUK BISA MENGUBAH DATA ASLINYA, MAKA GANTI POINTER DI PARAMETER FUNCTION NYA
}
