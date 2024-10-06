package main

import "fmt"

type HasName interface {
	GetName() string
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

func main() {
	person := Person{"Albarra Zikrillah"}
	sayHello(person)

}

func sayHello(value HasName) {
	fmt.Println("Hello", value.GetName())
}
