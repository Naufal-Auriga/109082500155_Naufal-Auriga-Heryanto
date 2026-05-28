package main

import "fmt"

func selectionSort1(arr []int) {

	var i, j, idxMin int
	var temp int

	for i = 0; i < len(arr)-1; i++ {

		idxMin = i

		for j = i + 1; j < len(arr); j++ {

			if arr[idxMin] < arr[j] {

				idxMin = j
			}
		}

		temp = arr[idxMin]
		arr[idxMin] = arr[i]
		arr[i] = temp
	}

	return
}

func selectionSort(arr []int) {

	var i, j, idxMin int
	var temp int

	for i = 0; i < len(arr)-1; i++ {

		idxMin = i

		for j = i + 1; j < len(arr); j++ {

			if arr[idxMin] > arr[j] {

				idxMin = j
			}
		}

		temp = arr[idxMin]
		arr[idxMin] = arr[i]
		arr[i] = temp
	}

	return
}

func main() {
	var semua [][]int
	var n, m int
	fmt.Scan(&n)

	fmt.Scan(&m)
	for i := 0; i < n; i++ {

		rumah := make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&rumah[j])
		}
		semua = append(semua, rumah)

		ganjil := []int{}
		genap := []int{}

		for j := 0; j < len(rumah); j++ {

			if rumah[j]%2 == 0 {

				genap = append(genap, rumah[j])

			} else {

				ganjil = append(ganjil, rumah[j])
			}
		}

		selectionSort(ganjil)
		selectionSort1(genap)

		for j := 0; j < len(ganjil); j++ {
			fmt.Print(ganjil[j], " ")
		}

		for j := 0; j < len(genap); j++ {
			fmt.Print(genap[j], " ")
		}
	}

}
