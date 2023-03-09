package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("masukan kata")
	var kata string

	//cara untuk scan line
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	kata = scanner.Text()

	palindrome := true

	for i := 0; i < (len(kata)-1)/2; i++ {
		if kata[i] != kata[(len(kata)-i)-1] {
			palindrome = false
		}
	}

	if palindrome {
		fmt.Println("kata palindorme")
	} else {
		fmt.Println("bukan palindrome")
	}
}
