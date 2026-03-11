package main

import "fmt"

func main() {
	var kendaraan string
	var tarif int
	var durasi int

	fmt.Print("kendaraan : ")
	fmt.Scan(&kendaraan)

	fmt.Print("durasi: ")
	fmt.Scan(&durasi)

	switch {
	case kendaraan == "motor" && durasi >= 1 && durasi <= 2:
		tarif = 2000
	case kendaraan == "motor" && durasi > 2:
		tarif = 2000 * durasi
	case kendaraan == "mobil" && durasi >= 1 && durasi <= 2:
		tarif = 5000
	case kendaraan == "mobil" && durasi > 2:
		tarif = 5000 * durasi
	case (kendaraan == "truk" || kendaraan == "truck") && durasi >= 1 && durasi <= 2:
		tarif = 20000
	case (kendaraan == "truk" || kendaraan == "truck") && durasi > 2:
		tarif = 20000 * durasi

	default:
		fmt.Println("bukan kendaraan")
	}

	fmt.Printf("maka tarif anda %d", tarif)
}
