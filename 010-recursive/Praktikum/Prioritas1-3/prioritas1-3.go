package main

import (
	"fmt"
)

func primeX(number int) int {
	prime := []int{2}
	var temp = 0
	var primeNumber = 3
	for len(prime) <= number {
		for i := 2; i < primeNumber; i++ {
			if primeNumber%i == 0 {
				temp++
				break
			}
		}

		if temp == 0 {
			prime = append(prime, primeNumber)
			primeNumber++
		} else {
			temp = 0
			primeNumber++
		}
	}

	return prime[number-1]
}

func main() {
	fmt.Println(primeX(1))
	fmt.Println(primeX(5))
	fmt.Println(primeX(8))
	fmt.Println(primeX(9))
	fmt.Println(primeX(10))

}
