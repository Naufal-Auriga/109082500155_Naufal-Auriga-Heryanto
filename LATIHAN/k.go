package main
import "fmt"


func main() {
	var a, b ont

	fmt.Scan(&a, &b)

	switch (a+b)%4 {
	case 0:
		fmt.Println("hasilnya", a, "+", b, "%4 adalah", (a+b)%4
	)
	}
}