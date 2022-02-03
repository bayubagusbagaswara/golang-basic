package main

import (
	"fmt"
	"strconv"
)

func main() {
	boolean, err := strconv.ParseBool("true")

	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println("Error", err.Error())
	}

	number, err := strconv.ParseInt("1000000", 10, 64)

	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err.Error())
	}

	// jadi okta
	value := strconv.FormatInt(1000000, 8)
	fmt.Println(value)

	// ubah string ke int
	valueInt, _ := strconv.Atoi("20000000")
	fmt.Println(valueInt)
}
