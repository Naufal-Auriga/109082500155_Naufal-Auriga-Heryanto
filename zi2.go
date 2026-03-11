package main

import "fmt"

func main() {
	var kamar, main, kebun int
	fmt.Scan(&kamar, &main, &kebun)

	if kamar >= 60 && main >= 75 && kebun >= 60 {
		fmt.Println("Ice Cream")
	} else if kamar >= 80 && kebun >= 80 {
		fmt.Println("Ice Cream")
	} else if kamar == 100 {
		fmt.Println("Ice Cream")
	} else {
		fmt.Println("Tidak")
	}
}
