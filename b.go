package main

import "fmt"

func main() {
	var a int
	var hasil bool

	fmt.Print("masukan angka:")
	fmt.Scan(&a)

	hasil = a < 0 && a%2 == 0

	fmt.Println(hasil)
}
