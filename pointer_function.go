package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func changeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	address := Address{}
	changeCountryToIndonesia(&address)
	fmt.Println(address.Country)

	var address2 *Address = &Address{}
	changeCountryToIndonesia(address2)

	fmt.Println(address2.Country)
}
