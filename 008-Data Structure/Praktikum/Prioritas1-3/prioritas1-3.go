package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(munculSekali("1234321"))
	fmt.Println(munculSekali("76523752"))
	fmt.Println(munculSekali("1122334455"))
	fmt.Println(munculSekali("0872504"))
}

func munculSekali(angka string) []int {

	if angka == "" {
		return []int{}
	}

	temp := make(map[int]int)
	for i := range angka {
		str := string(angka[i])
		t, _ := strconv.Atoi(str)
		if _, value := temp[t]; !value {
			temp[t] = 1
		} else {
			temp[t]++
		}
	}

	result := []int{}

	for key, _ := range temp {
		if temp[key] == 1 {
			result = append(result, key)
		}
	}

	return result
}
