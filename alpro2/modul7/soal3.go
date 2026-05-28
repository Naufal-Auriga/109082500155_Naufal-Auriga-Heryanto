package main

import "fmt"

func insertionSort(arr []int) {

	var i, j, temp int

	for i = 1; i < len(arr); i++ {

		temp = arr[i]
		j = i - 1

		for j >= 0 && arr[j] > temp {

			arr[j+1] = arr[j]
			j--
		}

		arr[j+1] = temp
	}
}

func main() {

	var x int
	data := []int{}

	for {

		fmt.Scan(&x)

		if x == -5313 {
			break
		}

		if x != 0 {

			data = append(data, x)

		} else {

			insertionSort(data)

			n := len(data)

			if n%2 == 1 {

				fmt.Println(data[n/2])

			} else {

				kiri := data[(n/2)-1]
				kanan := data[n/2]

				fmt.Println((kiri + kanan) / 2)
			}
		}
	}
}
