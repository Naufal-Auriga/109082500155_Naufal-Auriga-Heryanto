package main

import "fmt"

func main() {
	var x, a, b int

	fmt.Scan(&x, &a, &b)
	fmt.Printf("ini fogoh %d = %d\n", x, (fogoh(x)))
	fmt.Printf("ini hofog %d = %d\n", x, (hofog(x)))
	fmt.Printf("ini gohof %d = %d\n", x, (gohof(x)))

	fmt.Printf("ini fogoh %d = %d\n", a, (fogoh(a)))
	fmt.Printf("ini hofog %d = %d\n", a, (hofog(a)))
	fmt.Printf("ini gohof %d = %d\n", a, (gohof(a)))

	fmt.Printf("ini fogoh %d = %d\n", b, (fogoh(b)))
	fmt.Printf("ini hofog %d = %d\n", b, (hofog(b)))
	fmt.Printf("ini gohof %d = %d\n", b, (gohof(b)))

}

func f(x int) int {
	return x * x
}

func g(x int) int {
	return x - 2
}

func h(x int) int {
	return x + 1

}

func fogoh(x int) int {
	result := f(g(h(x)))

	return result
}

func hofog(x int) int {
	result := h(f(g(x)))
	return result
}

func gohof(x int) int {
	result := g(h(f(x)))
	return result
}
