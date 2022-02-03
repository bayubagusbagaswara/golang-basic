package main

import (
	"errors"
	"fmt"
)

// bikin fucntion Pembagi
// jika suatu function itu bisa menyebabkan error, maka return value nya ada 2
// yakni data suksesnya dan data errornya
func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh 0")
	} else {
		// jika tidak terjadi error, maka kita balikan nil untuk data errornya
		result := nilai / pembagi
		return result, nil
	}
}

func main() {

	// simpan return value Pembagian didalam dua variable
	hasil, err := Pembagian(100, 10)
	if err == nil {
		fmt.Println("Hasil", hasil)
	} else {
		fmt.Println("Error", err.Error())
	}

}
