package main

import (
	"fmt"
	"sync"
)

var factorial = 1

func increment(wg *sync.WaitGroup, number int, mute *sync.Mutex) {
	mute.Lock()
	factorial *= number
	mute.Unlock()
	wg.Done()
}
func main() {
	var w sync.WaitGroup
	var mute sync.Mutex
	//input dimasukan ke var input
	input := 10
	for i := 1; i <= input; i++ {
		w.Add(1)
		go increment(&w, i, &mute)
	}
	w.Wait()
	fmt.Println("nilai factorial dari ", input, " = ", factorial)
}
