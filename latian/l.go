package main

import "fmt"

func main() {
	var aku int

	fmt.Print("masukan angka :")
	fmt.Scan(&aku)

	if aku%2 == 0 {

		fmt.Printf("ini bilngan genap %d", aku)
	} else {
		fmt.Printf(" %d bukan genap tapi ganjil", aku)
	}
}
