package main

import "fmt"

func main() {
	var aku int
	hasil := 0

	fmt.Print("Masukan angka: ")
	fmt.Scan(&aku)

	for aku > 0 {
		digit := aku % 10
		hasil = hasil*10 + digit
		aku = aku / 10

	}

	fmt.Println("Jumlah digit =", hasil)
}
