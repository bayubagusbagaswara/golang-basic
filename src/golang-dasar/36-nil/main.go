package main

import "fmt"

// bikin function NewMap, dimana return valuenya adalah map
func NewMap(name string) map[string]string {

	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {

	var data map[string]string = NewMap("Bayu")

	// data nil, biasanya kita gunakan untuk pengecekan
	if data == nil {
		fmt.Println("Data Kosong")
	} else {
		fmt.Println(data)
	}
}
