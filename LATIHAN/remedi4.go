package main

import "fmt"

func main() {
	var nilai, x int
	fmt.Print("Masukkan nilai : ")
	fmt.Scan(&nilai)
	fmt.Print("Masukkan x : ")
	fmt.Scan(&x)

	z := 0
	f := 1

	for i := 1; i <= x-1; i++ {
		f = f * 10
	}

	for j := 1; j <= x; j++ {
		m := nilai / f
		nilai = nilai % f
		f = f / 10

		n := 1
		for k := 1; k <= j; k++ {
			n = n * m
		}

		fmt.Print(n, " ")
		z = z + n
	}

	fmt.Println()
	fmt.Println(z)
}
