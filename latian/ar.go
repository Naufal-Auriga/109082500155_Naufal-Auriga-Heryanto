package main

import "fmt"

func main() {

	arr := []int{4, 7, 2, 9, 6}

	for i := 1; i < len(arr); i++ {
		if arr[i]%2 == 0 {
			fmt.Println("genap broo", arr[i])
		}

	}
}
