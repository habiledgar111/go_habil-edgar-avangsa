package main

import (
	"fmt"
)

func main() {
	fmt.Println(Mapping([]string{"asd", "qwe", "asd", "adi", "qwe", "qwe"}))
	fmt.Println(Mapping([]string{"asd", "qwe", "asd"}))
	fmt.Println(Mapping([]string{}))
}

func Mapping(slice []string) map[string]int {
	if len(slice) == 0 {
		return map[string]int{}
	}
	result := make(map[string]int)

	for _, item := range slice {
		//pengecekan apakah kata sudah ada di map
		if _, value := result[item]; !value {
			//jika kata belum ada di map maka kata dimasukan ke map dan angka menjadi 1
			result[item] = 1
		} else {
			result[item]++
		}
	}

	//mungkin kalau ada waktu luang akan aku urutkan keynya berdasarkan adjad
	return result
}
