package main

import "fmt"

func getFullName() (string, string, int) {
	return "Bayu", "Bagaswara", 25
}

func main() {

	// TANGKAP SEMUA RETURN VALUE, SIMPAN KE MASING-MASING VARIABLE
	firstName, lastName, age := getFullName()
	fmt.Println(firstName, lastName, age)

	// IGNORE RETURN VALUE
	firstName1, _, _ := getFullName()
	fmt.Println(firstName1)
}
