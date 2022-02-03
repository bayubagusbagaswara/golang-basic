package main

import "fmt"

func getCompleteName() (firstName, lastName string, age int, status bool) {

	firstName = "Bayu"
	lastName = "Bagaswara"
	age = 25
	status = true

	return
}

func main() {

	firstName, lastName, age, status := getCompleteName()
	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(age)
	fmt.Println(status)

	firstName1, _, _, status1 := getCompleteName()
	fmt.Println(firstName1, status1)
}
