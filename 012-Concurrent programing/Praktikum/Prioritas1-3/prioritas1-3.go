package main

import "fmt"

func main() {
	arr := []int{3, 6, 9}
	ch := make(chan int, len(arr))
	go func(arr []int, x chan int) {
		for _, item := range arr {
			ch <- item
		}
	}(arr, ch)

	for i := 0; i < len(arr); i++ {
		fmt.Println("bilangan ", <-ch)
	}
}
