package main

import "fmt"

// buat struct Address
type Address struct {
	City, Province, Country string
}

func main() {

	// PASS BY VALUE

	// kita buat address1
	address1 := Address{
		City:     "Kediri",
		Province: "Jawa Timur",
		Country:  "Indonesia",
	}

	// lalu kita assign address1 ke address2
	// atau address2 ini mengambil nilai dari address1
	address2 := address1

	// lalu kita coba ubah data di address2
	address2.City = "Surabaya"

	fmt.Println("Address1 : ", address1) // Kediri, Jawa Timur, Indonesia
	fmt.Println("Address2 : ", address2) // Surabaya, Jawa Timur, Indonesia
	fmt.Println("===============")

	// DI MEMORI SEBENARNYA ADDRESS1 DAN ADDRESS2 INI BERBEDA LOKASI MEMORI NYA

	// PASS BY REFERENCE

	address3 := Address{
		City:     "Bandung",
		Province: "Jawa Barat",
		Country:  "Indonesia",
	}

	// address4 akan reference ke address3
	address4 := &address3

	address4.City = "Sumedang"

	fmt.Println("Address3 : ", address3)
	fmt.Println("Address4 : ", address4)

	fmt.Println("=================")

	// OPERATOR *
	var address5 Address = Address{
		City:     "Subang",
		Province: "Jawa Barat",
		Country:  "Indonesia",
	}
	// address5 assign address6
	var address6 *Address = &address5
	var address7 *Address = &address5

	// SEBELUMMNYA KITA HANYA BUAT address6 reference ke address5
	// yang artinya tipa periubahan field di address6, juga akan mengubah field di address5
	// NAH BAGAIMANA JIKA address6 ini diubah secara total, tapi kita juga inging di address5 juga berubah total ??
	// caranya adalah menambahkan * di variable address6, lalu buat address6 dari struct Address yang baru (ditimpa semua data struct nya)
	// hasilnya adalah memory address5 akan berpindah ke memory address6
	// JADI SEMUA yang merujuk ke Address, datanya sudah berganti ke data yang baru (yang dibuat dari address6), dan data struct address5 sebelumnya sudah hilang
	*address6 = Address{
		City:     "Malang",
		Province: "Jawa Timur",
		Country:  "Indonesia",
	}

	address5.City = "Jayapura"

	fmt.Println("Address5 : ", address5)
	fmt.Println("Address6 : ", address6)
	fmt.Println("Address7 : ", address7)

	fmt.Println("=================")

	// FUNCTION NEW
	// akan dibuatkan Address yang isinya kosong
	alamat1 := new(Address)

	// alamat2 dibuat dari alamat1
	alamat2 := alamat1

	// set alamat2, karena awalnya datanya adalah kosong
	alamat2.Country = "Indonesia"

	fmt.Println(alamat1) // alamat 1 berubah
	fmt.Println(alamat2)

}
