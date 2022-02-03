package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total = total + number
	}
	return total
}

func main() {
	total := sumAll(10, 20, 30, 40, 50)
	fmt.Println(total)

	// kita punya variable slice
	numbers := []int{60, 70, 80, 90, 100}

	// lalu kita ingin memanggil function sumAll, dimana dia adalah variadic function
	// maka cukup tuliskan nama silce diikuti 3 titik
	// artinya parameter berupa slice ini dijadikan variadic argument parameter
	total = sumAll(numbers...)
	fmt.Println(total)
}
