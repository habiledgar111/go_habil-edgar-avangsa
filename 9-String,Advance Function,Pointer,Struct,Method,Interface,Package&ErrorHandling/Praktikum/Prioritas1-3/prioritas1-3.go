package main

import (
	"fmt"
	"strings"
)

func compare(a, b string) string {
	status := false
	if len(a) > len(b) {
		status = strings.Contains(a, b)
		if status {
			return b
		}
	}

	status = strings.Contains(b, a)
	if status {
		return a
	}
	return ""
}

func main() {
	fmt.Println(compare("AKA", "AKASHI"))
	fmt.Println(compare("KANGOORO", "KANG"))
	fmt.Println(compare("KI", "KIJANG"))
	fmt.Println(compare("KUPU-KUPU", "KUPU"))
	fmt.Println(compare("ILALANG", "ILA"))
}
