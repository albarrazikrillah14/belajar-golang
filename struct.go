package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "my name is", customer.Name)
}

func main() {
	var person Customer
	person.Name = "Albarra Zikrillah"
	person.Address = "Bogor"
	person.Age = 34

	fmt.Println(person)
	fmt.Println(person.Name)

	person1 := Customer{
		Name:    "Albarra Zikrillah",
		Address: "Bogor",
		Age:     252,
	}

	person2 := Customer{
		"Albarra Zikrillah",
		"Bogor",
		40,
	}

	fmt.Println(person1)
	fmt.Println(person2)

	person2.sayHello("sumbul")
}
