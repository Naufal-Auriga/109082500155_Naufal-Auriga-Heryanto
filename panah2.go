package main

import "fmt"

func main() {
	var skor int

	total := 0
	panah := 0

	for skor >= 0 && skor <= 3 {
		fmt.Scan(&skor)

		if skor == 0 || skor == 1 || skor == 2 || skor == 3 {
			total = total + skor

			panah++
		}

	}
	fmt.Println(total)
	fmt.Println(panah)
}
