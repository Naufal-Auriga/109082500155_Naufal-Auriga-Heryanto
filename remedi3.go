package main

import "fmt"

func main() {
	var x, y, digit, temp int
	var status bool
	var total, j int

	fmt.Print("masukan x : ")
	fmt.Scan(&x)
	temp = x

	fmt.Print("masukan y : ")
	fmt.Scan(&y)

	total = 0

	for j = 1; j <= y; j++ {
		digit = x % 10
		
		pangkat := 1
		for i := 1; i <= y; i++ {
			pangkat = pangkat * digit
		}
		total = total + pangkat
		x = x / 10
	}
	status = total == temp

	fmt.Println(total, status)

}
