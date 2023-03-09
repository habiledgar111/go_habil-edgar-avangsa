package main

import (
	"fmt"
)

type pair struct {
	name  string
	count int
}

func mostAppearItem(items []string) []pair {
	if len(items) == 1 {
		var temp []pair
		temp = append(temp, pair{
			name:  items[0],
			count: 1,
		})
		return temp
	}

	result := mostAppearItem(items[:len(items)-1])
	for i := 0; i < len(result); i++ {
		if items[len(items)-1] == result[i].name {
			result[i].count += 1
			return result
		}
	}

	result = append(result, pair{
		name:  items[len(items)-1],
		count: 1,
	})

	for i := 0; i < len(result); i++ {
		for j := i + 1; j < len(result); j++ {
			if result[i].count > result[j].count {
				temp := result[i]
				result[i] = result[j]
				result[j] = temp
			}
		}
	}
	return result
}
func main() {
	fmt.Println(mostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}))
	fmt.Println(mostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}))
	fmt.Println(mostAppearItem([]string{"football", "basketball", "tenis"}))
}
