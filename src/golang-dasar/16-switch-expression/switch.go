package main

import "fmt"

func main() {

	name := "Bayu"

	switch name {
	case "Bayu":
		fmt.Println("Hello Bayu")
	case "Joko":
		fmt.Println("Hello Joko")
	default:
		fmt.Println("Hi, Boleh Kenalan")
	}

	// SWITCH SORT STATEMENT
	// tidak harus kondisinya bernilai boolean
	// karena kita cek length > 5, maka menghasilkan nilai boolean
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Name Terlalu Panjang")
	case false:
		fmt.Println("Nama Sudah Benar")
	}

	// SWITCH TANPA KONDISI
	length1 := len(name)
	switch {
	case length1 > 10:
		fmt.Println("Nama Terlalu Panjang")
	case length1 > 5:
		fmt.Println("Nama Lumayan Panjang")
	default:
		fmt.Println("Nama Sudah Benar")
	}

}
