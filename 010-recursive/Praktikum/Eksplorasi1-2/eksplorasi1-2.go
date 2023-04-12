package main

import (
	"fmt"
)

func maximumBuyProduct(money int, productPrices []int) int {
	sortedArr := sortProduct(productPrices)
	result := 0
	for i := 0; i < len(sortedArr); i++ {
		if money >= sortedArr[i] && money > 0 {
			money -= sortedArr[i]
			result++
		}
	}
	return result
}
func sortProduct(product []int) []int {
	for i := 0; i < len(product); i++ {
		for j := i + 1; j < len(product); j++ {
			if product[j] < product[i] {
				temp := product[j]
				product[j] = product[i]
				product[i] = temp
			}
		}
	}
	return product
}
func main() {
	fmt.Println(maximumBuyProduct(50000, []int{25000, 25000, 10000, 14000}))
	fmt.Println(maximumBuyProduct(30000, []int{15000, 10000, 12000, 5000, 3000}))
	fmt.Println(maximumBuyProduct(10000, []int{2000, 3000, 1000, 2000, 10000}))
	fmt.Println(maximumBuyProduct(4000, []int{7500, 3000, 2500, 2000}))
	fmt.Println(maximumBuyProduct(0, []int{10000, 30000}))
}
