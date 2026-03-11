package main

import "fmt"

func main() {
	var tugas, kuis, ujian, praktikum int
	var hasil float64

	fmt.Print("masukan nilai tugas :")
	fmt.Scan(&tugas)
	fmt.Print("masukan nilai kuis :")
	fmt.Scan(&kuis)
	fmt.Print("masukan nilai ujian :")
	fmt.Scan(&ujian)
	fmt.Print("masukan nilai praktikum ;")
	fmt.Scan(&praktikum)

	hasil = 0.1*float64(tugas) + 0.2*float64(tugas) + 0.5*float64(ujian) + 0.2*float64(praktikum)

	fmt.Println(hasil)

}
