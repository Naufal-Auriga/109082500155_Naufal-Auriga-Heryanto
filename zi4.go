package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	hitungX := 0
	hitung0 := 0

	for i := 1; i <= 9; i++ {
		var angka int
		fmt.Scan(&angka)

		if angka == x {
			hitungX++
		} else {
			hitung0++
		}
	}

	if hitungX > hitung0 {
		fmt.Println("Modus =", x)
	} else {
		fmt.Println("Modus = 0")
	}
}
