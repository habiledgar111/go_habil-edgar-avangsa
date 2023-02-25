package main

import (
	"fmt"
)

func main() {
	fmt.Println(countMatrix([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{9, 8, 9},
	}))

	//intputan tambahan
	fmt.Println(countMatrix([][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
		{21, 22, 23, 24, 25},
	}))
}

func countMatrix(matrix [][]int) int {
	sum1 := 0
	sum2 := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {
			if i == j {
				sum1 += matrix[i][j]
			}
			if i+j == len(matrix)-1 {
				sum2 += matrix[i][j]
			}
		}
	}

	if sum1 > sum2 {
		return sum1 - sum2
	}

	return sum2 - sum1
}
