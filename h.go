package main

import "fmt"

func main() {
	var fak, i, jumlah int

	fmt.Print("masukin angka ;")
	fmt.Scan(&fak)
	jumlah = 0

	for i = 1; i <= fak; i++ {
		if fak%i == 0 {
			fmt.Print(" ", i)

			jumlah++

		}

	}
	fmt.Println()

	if jumlah == 2 {
		fmt.Print("true")
	} else {
		fmt.Print("false")
	}

}
