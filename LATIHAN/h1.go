package main

import "fmt"

func main() {
	pita := ""
	jumlah := 0

	for {
		var bunga string
		fmt.Print("Bunga ", jumlah+1, ": ")
		fmt.Scan(&bunga)

		if bunga == "SELESAI" || bunga == "selesai" {
			break
		}

		pita = pita + bunga + " - "
		jumlah = jumlah + 1
	}

	fmt.Println("Pita:", pita)
	fmt.Println("Bunga:", jumlah)
}
