package main

import "fmt"

func main() {
	var parsel, berat, biaya, total, sisa, ongkir int

	fmt.Print("masukaan parsel :")
	fmt.Scan(&parsel)

	berat = parsel / 1000

	sisa = parsel % 1000

	ongkir = berat * 10000

	fmt.Printf("maka %d kg + %d gram\n", berat, sisa)

	if parsel >= 10000 {
		biaya = 0
	} else {
		if parsel >= 500 {
			biaya = sisa * 5
		} else {
			biaya = sisa * 15
		}
	}
	total = ongkir + biaya

	fmt.Println("ini toal semuanya rp", total)

}
