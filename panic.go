package main

import "fmt"

func main() {
	runApp(true)

}

func endApp() {
	fmt.Println("End app")
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("Ups Error")
	}
}
