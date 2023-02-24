package main

import (
	"fmt"
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
	fmt.Println("angka : ", angka[0])
	return []int{}
}
