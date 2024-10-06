package main

import "fmt"

func main() {
	type noKtp string

	var ktp noKtp = "1232323232"

	fmt.Println(ktp)
}
