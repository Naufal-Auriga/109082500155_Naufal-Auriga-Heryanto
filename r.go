package main

import "fmt"

func main() {

	var e, x, y, y1 float64

	e = 0.0000001
	x = 2.0
	y = 0.0
	y1 = x
	for y1-y > e || y1-y < -e {
		y = y1
		y1 = 0.5*y + 0.5*(x/y)
	}
	fmt.Printf("sqrt(%v)=%v\n", x, y)
}
