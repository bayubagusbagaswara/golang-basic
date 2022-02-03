package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {

	// Membuat Struct atau membuat object
	// nama data/object nya adalah bayu
	// tipe data struct nya adalah Customer
	// mirip Customer bayu = new Customer();
	var bayu Customer
	// set data
	bayu.Name = "Bayu Bagaswara"
	bayu.Address = "Indonesia"
	bayu.Age = 25

	fmt.Println(bayu.Name)
	fmt.Println(bayu.Address)
	fmt.Println(bayu.Age)

	// cara lain membuat data dari struct
	// mirip seperti Customer joko = new Customer("Joko", "Indonesia", 30)
	joko := Customer{
		Name:    "Joko",
		Address: "Indonesia",
		Age:     30,
	}
	fmt.Println(joko)

	budi := Customer{"Budi", "Indonesia", 30}
	fmt.Println(budi)
}
