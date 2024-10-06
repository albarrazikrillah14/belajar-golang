package main

import "fmt"

func main() {
	var nilai32 int32 = 33243
	var nilai64 int64 = int64(nilai32)
	//overflow
	var nilai16 int16 = int16(nilai64)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	var name = "Albarra Zikrillah"
	var r = name[4]
	var rString = string(r)

	fmt.Println(r)
	fmt.Println(rString)
}
