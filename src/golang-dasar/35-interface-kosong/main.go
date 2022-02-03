package main

import "fmt"

// bikin function dimana return valuenya adalah interface kosong
func Ups() interface{} {

	// karena interface kosong, maka returnya bisa berbagai tipe data
	// seperti int, boolean, string
	// return 1
	// return true
	return "Ups"
}

func Hai(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Ups"
	}
}

func main() {

	var data interface{} = Hai(1)
	fmt.Println(data)

}
