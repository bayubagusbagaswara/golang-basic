package main

import "fmt"

// kita punya interface HasName
// dimana didalam interface terdapat kontrak method GetName() dengan return valuenya string
type HasName interface {
	GetName() string
}

// kita bikin function SayHello, dimana parameternya adalah Interface HasName
// sehingga didalam function SayHello kita bisa mengakses method interface GetName
func SayHello(hasName HasName) {
	fmt.Println("Hello ", hasName.GetName())
}

// kita bikin Struct Person
type Person struct {
	Name string
}

// kita bikin struct function, dimana namanya adalah GetName()
// Struct Function ini secara otomatis dideteksi sebagai implementasi Interface, karena di interface kita sudah punya GetName()
func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {

	// bikin data struct
	person := Person{
		Name: "Bayu",
	}

	// secara otomatis person sudah mengikuti kontrak dari HasName
	// makanya person bisa menggunakan function SayHello
	SayHello(person)

	var cat = Animal{
		Name: "Push",
	}

	SayHello(cat)

}
