package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (userSlice UserSlice) Len() int {
	return len(userSlice)
}

func (userSlice UserSlice) Less(i, j int) bool {
	return userSlice[i].Age < userSlice[j].Age
}

func (userSlice UserSlice) Swap(i, j int) {
	userSlice[i], userSlice[j] = userSlice[j], userSlice[i]
}

func main() {

	users := []User{
		{"Bayu", 25},
		{"Agung", 20},
		{"Aan", 26},
		{"Bagus", 25},
	}

	// konversi dari users ke UserSlice
	sort.Sort(UserSlice(users))
	fmt.Println(users)

}
