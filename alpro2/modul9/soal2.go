package main

import "fmt"

const NMAX = 1000

type arrBerat [NMAX]float64

func main() {
	var ikan arrBerat
	var x, y int

	fmt.Print("Masukkan jumlah ikan (x): ")
	fmt.Scan(&x)

	fmt.Print("Masukkan kapasitas tiap wadah (y): ")
	fmt.Scan(&y)

	for i := 0; i < x; i++ {
		fmt.Printf("Masukkan berat ikan ke-%d: ", i+1)
		fmt.Scan(&ikan[i])
	}

	jumlahWadah := (x + y - 1) / y

	fmt.Println("\nTotal berat tiap wadah:")

	for i := 0; i < jumlahWadah; i++ {
		var total float64
		var banyak int

		for j := i * y; j < (i+1)*y && j < x; j++ {
			total += ikan[j]
			banyak++
		}

		fmt.Printf("Wadah %d = %.2f kg\n", i+1, total)
		fmt.Printf("Rata-rata wadah %d = %.2f kg\n", i+1, total/float64(banyak))
	}
}
