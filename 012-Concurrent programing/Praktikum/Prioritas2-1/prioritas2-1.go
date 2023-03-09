package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

var mapp = make(map[string]int)

func main() {
	fmt.Println("masukan input")
	buff := bufio.NewReader(os.Stdin)
	input, _ := buff.ReadString('\n')
	var w sync.WaitGroup
	var mute sync.Mutex

	for i := 0; i < len(input); i++ {
		w.Add(1)
		go countChar(input[i:i+1], &w, &mute)
	}

	w.Wait()
	fmt.Println(mapp)

}
func countChar(input string, w *sync.WaitGroup, mute *sync.Mutex) {
	mute.Lock()
	if _, value := mapp[input]; !value {
		mapp[input] = 1
	} else {
		mapp[input]++
	}
	mute.Unlock()
	w.Done()
}
