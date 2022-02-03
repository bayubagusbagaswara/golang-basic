package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication() {

	// deklarasi function defer logging
	defer logging()

	// artinya setelah kita mengeksekusi function runApplication(), maka diakhir akan mengeksekusi function logging()
	fmt.Println("Run Application")
}

func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("Error dengan message:", message)
	}
	fmt.Println("Aplikasi Selesai")
}

func runApp(error bool) {
	// defer function akan tetap dieksekusi meksi terjadi panic dari kode program bawah
	defer endApp()

	// jika terjadi error, maka eksekusi panic
	if error {
		panic("Aplikasi Error")
	}
	// setelah panic maka akan dilempar keatas, sehingga kode program dibawah tidak akan dieksekusi
	// Atau biasa disebut Throw Exception
	// oleh itu, kita taruhr recover diluar function, yakni di function endApp()
	// dimana endApp selalu dieksekusi diakhir (DEFER)

	fmt.Println("Aplikasi Berjalan")
}

func main() {

	// eksekusi function runApplication()
	// runApplication()

	// eksekusi runApp()
	runApp(true)
	fmt.Println("Aplikasi masih tetap berjalan")
}
