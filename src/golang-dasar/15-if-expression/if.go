package main

import "fmt"

func main() {

	// IF EXPRESSION
	name := "Bayu"

	if name == "Bayu" {
		fmt.Println("Hello Bayu")
	}

	// ELSE EXPRESSION
	if name == "Bayu" {
		fmt.Println("Hello Bayu")
	} else {
		// jika if bernilai false, maka blok di else yang dieksekusi
		fmt.Println("Hi")
	}

	// ELSE IF

	name = "Joko"
	if name == "Bayu" {
		fmt.Println("Hello Bayu")
	} else if name == "Joko" {
		fmt.Println("Hello Joko")
	} else {
		// jika if bernilai false, maka blok di else yang dieksekusi
		fmt.Println("Hi")
	}

	// SHORT STATEMENT
	// misal kita cek panjang name, apakah terlalu panjang atau gak
	// jika panjang name itu lebih dari 5, maka terlalu panjang
	// ini seperti kita buat dua buah variable, tapi dipersingkat menjadi satu dengan syarat kondisi

	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
