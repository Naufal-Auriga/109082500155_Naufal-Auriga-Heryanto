package main

import "fmt"

func main() {

	var a, b int

	fmt.Scan(&a, &b)

	for (a+b)%2 == 0 {
		fmt.Println("hasil jumlah", a+b)
		fmt.Scan(&a, &b)
	}

	fmt.Println("end program ")
}
