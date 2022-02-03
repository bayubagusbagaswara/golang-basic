package helper

// bisa bikin function
func SayHello(name string) string {
	return "Hello " + name
}

// jika sama dalam satu package, maka tidak boleh membuat functin atau struct dengan nama yang sama
// misal function SayHello, ini tidak bisa
// karena SayHello sudah ada di package helper
// tapi kalau package nya berbeda, maka itu diperbolehkan
