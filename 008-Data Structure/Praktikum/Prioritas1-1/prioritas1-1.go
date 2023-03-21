package main

import (
	"fmt"
)

func main() {
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))
	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))
	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
	fmt.Println(ArrayMerge([]string{}, []string{}))

}

func ArrayMerge(arrayA, arrayB []string) []string {
	if len(arrayA) == 0 && len(arrayB) == 0 {
		return []string{}
	}

	if len(arrayA) == 0 {
		return arrayB
	}

	if len(arrayB) == 0 {
		return arrayA
	}

	sliceReturn := []string{}
	tempslice := append(arrayA, arrayB...)
	temp := make(map[string]bool)

	//pengecekan apakah ada yang sama atau tidak
	for _, item := range tempslice {
		//mengecek apakah map dengan key = item sudah ada atau tidak jika belum ada maka masukan ke dalam map kalau sudah ada
		//maka lewati
		_, value := temp[item]
		if !value {
			temp[item] = true
			sliceReturn = append(sliceReturn, item)
		}
	}
	return sliceReturn
}
