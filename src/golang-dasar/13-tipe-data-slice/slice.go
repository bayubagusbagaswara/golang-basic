package main

import "fmt"

func main() {

	// buat array dulu
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"Nopember",
		"Desember",
	}

	// lalu kita bikin slice yang pertama
	var slice1 = months[4:7]
	fmt.Println(slice1)      // [Mei, Juni, Juli], atau index [4, 5, 6], index 7 tidak ikut
	fmt.Println(len(slice1)) // len : 3
	fmt.Println(cap(slice1)) // cap : 8, dari Mei-Desember ada 8 data

	// kita bikin slice lagi, dari index ke-10 sampai habis
	var slice2 = months[10:]
	fmt.Println(slice2) // [November, Desember]

	// misal kita buat slice3, yang isinya adalah slice2 ditambah dengan data baru
	var slice3 = append(slice2, "Bayu")
	fmt.Println(slice3) // [November, Desember, Bayu]

	// kita ubah slice3, misal ubah index-1 milik slice3, yakni Desember
	slice3[1] = "Bukan Desember"
	// hasilnya tidak berubah

	fmt.Println("=============================")

	// var fruitsA = []string{"apple", "grape"}     // slice
	// var fruitsB = [2]string{"banana", "melon"}   // array
	// var fruitsC = [...]string{"papaya", "grape"} // array

	// misal kita punya slice dari array
	var fruits = []string{"apple", "grape", "banana", "melon"}

	// aFruits memiliki pointer yang mengacu ke index-0 pada array fruits,
	var aFruits = fruits[0:3]
	// bFruits memiliki pointer yang mengacu ke index-1 pada array fruits
	var bFruits = fruits[1:4]

	// kita slice lagi dari slice
	var aaFruits = aFruits[1:2]
	var baFruits = bFruits[0:1]

	fmt.Println(fruits)   // [apple grape banana melon]
	fmt.Println(aFruits)  // [apple grape banana]
	fmt.Println(bFruits)  // [grape banana melon]
	fmt.Println(aaFruits) // [grape]
	fmt.Println(baFruits) // [grape]

	// TAPI SAAT KITA UBAH DATA DI SLICE, MAKA JUGA AKAN MENGUBAH STRUKTUR DATA DI ARRAY NYA

	// Buah "grape" diubah menjadi "pinnaple"
	baFruits[0] = "pinnaple"

	fmt.Println(fruits)   // [apple pinnaple banana melon]
	fmt.Println(aFruits)  // [apple pinnaple banana]
	fmt.Println(bFruits)  // [pinnaple banana melon]
	fmt.Println(aaFruits) // [pinnaple]
	fmt.Println(baFruits) // [pinnaple]

	// FUNGSI LEN
	var fruits1 = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(len(fruits1)) // 4

	// FUNGSI CAP
	var fruits2 = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(len(fruits)) // len: 4
	fmt.Println(cap(fruits)) // cap: 4

	var aFruits2 = fruits2[0:3]
	fmt.Println(len(aFruits2)) // len: 3 yaitu [0,1,2]
	fmt.Println(cap(aFruits2)) // cap: 4 yaitu [0,1,2,3]

	var bFruits2 = fruits2[1:4]
	fmt.Println(len(bFruits2)) // len: 3 yaitu [1,2,3]
	fmt.Println(cap(bFruits2)) // cap: 4 yaitu [1,2,3]

	// FUNGSI APPEND
	var fruits3 = []string{"apple", "grape", "banana"}
	var bFruits3 = fruits3[0:2]

	fmt.Println(cap(bFruits3)) // 3
	fmt.Println(len(bFruits3)) // 2

	fmt.Println(fruits3)  // ["apple", "grape", "banana"]
	fmt.Println(bFruits3) // ["apple", "grape"]

	var cFruits3 = append(bFruits3, "papaya")

	fmt.Println(fruits3)  // ["apple", "grape", "papaya"]
	fmt.Println(bFruits3) // ["apple", "grape"]
	fmt.Println(cFruits3) // ["apple", "grape", "papaya"]

	// FUNGSI COPY
	dst := []string{"potato", "potato", "potato"}
	src := []string{"watermelon", "pinnaple"}
	n := copy(dst, src)

	fmt.Println(dst) // watermelon pinnaple potato
	fmt.Println(src) // watermelon pinnaple
	fmt.Println(n)   // 2

	// AKSES ELEMEN SLICE DENGAN 3 INDEX
	var fruits5 = []string{"apple", "grape", "banana"}
	var aFruits5 = fruits5[0:2]
	var bFruits5 = fruits5[0:2:2]

	fmt.Println(fruits5)      // ["apple", "grape", "banana"]
	fmt.Println(len(fruits5)) // len: 3
	fmt.Println(cap(fruits5)) // cap: 3

	fmt.Println(aFruits5)      // ["apple", "grape"]
	fmt.Println(len(aFruits5)) // len: 2
	fmt.Println(cap(aFruits5)) // cap: 3

	fmt.Println(bFruits5)      // ["apple", "grape"]
	fmt.Println(len(bFruits5)) // len: 2
	fmt.Println(cap(bFruits5)) // cap: 2

}
