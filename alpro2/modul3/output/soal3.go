package main

import (
	"fmt"
)

func main() {
	var t1, t2, r1 int //lingkaran1

	var t3, t4, r2 int //lingkaran2

	var titikx, titiky int

	fmt.Scan(&t1, &t2, &r1)
	fmt.Scan(&t3, &t4, &r2)
	fmt.Scan(&titikx, &titiky)

	l1 := ceklingkaran(t1, t2, r1, titikx, titiky)
	l2 := ceklingkaran(t3, t4, r2, titikx, titiky)

	if l1 && l2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if l1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if l2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}

}

func ceklingkaran(a, b, r, x, y int) bool {
	A := x - a
	B := y - b

	return A*A+B*B <= r*r
}
