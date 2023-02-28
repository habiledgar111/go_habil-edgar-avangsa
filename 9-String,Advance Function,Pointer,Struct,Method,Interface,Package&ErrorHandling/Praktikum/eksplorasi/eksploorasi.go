package main

import (
	"fmt"
)

type student struct {
	name       string
	nameEncode string
	score      int
}

type chiper interface {
	encode() string
	decode() string
}

func (s student) encode() string {
	//dari stirng ke chiper code
	var nameEncode = s.name
	var key = 15
	result := ""
	for i := 0; i < len(nameEncode); i++ {
		temp := nameEncode[i] + byte(key)
		if temp > 122 {
			temp -= 122
			temp += 96
		}
		result += string(temp)
	}
	return result
}

func (s student) decode() string {
	//dari chiper code ke string
	var nameDecode = s.nameEncode
	var key = 15
	result := ""
	for i := 0; i < len(nameDecode); i++ {
		temp := nameDecode[i] - byte(key)
		if temp < 97 {
			temp = 97 - temp
			temp = 123 - temp
		}
		result += string(temp)
	}
	return result
}

func main() {
	var menu int
	var a student = student{}
	var c chiper = &a

	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu? ")
	fmt.Scan(&menu)

	if menu == 1 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&a.name)
		fmt.Print("\nEncode of student’s name " + a.name + "is : " + c.encode())
	} else if menu == 2 {
		fmt.Print("\nInput Student Name Code: ")
		fmt.Scan(&a.nameEncode)
		fmt.Print("\nDecode of student’s name " + a.nameEncode + "is : " + c.decode())
	}
}
