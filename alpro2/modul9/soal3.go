package main

import "fmt"

const NMAX = 100

type arrBalita [NMAX]float64

// Prosedur mencari berat minimum dan maksimum
func hitungMinMax(arr arrBalita, n int, bMin, bMax *float64) {
	*bMin = arr[0]
	*bMax = arr[0]

	for i := 1; i < n; i++ {
		if arr[i] < *bMin {
			*bMin = arr[i]
		}
		if arr[i] > *bMax {
			*bMax = arr[i]
		}
	}
}

// Fungsi menghitung rerata
func rerata(arr arrBalita, n int) float64 {
	var total float64

	for i := 0; i < n; i++ {
		total += arr[i]
	}

	return total / float64(n)
}

func main() {
	var berat arrBalita
	var n int
	var min, max float64

	fmt.Print("Masukan banyak data berat balita : ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("Masukan berat balita ke-%d : ", i+1)
		fmt.Scan(&berat[i])
	}

	hitungMinMax(berat, n, &min, &max)

	rata := rerata(berat, n)

	fmt.Printf("\nBerat balita minimum: %.2f kg\n", min)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", max)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rata)
}
