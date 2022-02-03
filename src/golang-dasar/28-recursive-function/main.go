package main

import "fmt"

// parameter adalah value dengan tipe data int
// dan return value function adalah int
func factorialLoop(value int) int {

	// initialisasi nilai awal
	result := 1

	for i := value; i > 0; i-- {
		// i nya dimulai dari value, terus menurun kebawah
		// misal i = 5, maka selanjutnya adalah 4,3,2,1
		result = result * i
		// result = 1 * 5 = 5
		// result = 5 * 4 = 20
		// result = 20 * 3 = 60
		// result = 60 * 2 = 120
		// result = 120 * 1 = 120
		// i = 0, keluar dari perulangan
	}
	// selesai perulangan
	return result
}

// kita bikin function recursive
func factorialRecursive(value int) int {

	// jika input 1 artinya bahwa dia menghitung 1!, ya hasilnya adalah 1
	// ini juga berfungsi sebegai penanda jika value nya minimal 1, maka recursive harus berhenti
	if value == 1 {
		return 1
	} else {
		// kita lakukan recursive
		fmt.Println(value)
		return value * factorialRecursive(value-1)
	}
}

func main() {

	// kita panggil factorialLoop, dan masukkan parameternya
	// misal 5, artinya kita menghitung nilai dari 5!
	// result := factorialLoop(5)
	result := factorialRecursive(5)
	fmt.Println(result)
}
