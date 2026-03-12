package main

import "fmt"

func main() {

	var berat, kg, gram, biaya int

	fmt.Print("masukan berat: ")
	fmt.Scan(&berat)

	kg = berat / 1000
	gram = berat % 1000

	if kg >= 10 {
		biaya = kg * 10000
	} else if kg <= 10 && gram >= 500 {
		biaya = kg*10000 + gram*5

	} else if kg <= 10 && gram <= 500 {
		biaya = kg*10000 + gram*15
	}
	fmt.Printf("berat %d kg + %d gram dan hasil totalnya ", kg, gram)
	fmt.Println(biaya)
}
