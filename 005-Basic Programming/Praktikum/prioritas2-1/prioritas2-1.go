package main

import "fmt"

func main() {
	fmt.Println("masukan bilangan")
	var bil int
	fmt.Scan(&bil)

	for i := 0; i < bil; i++ {
		//spacing
		for a := 1; a < (bil - i); a++ {
			fmt.Print(" ")
		}
		for j := 0; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println("")
	}
}
