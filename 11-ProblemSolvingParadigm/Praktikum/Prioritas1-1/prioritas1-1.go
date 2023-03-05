package main

import (
	"fmt"
)

func binerarr(n int) []string {
	result := []string{}
	if n == 0 {
		return []string{"0"}
	}
	// for i := 0; i <= n; i++ {
	// 	// result = append(result, strconv.FormatInt(int64(i), 2))
	// }
	result = append(result, "0")
	for i := 0; i <= n; i++ {
		biner := ""
		temp := ""
		index := i
		for index > 0 {
			temp = ""
			if (index & 1) == 0 {
				temp = "0"
			} else {
				temp = "1"
			}
			biner = temp + biner
			index >>= 1
		}
		result = append(result, biner)
	}
	return result
}

func main() {
	fmt.Println(binerarr(5))
	fmt.Println(binerarr(38))
}
