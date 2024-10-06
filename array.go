package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Mr"
	names[1] = "Albarra"
	names[2] = "Zikrillah"

	fmt.Print(names[0])
	fmt.Print(names[1])
	fmt.Print(names[2])

	var values = [...]int{
		1,
		2,
		3,
		4,
		5,
		6,
	}

	fmt.Println(values)

	fmt.Println("panjang array", len(values))

}
