package main

import (
	"45-package-initialization/database"
	"fmt"
)

func main() {
	result := database.GetDatabase()

	fmt.Println(result)

}
