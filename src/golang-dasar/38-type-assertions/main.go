package main

import "fmt"

// kita bikin function dengan return value nya adalah interface kosong
// tetapi di dalam function hanya mengembalikan tipe data string
func random() interface{} {
	return "OK"
}

func main() {

	// random() ini mengembalikan interface kosong, artinya bisa berbagai macam tipe data
	result := random()

	// kita ubah result menjadi string
	resultString := result.(string)
	fmt.Println(resultString)

	// jika kita salah mengubah tipe data dari interface kosong, maka akan terjadi panic
	resultInt := result.(int) // panic
	fmt.Println(resultInt)

	// MENGGUNAKAN SWITCH
	switch value := result.(type) {
	case string:
		fmt.Println("String", value, "is string")
	case int:
		fmt.Println("Int", value, "is int")
	default:
		fmt.Println("Unknown")
	}
}
