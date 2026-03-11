package main

import "fmt"

func main() {
	for {
		var kiri, kanan float64

		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		fmt.Scan(&kiri, &kanan)

		if kiri < 0 || kanan < 0 || kiri+kanan > 150 {
			fmt.Println("Proses selesai.")
			break
		}

		var selisih float64
		if kiri > kanan {
			selisih = kiri - kanan
		} else {
			selisih = kanan - kiri
		}

		if selisih >= 9 {
			fmt.Println("Sepeda motor pak Andi akan oleng: true")
		} else {
			fmt.Println("Sepeda motor pak Andi akan oleng: false")
		}
	}
}
