package main

import "fmt"

func main() {
	var nilai int
	var ada bool
	var hadir float64
	var index string

	nilai = 65
	ada = true
	hadir = 0.0

	if nilai > 75 && ada {
		index = "A"
	} else if nilai > 64 {
		index = "B"
	} else if nilai > 50 && hadir > 0.7 {
		index = "C"
	} else {
		index = "E"
	}
	fmt.Printf("nilai %v, dan presentase kehadiran %.2f, dan buat tubes %v mendapat %v", nilai, hadir, ada, index)
}
