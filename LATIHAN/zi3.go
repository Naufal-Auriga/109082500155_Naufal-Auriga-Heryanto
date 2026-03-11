package main

import "fmt"

func main() {
	var P, R int
	fmt.Scan(&P, &R)

	hari := P / R

	if P%R != 0 {
		hari = hari + 1
	}

	fmt.Println(hari, "hari")
}
