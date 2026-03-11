package main

import "fmt"

func main() {
	var tanaman string

	fmt.Print("masukan nama tanaman : ")
	fmt.Scan(&tanaman)

	switch tanaman {
	case "nepenthus", "drosera":
		fmt.Println("tanamn karnivora dari indo")
	case "venus", "sarrencia":
		fmt.Println("tanamn karnivora bukan dari indo")
	default:
		fmt.Println("bukan tanamn karnivora ")
	}
}
