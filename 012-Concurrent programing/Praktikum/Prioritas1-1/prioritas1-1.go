package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	number := 10
	go func(number int) {
		ticker := time.NewTicker(3 * time.Second)
		// number := <-ch
		for {
			select {
			case t := <-ticker.C:
				number *= number
				fmt.Println("Time : ", t, "number : ", number)
			case <-ch:
				ticker.Stop()
				return
			}
		}
	}(number)
	time.Sleep(10 * time.Second)
}
