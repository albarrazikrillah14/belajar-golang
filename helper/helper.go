package helper

var version = "1.0.0"
var Application = "golang"

func SayHello(name string) string {
	return "Hello, " + name
}

// fungsi private hanya bisa diakses pada module yang sama
func sayGoodbye(name string) string {
	return "Goodbye, " + name
}
