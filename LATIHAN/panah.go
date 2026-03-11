package main

import "fmt"

func main() {
	var skor int
	fmt.Scan(&skor)

	total := 0
	panah := 0

	for skor < 15 {

		fmt.Scan(&skor)

		if skor == 0 || skor == 1 || skor == 2 || skor == 3 {
			total += skor

			panah++
		}

	}
	fmt.Println(total)
	fmt.Println(panah)
}
