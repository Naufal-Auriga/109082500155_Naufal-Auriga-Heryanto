package main

import "fmt"

func main() {
	var target int
	fmt.Scan(&target)

	total := 0
	donatur := 0

	for total < target {
		var sumbangan int
		fmt.Scan(&sumbangan)

		donatur++
		total = total + sumbangan

		fmt.Printf(
			"Donatur %d: Menyumbang %d. Total terkumpul: %d\n",
			donatur, sumbangan, total,
		)
	}

	fmt.Printf(
		"Target tercapai! Total donasi: %d dari %d donatur.\n",
		total, donatur,
	)
}
