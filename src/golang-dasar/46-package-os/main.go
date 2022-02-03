package main

import (
	"fmt"
	"os"
)

func main() {

	args := os.Args
	fmt.Println("Argument: ", args)

	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println(hostname)
	} else {
		fmt.Println("Error", err.Error())
	}

	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")

	fmt.Println(username)
	fmt.Println(password)
}
