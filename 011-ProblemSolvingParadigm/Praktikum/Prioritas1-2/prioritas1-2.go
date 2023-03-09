package main

import (
	"fmt"
)

func pascalTriangle(n int) [][]int {
	result := [][]int{}

	for i := 0; i < n; i++ {
		value := []int{}
		for j := 0; j <= i; j++ {
			//formula nCr = n!/(n-r)!*r!
			temp := findFactorial(i) / (findFactorial(i-j) * findFactorial(j))
			// value += strconv.Itoa(temp)
			value = append(value, temp)

		}
		result = append(result, value)
	}
	return result
}

func findFactorial(n int) int {
	if n == 1 || n == 0 {
		return 1
	}
	return n * findFactorial(n-1)
}
func main() {
	fmt.Println(pascalTriangle(5))
	fmt.Println(pascalTriangle(10))
}
