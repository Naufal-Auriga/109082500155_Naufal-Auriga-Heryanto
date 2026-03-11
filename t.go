package main

import "fmt"

func main() {
	var token string
	var i int
	var nominal int
	fmt.Print("masukan token :")
	fmt.Scan(&token)

	for i = 0; i < 3 && token != "123abc"; i++ {
		fmt.Print("token yang anda masukan salah masukan lagi :")
		fmt.Scan(&token)

	}
	if token == "123abc" {
		fmt.Print("masukan nominal : ")
		fmt.Scan(&nominal)
		fmt.Printf("witdral rupiah %d thank uu\n", nominal)
		fmt.Printf("Selamat Anda berhasil login dengan %d percobaan\n", i)
	} else {
		fmt.Printf("anda gagal dengan %d percobaan dan di block\n", i)
	}
}
