package main

import "fmt"

type bahan struct {
	w1, w2, w3, w4 string
}

func input() bahan {
	var b bahan
	fmt.Scan(&b.w1, &b.w2, &b.w3, &b.w4)

	return b
}

func cek(b bahan) bool {
	return b.w1 == "merah" && b.w2 == "kuning" && b.w3 == "hijau" && b.w4 == "ungu"
}

func main() {
	data := input()

	fmt.Println("sukses: ", cek(data))
}
