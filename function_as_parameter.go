package main

import "fmt"

func main() {
	sayHello("Albarra Zikrillah", spamFilter)
	sayHello("anjing", spamFilter)
}

type Filter func(string) string

func sayHello(name string, filter Filter) {
	fmt.Println("Hello", filter(name))

}

func spamFilter(name string) string {
	if name == "anjing" {
		return "*****"
	}

	return name
}
