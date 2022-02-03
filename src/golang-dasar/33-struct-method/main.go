package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

// kita buat function yang ditempelkan ke struct
func (pelanggan Customer) sayHello(name string) {
	fmt.Println("Hello", pelanggan.Name, "My Name is", name)
}

func (pelanggan Customer) sayHuu() {
	fmt.Println("Huuu", pelanggan.Name)
}

func main() {

	var bayu Customer
	bayu.Name = "Bayu"
	bayu.Address = "Indonesia"
	bayu.Age = 25

	// panggil struct function
	// function nya bisa langsung nempel ke struct data nya
	bayu.sayHello("Joko")
	bayu.sayHuu()
}
