package main

import "fmt"

func main() {
	fmt.Println("Bayu")
	fmt.Println("Bagus")
	fmt.Println("Bagaswara")

	// hitung length Bagaswara
	fmt.Println(len("Bagaswara"))

	// ambil karakter[0], tapi ini balikannya adalah number representasi karakter B
	fmt.Println("Bayu Bagus"[0])

	// ambil string index[1]
	fmt.Println("Bayu Bagus Bagaswara"[1])
}
