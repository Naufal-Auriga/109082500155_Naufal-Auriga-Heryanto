package main
import "fmt"

func main() {
	var (
		maxf, f0, f1, f2 int = 100, 0, 1, 1
		selesai bool
	)

	fmt.Println("Bilangan Pertama: ", f1)
	for selesai = false; !selesai; {
		f0 = f1
		f1 = f2
		f2 = f1 + f0
		fmt.Println("Bilangan Berikutnya: ", f1)
		selesai = f2 > maxf
	}
}
