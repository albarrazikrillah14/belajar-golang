package main

import "fmt"

func random() any {
	return "Ok"
}

func main() {
	var result any = random()

	switch value := result.(type) {
	case string:
		fmt.Println("string", value)
	case int:
		fmt.Println("int", value)
	default:
		fmt.Println("value", value)
	}
}
