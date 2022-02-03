package helper

var version = "1.0.0"      // tidak bisa diakses dari luar package
var Application = "golang" // bisa diakses dari luar pakcage

// tidak bisa diakses dari luar package
func sayGoodBye(name string) string {
	return "Hello " + name
}

// bisa diakses dari luar package
func SayHello(name string) string {
	return "Hello " + name
}
