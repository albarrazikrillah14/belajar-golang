package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := &address1

	address2.City = "Bogor"

	fmt.Println(address1) //{Bogor Jawa Barat Indonesia}
	fmt.Println(address2) //&{Bogor Jawa Barat Indonesia}

	*address2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	fmt.Println(address1) //{Jakarta DKI Jakarta Indonesia}
	fmt.Println(address2) //&{Jakarta DKI Jakarta Indonesia}
}
