package main

import "fmt"

func main() {
	fmt.Println("masukan bialngan")
	var bil int
	fmt.Scan(&bil)

	for i := 1; i <= bil; i++ {
		if bil%i == 0 {
			fmt.Println(i)
		}
	}
}
