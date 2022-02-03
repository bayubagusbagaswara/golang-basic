package main

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye " + name
}

func main() {

	// simpan function di dalam sebuah variable
	sayGoodBye := getGoodBye

	// panggil variable dan masukkan parameternya
	result := sayGoodBye("Bayu")
	fmt.Println(result)

	// atau panggil langsung dan masukkan parameter
	fmt.Println(getGoodBye("Bayu"))
}
