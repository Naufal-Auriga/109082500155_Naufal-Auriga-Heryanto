package main

import "fmt"

func main() {
	var a, b, c, d int

	fmt.Scan(&a, &b, &c, &d)

	if a >= c {
		fmt.Println(permutasi(a, c))
		fmt.Println(kombinasi(a, c))
	} else {
		fmt.Println(permutasi(c, a))
		fmt.Println(kombinasi(c, a))
	}

	if b >= d {
		fmt.Println(permutasi(b, d))
		fmt.Println(kombinasi(b, d))

	} else {
		fmt.Println(permutasi(d, b))
		fmt.Println(kombinasi(d, b))

	}
}

func permutasi(n int, r int) int {
	return faktorial(n) / faktorial(n-r)
}

func kombinasi(n int, r int) int {
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}

func faktorial(n int) int {
	hasil := 1

	for i := 1; i <= n; i++ {
		hasil = hasil * i

	}
	return hasil
}
