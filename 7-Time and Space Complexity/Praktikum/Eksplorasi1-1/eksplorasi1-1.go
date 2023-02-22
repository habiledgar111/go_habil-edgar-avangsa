package main

import (
	"fmt"
	"math"
)

func main() {
	var bilangan int
	fmt.Println("masukan bilangan")
	fmt.Scanln(&bilangan)
	fmt.Println(primeNumber(bilangan))
}

func primeNumber(bil int) string {
	if bil <= 1 {
		return "Bukan Bilangan Prima"
	}
	for i := 2; i < int(math.Sqrt(float64(bil))); i++ {
		if bil%i == 0 {
			return "Bukan Bilangan Prima"
		}
	}

	return "Bilangan Prima"
}
