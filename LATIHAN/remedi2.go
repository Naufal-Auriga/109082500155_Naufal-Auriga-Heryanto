package main

import "fmt"

func main() {
	var tinggi_pohon, bayangan_tongkat, bayangan_pohon, tangga float64
	var bandingin bool

	fmt.Print("masukan tangga :")
	fmt.Scan(&tangga)
	fmt.Print("masukan bayangan_tongkat :")
	fmt.Scan(&bayangan_tongkat)
	fmt.Print("masukan tinggi_pohon :")
	fmt.Scan(&bayangan_pohon)

	tinggi_tongkat := 1.0
	tinggi_pohon = (bayangan_pohon * tinggi_tongkat) / bayangan_tongkat

	bandingin = tangga >= tinggi_pohon

	fmt.Println("tinggi pohon", tinggi_pohon, bandingin)

}
