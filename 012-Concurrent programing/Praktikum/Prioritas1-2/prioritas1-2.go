package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		ch <- 3
	}()
	for i := <-ch; i <= 12; i += 3 {
		fmt.Println("bilangan : ", i)
	}
}
