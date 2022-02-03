package main

import "fmt"

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func main() {
	// membuat variable dulu
	firstName := "Bayu"
	sayHelloTo(firstName, "Bagaswara")

	// atau bisa memasukkan data langsung di parameter function nya
	sayHelloTo("Budi", "Nugraha")
}
