package main

import (
	"bufio"
	"fmt"
	"os"
)

const DiskonPersen = 0.35 // 35% diskon

func main() {
	// Contoh kasus
	fmt.Println("--- Contoh Kasus ---")
	fmt.Printf("Input: 10000, true  -> Output: %.0f\n", hitungBelanjaAkhir(10000, true))
	fmt.Printf("Input: 10000, false -> Output: %.0f\n", hitungBelanjaAkhir(10000, false))
	fmt.Printf("Input: 25100, true  -> Output: %.0f\n", hitungBelanjaAkhir(25100, true))

	// Input dari user
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\n--- Coba Input Anda ---")

	var totalBelanja float64
	var ikutAsesmen bool

	fmt.Print("Masukkan total belanja: ")
	fmt.Fscan(reader, &totalBelanja)

	fmt.Print("Apakah sedang asesmen (true/false): ")
	fmt.Fscan(reader, &ikutAsesmen)

	hasil := hitungBelanjaAkhir(totalBelanja, ikutAsesmen)
	fmt.Printf("Total belanja akhir: %.0f\n", hasil)
}
