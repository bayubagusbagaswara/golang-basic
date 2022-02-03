package main

import "fmt"

func main() {

	// kita bikin map[string]string{}, artinya key nya bertipe string, dan valuenya bertipe string
	person := map[string]string{
		"name":    "Bayu",
		"address": "Kediri",
	}

	fmt.Println(person)
	// cara mengaksesnya yaitu dengan cara ambil key nya
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	// untuk menambah data ke map
	person["title"] = "Programmer"

	// hitung panjang map
	fmt.Println(len(person))

	// membuat map baru
	var book map[string]string = make(map[string]string)
	book["title"] = "Belajar Go-Lang"
	book["author"] = "Bayu"
	book["ups"] = "Salah"

	// menghapus data di map, cukup sebutkan nama map nya, lalu key nya
	delete(book, "ups")

	fmt.Println(book)
}
