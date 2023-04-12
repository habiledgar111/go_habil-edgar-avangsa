package main

import (
	"fmt"
)

func main() {
	fmt.Println("masukan bilangan")
	var bilangan int
	fmt.Scanln(&bilangan)
	fmt.Println("masukan pangkat")
	var pangkat int
	fmt.Scanln(&pangkat)
	fmt.Println(pow(bilangan, pangkat))
}

// func pow(x int, n int) int {
// 	//big O (6+n/2) atau big O (5+n/2)????
// 	if n == 0 {
// 		return 1
// 	}
// 	if n == 1 {
// 		return x
// 	}
// 	if n == 2 {
// 		return x * x
// 	}
// 	result := 1
// 	for i := 1; i <= n/2; i++ {
// 		result *= x
// 	}

// 	if n%2 == 0 {
// 		return result * result
// 	}

// 	return result * result * x
// }

func pow(x int, n int) int {
	//big O (log n)
	if n == 0 {
		return 1
	}

	result := pow(x, n/2)

	if n%2 == 0 {
		return result * result
	}
	return result * result * x
}
