package main

import "fmt"

func main() {
	var angka int
	fmt.Scan(&angka)

	jumlah := 0

	for angka >= 0 {
		fmt,Scan(&angka)
		jumlah = jumlah + angka
	}

	fmt.Println(jumlah)
}
