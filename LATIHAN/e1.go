package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	ganjil := 0

	for i := 1; i <= n; i++ {
		if i%2 == 1 {
			ganjil = ganjil + 1
		}
	}
	fmt.Println("terdapat", ganjil, "bil ganj")
}
