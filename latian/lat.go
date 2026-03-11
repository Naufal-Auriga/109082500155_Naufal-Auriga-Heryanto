package main

import "fmt"

func main() {
	var aku int
	total := 0

	fmt.Print("Masukan angka: ")
	fmt.Scan(&aku)

	for aku > 0 {
		digit := aku % 10
		total = total + digit
		aku = aku / 10
	}

	fmt.Println("Jumlah digit =", total)
}
