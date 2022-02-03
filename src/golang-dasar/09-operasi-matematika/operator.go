package main

import "fmt"

func main() {

	var a = 10
	var b = 10
	var c = a + b
	fmt.Println(c)

	var i = 10
	i += 10 // i = i + 10
	fmt.Println(i)

	i++ // i = i + 1
	fmt.Println(i)

	negative := -100
	positive := +100
	fmt.Println(negative)
	fmt.Println(positive)
}
