package main

import "fmt"

func area(z int) int {
	thearea := z * 4
	return thearea
}
func main() {

	hth := 8
	vol := hth * area(2)
	fmt.Println(vol)

}
