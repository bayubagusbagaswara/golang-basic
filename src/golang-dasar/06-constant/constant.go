package main

import "fmt"

func main() {
	const firstName string = "Bayu"
	const lastName = "Bagaswara"
	const age = 25

	// error, variable constant tidak bisa diubah lagi
	// firstName = "Tidak Bisa Diubah"
	// lastName = "Tidak Bisa Diubah"

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(age)
}
