package main

import "fmt"

func main() {
	var name string

	name = "Bayu Bagaswara"
	fmt.Println(name)

	name = "Bayu Bagus"
	fmt.Println(name)

	// bisa tanpa menulis keyword var diganti dengan :=
	country := "Indonesia"
	fmt.Println(country)

	country = "Germany"
	fmt.Println(country)

	// multiple variable
	var (
		firstName = "Bayu"
		lastName  = "Bagus"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}
