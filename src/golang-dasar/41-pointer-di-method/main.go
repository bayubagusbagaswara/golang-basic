package main

import "fmt"

// bikin struct
type Man struct {
	Name string
}

// contoh method struct yang tidak menggunakan pointer
// jadi parameter struct man disini adalah pass by value, bukan data struct yang asli
func (man Man) Married() {
	// kita ubah name menjadi data baru
	man.Name = "Mr. " + man.Name
	fmt.Println("Di Method", man.Name)
}

// SOLUSI agar data struct nya ikut berubah, maka gunakan pointer sebagai parameter data struct
func (man *Man) MarriedPointer() {
	// kita ubah name menjadi data baru
	man.Name = "Mr. " + man.Name
	fmt.Println("Di Method", man.Name)
}

func main() {

	bayu := Man{
		Name: "Bayu",
	}

	// panggil method struct
	bayu.Married()

	fmt.Println(bayu.Name) // hasilnya tetap Bayu

	bagus := Man{
		Name: "Bagus",
	}

	bagus.MarriedPointer()
	fmt.Println(bagus.Name) // sudah berubah Mr. Bagus

}
