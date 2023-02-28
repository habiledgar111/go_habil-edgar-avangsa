package main

import (
	"fmt"
)

func caesar(offset int, input string) string {
	geser := offset % 26
	result := ""
	for i := 0; i < len(input); i++ {
		temp := input[i] + byte(geser)
		if temp > 122 {
			temp -= 122
			temp += 97
		}
		result += string(temp)
	}
	return result
}

func main() {
	fmt.Println(caesar(3, "abc"))
	fmt.Println(caesar(2, "alta"))
	fmt.Println(caesar(10, "alterraacademy"))
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz"))
	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz"))
}
