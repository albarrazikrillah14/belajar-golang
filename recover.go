package main

import "fmt"

func main() {
	runApp(true)
	fmt.Println("Aplikasi tetap berjalan")
}

func endApp() {
	message := recover()
	fmt.Println("Terjadi panic", message)
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("ERROR")
	}
}
