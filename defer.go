package main

import "fmt"

func main() {
	runApplication()
}

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication() {
	defer logging()
	fmt.Println("run aplication")
}
