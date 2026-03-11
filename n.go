package main

import "fmt"

func main() {
	var n int

	fmt.Print("Masukkan angka: ")
	fmt.Scan(&n)

	switch {
	case n%2 == 0:
		fmt.Println("Genap dan kelipatan 10")
		fmt.Println(n * (n + 1))

		fmt.Println(n / 10)

	default:

		fmt.Println("Kelipatan 5 dan ganjil")
		fmt.Println(n * n)
		fmt.Println(n + (n + 1))
	}
}
