package main

import "fmt"

func main() {
	var jam int
	fmt.Scan(&jam)

	if jam == 0 {
		fmt.Println("12 AM")
	} else if jam < 12 {
		fmt.Println(jam, "AM")
	} else if jam == 12 {
		fmt.Println("12 PM")
	} else {
		fmt.Println(jam-12, "PM")
	}
}
