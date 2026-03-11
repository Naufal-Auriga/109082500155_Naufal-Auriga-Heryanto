package main

import "fmt"

func main() {
	var b1, b2, b3, max, min int

	fmt.Scan(&b1, &b2, &b3)
	if b1 > b2 {
		max = b1
		min = b2
	} else {
		max = b2
		min = b1
	}

	if b3 > max {
		max = b3
	}
	if b3 < min {
		min = b3
	}
	fmt.Println("Maksimal", max)
	fmt.Println("Minimum", min)
}
