package main

import "fmt"

func main() {
	fmt.Println("masukan tinggi")
	var tinggi int
	fmt.Scan(&tinggi)

	fmt.Println("masukan alas 1")
	var alas1 int
	fmt.Scan(&alas1)

	fmt.Println("masukan alas 2")
	var alas2 int
	fmt.Scan(&alas2)

	var alas = (alas1 + alas2) / 2
	var luas = alas * tinggi

	fmt.Println("luas = ", luas)
}
