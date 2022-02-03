package main

import (
	"43-package-dan-import/helper"
	"fmt"
)

func main() {

	// jika kita ingin mengakses function SayHello, maka kita harus lakukan import dulu package nya
	result := helper.SayHello("Bayu")

	fmt.Println(result)

}
