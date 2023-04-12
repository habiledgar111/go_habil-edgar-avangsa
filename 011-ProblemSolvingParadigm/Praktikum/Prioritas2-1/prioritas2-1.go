package main

import (
	"fmt"
	"math"
)

//greddy
// func Frog(jumps []int) int {
// 	pointer := 0
// 	result := 0
// 	for pointer < len(jumps)-2 {
// 		value1 := math.Abs(float64(jumps[pointer]) - float64(jumps[pointer+1]))
// 		value2 := math.Abs(float64(jumps[pointer]) - float64(jumps[pointer+2]))
// 		if value1 < value2 {
// 			result += int(value1)
// 			pointer++
// 			fmt.Println("result ", result, "pointer ", pointer)
// 		} else {
// 			result += int(value2)
// 			pointer += 2
// 			fmt.Println("result ", result, "pointer ", pointer)
// 		}
// 	}
// 	return result
// }

// dynamic
func Frog(jumps []int) int {
	//deklarasi 2 variable yang akan menyimpan setiap langkah yang akan dijalankan
	langkah1Sebelumnya := 0
	langkah2Sebelumnya := math.MaxInt
	//lakukan looping sepanjang array
	for i := 1; i < len(jumps); i++ {
		//loncat 1 akan digunakan untuk katak loncat 1 batu sedangan loncat2 akan digunakan katak untuk loncat 2 batu
		//lankah 1 akan selalu tertingal 1 index di belakang i jadi langkah satu akan selalu loncat 1 batu
		loncat1 := langkah1Sebelumnya + int(math.Abs(float64(jumps[i])-float64(jumps[i-1])))
		loncat2 := langkah2Sebelumnya
		//karena katak akan loncat 2 batu maka batu yang dituju tidak boleh dengan index == 1 atau harus lebih dari 1
		if i > 1 {
			//langkah2sebelumnya akan selalu tertingal 2 index di belakang i jadi langkah2 akan selalu loncat 2 batu
			loncat2 = langkah2Sebelumnya + int(math.Abs(float64(jumps[i])-float64(jumps[i-2])))
		}
		temp := 0
		//di lakukan pengecekan apakah loncat 1 dan loncat lebih kecil
		if loncat1 < loncat2 {
			temp = loncat1
		} else {
			temp = loncat2
		}
		//yang paling kecil akan dimasukan ke langkah 1 sebelumnya dan yang lebih besar akan diabaikan
		//setelah katak loncat nilai yang paling kecil akan disimpan pada langkah1sebelumya
		//dan nilai paling kecil dari langkah sebelumnya akan disimpan pada langkah2sebelumnya
		langkah2Sebelumnya = langkah1Sebelumnya
		langkah1Sebelumnya = temp
	}
	result := langkah1Sebelumnya
	return result
}

func main() {
	fmt.Println(Frog([]int{10, 30, 40, 20}))
	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50}))
}
