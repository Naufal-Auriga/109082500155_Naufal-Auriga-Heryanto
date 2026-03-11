package main

import "fmt"

func loss(i int) int {
	return i + 3

}

func main() {
	for i := 0; i <= 10; i++ {
		progress := i * 6
		os := loss(i)

		if i%3 == 0 {
			fmt.Printf("%d bingo\n", i)

		} else if i < 2 {
			fmt.Printf("%d train achieve %d%% accuracy but %d%% loss\n", i, progress, os)

		} else {
			fmt.Printf("%d Trains achieve %d%% accuracy but %d%% loss\n", i, progress, os)
		}

	}
}
