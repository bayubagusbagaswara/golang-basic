package main

import "fmt"

// bikin type declarations untuk function Blacklist
// dimana parameter function adalah string, dan return value nya adalah boolean
type Blacklist func(string) bool

// bikin function dimana paramternya adalah function lain
func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You Are Blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {

	// ANONYMOUS FUNCTION 1 disimpan di sebuah variable
	blacklist := func(name string) bool {
		return name == "admin"
	}

	// panggil function registerUser
	registerUser("admin", blacklist)
	registerUser("bayu", blacklist)

	// ANONYMOUS FUNCTION langsung di parameter function
	registerUser("root", func(name string) bool {
		return name == "root"
	})

	registerUser("bayu", func(name string) bool {
		return name == "root"
	})

}
