package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Print("masukan angka: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			hasil := i * j

			fmt.Print(hasil)
		}
		fmt.Println()
	}
}
