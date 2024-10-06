package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr " + man.Name
}

func main() {
	zikri := Man{"Albarra"}
	zikri.Married()

	fmt.Println(zikri)
}
