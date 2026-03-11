package main

import "fmt"

func main() {
	var aku int
	fmt.Print("masukan angka :")
	fmt.Scan(&aku)
	for i := 1; i <= aku; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}

	}
}
