package main

import "fmt"

func main() {
	var aku int
	fmt.Print("masukan angka :")
	fmt.Scan(&aku)
	a := 0
	b := 1
	for i := 1; i <= aku; i++ {
		fmt.Println(a)
		c := a + b

		a = b
		b = c

	}
}
