package main // paket utama

import (
	"fmt"  // library input dan output
	"math" // library untuk fungsi pembulatan ke atas
)

func main() { // fungsi utama
	var angka float64 // variabel untuk menyimpan angka desimal
	fmt.Scan(&angka)  // membaca input user

	batas := math.Floor(1.0) // menentukan pembulatan ke atas
	nilai := angka           // nilai awal dimulai dari angka input

	// loop selama nilai masih kurang dari batas
	for nilai < batas {
		nilai = nilai + 0.1         // menambahkan 0.1 setiap loop
		fmt.Printf("%.1f\n", nilai) // menampilkan nilai dengan 1 angka desimal
	}
}
