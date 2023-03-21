package main

import (
	"fmt"
)

func simpleEquations(a, b, c int) interface{} {
	//formula 1 x+y+z
	//formula 2 xyz
	//formula 3 x*x + y*y + z*z
	//y/a+y/b+y/c = x^2-z/2
	for x := 1; x <= b; x++ {
		for y := 1; y <= b; y++ {
			for z := 1; z <= b; z++ {
				if x+y+z == a && x*y*z == b && (x*x)+(y*y)+(z*z) == c {
					return []int{x, y, z}
				}
			}
		}
	}
	return "no solution"
}
func main() {
	fmt.Println(simpleEquations(1, 2, 3))
	fmt.Println(simpleEquations(6, 6, 14))
	fmt.Println(simpleEquations(12, 12, 28))
	fmt.Println(simpleEquations(7, 12, 17))
}
