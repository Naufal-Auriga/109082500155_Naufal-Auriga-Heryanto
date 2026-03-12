package main

import "fmt"

func main() {
	var b1, b2, b3, b4 string

	var succes bool
	succes = true

	for i := 0; i <= 5; i++ {
		fmt.Printf("percobaan %d masukin 4 warna : \n", i)
		fmt.Scan(&b1, &b2, &b3, &b4)
		succes = succes && (b1 == "merah" && b2 == "kuning" && b3 == "ijo" && b4 == "ungu")
	}

	fmt.Print("BERHASIL : ", succes)
}
