package main

import "fmt"

func main() {
	berhasil := true

	for i := 1; i <= 5; i++ {
		var a, b, c, d string
		fmt.Print("Percobaan ", i, ": ")
		fmt.Scan(&a, &b, &c, &d)

		if a != "merah" || b != "kuning" || c != "hijau" || d != "ungu" {
			berhasil = false
		}
	}

	fmt.Println("BERHASIL:", berhasil)
}
