package main

import "fmt"

func sayHello() {
	fmt.Println("Hello")
}

func main() {
	// misal kita eksekusi function sayHello() sebanyak 10 kali
	for i := 0; i < 10; i++ {
		sayHello()
	}
}
