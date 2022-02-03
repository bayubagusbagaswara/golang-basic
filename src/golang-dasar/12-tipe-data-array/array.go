package main

import "fmt"

func main() {

	// membuat variable bertipe array of string dengan panjang array 3
	var names [3]string

	// mengisi data di array berdasarkan index
	names[0] = "Bayu"
	names[1] = "Bagus"
	names[2] = "Bagaswara"

	// mendapatkan data array berdasarkan index
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// membuat array langsung ada nilainya
	var values = [3]int{
		90,
		80,
		95,
	}
	fmt.Println(values)

	age := [3]int{
		12, 13, 14,
	}
	fmt.Println(age)

	// function di array
	fmt.Println(len(names))
	fmt.Println(len(age))
}
