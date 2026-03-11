package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	pembagi := 0

	if n > 1 {
		for i := 1; i <= n; i++ {
			if n%i == 0 {
				pembagi = pembagi + 1
			}
		}
	}

	if pembagi == 2 {
		fmt.Println("prima")
	} else {
		fmt.Println("bukan prima")
	}

}
