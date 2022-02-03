package main

import "fmt"

func main() {

	// BREAK
	// bikin perulangan sampai 10 kali
	// ketika index bernilai 5, maka hentikan perulangannya atau keluar perulangan
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}

		// jadi hanya cetak index0 sampai index4
		fmt.Println("Perulangan ke", i)
	}

	fmt.Println("=================")
	// continue adalah menghentikan perulangan saat ini
	// lalu menlanjutkan ke perulangan selanjutnya, atau bisa dikatakan skip perulangan

	// CONTINUE
	// misal kita akan skip perulangan jika nilainya adalah genap
	// jadi hanya akan menyelesaikan perulangan bagi index ganjil saja
	for i := 0; i < 10; i++ {

		if i%2 == 0 {
			continue
		}
		fmt.Println("Perulangan ke", i)

	}

}
