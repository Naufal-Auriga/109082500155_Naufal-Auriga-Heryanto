package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		progress := i * 6

		if i%3 == 0 {
			fmt.Printf("%d bingo\n", i)

		} else if i < 2 {
			fmt.Printf("%d train achieve %d%% accuracy\n", i, progress)

		} else {
			fmt.Printf("%d Trains achieve %d%% accuracy\n", i, progress)
		}

	}
}
