package main

import "fmt"

func main() {

	// misal kita punya variable counter bernilai 1
	counter := 1

	// kita ingin mencetak counter dengan nilai 1-10
	// maka kita bikin perulangan 10 kali, diakhir tiap perulangan kita naikkan nilai counternya counter = counter + 1 atau counter++
	for counter <= 10 {
		fmt.Println("Perulangan ke ", counter)
		counter = counter + 1
	}

	// FOR DENGAN STATEMENT
	// init statement adalah counter1 = 1
	// kondisi adalah counter <= 10
	// post statement adalah counter++
	for counter1 := 1; counter1 <= 10; counter1++ {
		fmt.Println("Perulangan ke", counter1)
	}

	// bikin slice isinya array
	names := []string{"Bayu", "Bagus", "Bagaswara"}

	// kita mau iterasi dari data-datanya
	for index, name := range names {
		fmt.Println("Index", index, "=", name)
	}

	// bisa tidak cetak index nya, diganti dengan tanda underscore _
	for _, name1 := range names {
		fmt.Println("Name", name1)
	}

	// kita coba iterasi data dari map
	person := make(map[string]string)
	person["name"] = "Bayu"
	person["title"] = "Programmer"

	// ambil semua data dari map
	for key, value := range person {
		fmt.Println("Key:", key, "Value:", value)
	}
}
