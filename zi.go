package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	agenap := a%2 == 0
	bgenap := b%2 == 0

	aganjil := a%2 != 0
	bganjil := b%2 != 0

	if (agenap && bganjil) || (aganjil && bgenap) {
		fmt.Println("berhasil")
	}

}
