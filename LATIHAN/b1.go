package main

import "fmt"

func main() {
	var angka int
	fmt.Scan(&angka)

	jumlahDigit := 0

	for angka > 0 {
		jumlahDigit++
		angka = angka / 10
	}

	fmt.Println(jumlahDigit)
}
