package main

import "fmt"

func main() {
	greeting := sayHello("Albarra Zikrillah")

	fmt.Println(greeting)

	firstname, lastname := getFullname("Albarra", "Zikrllah")

	fmt.Println(firstname, lastname)

	total := sumAll(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(total)

	goodBye := sayGoodBye

	fmt.Println(goodBye("Albarra"))
}

func sayHello(name string) string {
	return "Hello, " + name
}

func getFullname(firstname string, lastname string) (string, string) {
	return firstname, lastname
}

func sumAll(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func sayGoodBye(name string) string {
	return "Good Bye " + name
}
