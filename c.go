package main

import "fmt"

func main() {
	// Deklarasi variabel
	var N int                    // Jumlah mahasiswa (Input)
	const kapasitasMobil int = 7 // Kapasitas setiap mobil

	// Membaca input N dari pengguna
	fmt.Print("Masukkan jumlah mahasiswa (N): ")
	// fmt.Scanln digunakan untuk membaca input integer
	_, err := fmt.Scanln(&N)

	// Pengecekan input
	if err != nil || N <= 0 {
		fmt.Println("Input tidak valid. Masukkan bilangan bulat positif.")
		return // Keluar dari program
	}

	/*
	 * Logika Pembulatan Ke Atas (Ceiling)
	 * Untuk mendapatkan nilai (N / Kapasitas) yang dibulatkan ke atas,
	 * kita gunakan rumus pembagian integer: (N + Kapasitas - 1) / Kapasitas
	 */
	jumlahMobil := (N + kapasitasMobil - 1) / kapasitasMobil

	// Menampilkan hasil
	fmt.Printf("Jumlah mobil yang diperlukan untuk %d mahasiswa adalah: %d\n", N, jumlahMobil)
}
