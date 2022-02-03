package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {

	// misal kita buat 5 data
	data := ring.New(5)

	for i := 0; i < data.Len(); i++ {
		data.Value = "Value " + strconv.FormatInt(int64(i), 10)
		// ambil data ring selanjutnya
		data = data.Next()
	}

	// Do untuk mengetahui jumlah data di ring
	data.Do(func(value interface{}) {
		// print semua element yang ada di dalam ring
		fmt.Println(value)
	})
}
