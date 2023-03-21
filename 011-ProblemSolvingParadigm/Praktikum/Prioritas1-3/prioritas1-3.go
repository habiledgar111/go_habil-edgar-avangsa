package main

import (
	"fmt"
)

func fibbo(n int) int {
	fibboarr := []int{}
	for i := 0; i <= n; i++ {
		if i == 0 || i == 1 {
			fibboarr = append(fibboarr, i)
		} else {
			fibboarr = append(fibboarr, (fibboarr[i-1] + fibboarr[i-2]))
		}
	}
	return fibboarr[n]
}

func main() {
	fmt.Println(fibbo(5))
	fmt.Println(fibbo(10))

}
