package main

import "fmt"

func main() {
	var name string

	name = "Albarra Zikrillah"
	fmt.Println(name)

	name = "Zikrillah Albarra"
	fmt.Println(name)

	var hobby = "Basket"
	fmt.Println(hobby)

	//tanpa var, tapi hanya pertama kali
	alamat := "Dramaga, Bogor"
	fmt.Println(alamat)

	var (
		firstName = "Albarra"
		lastName  = "Zikrillah"
	)

	fmt.Println(firstName, lastName)

	const pi = 3.14

	fmt.Println(pi * 10 * 10)
}
