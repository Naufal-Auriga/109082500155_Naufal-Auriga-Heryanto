package main

import (
	"fmt"
)

func main() {
	var angka float64
	var batas float64

	fmt.Scan(&angka, &batas)

	nilai := angka

	for nilai < batas {
		nilai = nilai + 0.1
		fmt.Printf("%.1f\n", nilai)
	}
}
