package main

import "fmt"

func main() {
	var user, pass string

	var succes bool

	for succes = false; !succes; {
		fmt.Print("user : ")
		fmt.Scanln(&user)
		fmt.Print("pass : ")
		fmt.Scanln(&pass)

		succes = user == "123A" && pass == "123A"
	}
	fmt.Println("welcome")
}
