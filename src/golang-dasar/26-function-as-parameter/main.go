package main

import "fmt"

// type declarations untuk alias
type Filter func(name string, foot int) (string, int)

func sayHelloWithFilter(name string, foot int, filter Filter) {
	name1, foot1 := spamFilter(name, foot)
	fmt.Println("Ini ", name1, "Kakinya", foot1)
}

func spamFilter(name string, foot int) (string, int) {
	if name == "Cat" && foot == 4 {
		return "Cat", 4
	} else {
		return "Bukan Kucing", 0
	}
}

func main() {

	// panggil function sayHelloWithFilter dan kirimkan parameter name dan parameter berupa function
	sayHelloWithFilter("Cat", 4, spamFilter)

	// function sebagai value, atau simpan function spamFilter di variable filter
	filter := spamFilter
	// panggil function sayHelloWithFilter, dan kirimkan parameter name dan parameter function
	sayHelloWithFilter("Anjing", 2, filter)
}
