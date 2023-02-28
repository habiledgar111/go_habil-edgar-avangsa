package main

import (
	"fmt"
)

type student struct {
	name  string
	score int
}

func max(arr []student) student {
	max := 0
	for i := 0; i < len(arr); i++ {
		if max < arr[i].score {
			max = arr[i].score
		}
	}
	for j := 0; j < len(arr); j++ {
		if max == arr[j].score {
			return arr[j]
		}
	}
	return student{}
}

func min(arr []student) student {
	min := arr[0].score
	for i := 1; i < len(arr); i++ {
		if min > arr[i].score {
			min = arr[i].score
		}
	}
	for j := 0; j < len(arr); j++ {
		if min == arr[j].score {
			return arr[j]
		}
	}
	return student{}
}

func main() {
	var nama string
	var score int
	var students []student
	for i := 0; i < 5; i++ {
		fmt.Println("masukan nama siswa ke-", i)
		fmt.Scanln(&nama)
		fmt.Println("masukan score siswa ke-", i)
		fmt.Scanln(&score)
		students = append(students, student{
			name:  nama,
			score: score,
		})
	}
	sum := 0
	for i := 0; i < len(students); i++ {
		sum += students[i].score
	}
	sum = sum / 5
	fmt.Println("average : ", sum)
	fmt.Printf("Min score of student : %v\n", min(students))
	fmt.Printf("Max score of student : %v\n", max(students))
}
