package main

import (
	"fmt"
	"reflect"
)

// bikin struct
type Sample struct {
	Name string `required:"true" max:"10"`
}

func IsValid(data interface{}) bool {
	// jika valid maka returnya adalah true
	// ambil data reflect nya
	// lalu iterasi semua field nya
	// cek apakah fieldnya ada required == true
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			return reflect.ValueOf(data).Field(i).Interface() != ""
		}
	}
	return true
}

type ContohaLagi struct {
	Name        string
	Description string
}

func main() {

	// kita ingin membaca ada berapa field dan sebagainya dari struct Sample, yakni menggunakan reflect
	sample := Sample{"Bayu"}

	// bikin reflection nya dan tipe nya apa, misal tipe nya sample
	var sampleType reflect.Type = reflect.TypeOf(sample)

	// kita ambil field yang ke 0
	structField := sampleType.Field(0)

	// ambil tag
	required := sampleType.Field(0).Tag.Get("required")
	max := sampleType.Field(0).Tag.Get("max")

	fmt.Println(structField.Name)
	fmt.Println(required)
	fmt.Println(max)

	fmt.Println(IsValid(sample)) // true, ada field name
	sample.Name = ""
	fmt.Println(IsValid(sample)) // false, karena field name sudah kita set string kosong

	contohLagi := ContohaLagi{"Bayu", "Oke"}
	fmt.Println(IsValid(contohLagi)) // tetep true
}
