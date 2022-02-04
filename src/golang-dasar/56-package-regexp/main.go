package main

import (
	"fmt"
	"regexp"
)

func main() {

	var regex = regexp.MustCompile(`b([a-z])([a-z])u`)

	fmt.Println(regex.MatchString("bayu"))
	fmt.Println(regex.MatchString("batu"))
	fmt.Println(regex.MatchString("baYu"))

	fmt.Println(regex.FindAllString("bayu batu baru baju b10 biru", 10))
}
