package main

import "fmt"

func playDomino(cards [][]int, deck []int) interface{} {
	//[n][2]
	for i := 0; i < len(cards); i++ {
		for j := 0; j < 2; j++ {
			if cards[i][j] == deck[0] || cards[i][j] == deck[1] {
				return cards[i]
			}
		}
	}
	return "tutup kartu"
}
func main() {
	fmt.Println(playDomino([][]int{[]int{6, 5}, []int{3, 4}, []int{2, 1}, []int{3, 3}}, []int{4, 3}))
	fmt.Println(playDomino([][]int{[]int{6, 5}, []int{3, 3}, []int{3, 4}, []int{2, 1}}, []int{3, 6}))
	fmt.Println(playDomino([][]int{[]int{6, 6}, []int{2, 4}, []int{3, 6}}, []int{5, 1}))
}
