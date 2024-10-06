package main

import "fmt"

func main() {
	registerUser("Albarra Zikrillah", func(name string) bool {
		return name == "anjing"
	})

	blacklist := func(name string) bool {
		return name == "anjing"
	}

	registerUser("anjing", blacklist)
}

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}
