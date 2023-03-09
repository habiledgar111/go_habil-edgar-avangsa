package main

import (
	"fmt"
)

func main() {
	fmt.Println(pairSum2([]int{1, 2, 3, 4, 6}, 6))
	fmt.Println(pairSum2([]int{2, 5, 9, 11}, 11))
	fmt.Println(pairSum2([]int{1, 3, 5, 7}, 12))
	fmt.Println(pairSum2([]int{1, 4, 6, 8}, 10))
	fmt.Println(pairSum2([]int{1, 5, 6, 7}, 6))
}

func pairSum2(arr []int, target int) []int {
	if len(arr) == 0 || target == 0 {
		return []int{}
	}

	result := []int{}
	for i := range arr {
		for a := i + 1; a < len(arr); a++ {
			if arr[i]+arr[a] == target {
				result = append(result, i)
				result = append(result, a)
			}
		}
	}
	return result
}
