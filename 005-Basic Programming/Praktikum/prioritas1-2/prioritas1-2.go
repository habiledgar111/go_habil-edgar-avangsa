package main

import "fmt"

func main() {
	fmt.Println("masukan bilangan")
	var bil int
	fmt.Scan(&bil)

	if bil%2 == 0 {
		fmt.Println("bilangan genap")
	} else {
		fmt.Println("bilangan ganjil")
	}
}
