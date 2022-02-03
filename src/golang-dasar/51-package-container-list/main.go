package main

import (
	"container/list"
	"fmt"
)

func main() {

	data := list.New()
	data.PushBack("Bayu")
	data.PushBack("Bagus")
	data.PushBack("Bagaswara")

	// iterasi ambil data
	// jika e tidak sama dengan nil, artinya kita akan asign e dengan data Next
	// sampai menemui nil
	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

}
