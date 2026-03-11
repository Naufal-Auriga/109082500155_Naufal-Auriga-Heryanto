package main

import "fmt"

func main() {
	var total int
	var diskon int

	fmt.Scan(&total)

	if total > 1000000 {
		diskon = total * 20 / 100
		if diskon > 2000000 {
			diskon = 2000000
		}
	} else if total <= 100000 {
		diskon = total * 5 / 100
	} else {
		diskon = total * 15 / 100
	}

	bayar := total - diskon

	fmt.Println(diskon)
	fmt.Println(bayar)
}
