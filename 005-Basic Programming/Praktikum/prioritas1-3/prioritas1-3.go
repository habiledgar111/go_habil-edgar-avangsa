package main

import "fmt"

func main() {
	fmt.Println("masukan nilai")
	var nilai int
	fmt.Scan(&nilai)

	if nilai < 0 || nilai > 100 {
		fmt.Println("nilai tidak valid")
	}

	if nilai >= 0 && nilai <= 34 {
		fmt.Println("E")
	}

	if nilai >= 35 && nilai <= 49 {
		fmt.Println("D")
	}

	if nilai >= 50 && nilai <= 64 {
		fmt.Println("C")
	}

	if nilai >= 65 && nilai <= 79 {
		fmt.Println("B")
	}

	if nilai >= 80 && nilai <= 100 {
		fmt.Println("A")
	}
}
