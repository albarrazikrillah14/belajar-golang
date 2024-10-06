package main

import "fmt"

func main() {
	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}

	daySlice1 := days[5:]
	fmt.Println(daySlice1)

	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"

	fmt.Println(daySlice1)
	fmt.Println(days)

	daySlice2 := append(daySlice1, "Libur")

	fmt.Println(daySlice1)
	fmt.Println(daySlice2)

	daySlice2[0] = "Belajar"

	fmt.Println(daySlice1)
	fmt.Println(daySlice2)
}
