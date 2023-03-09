package main

import "fmt"

func decToRoman(n int) string {
	// ribuan := []string{"", "M", "MM", "MMM", "MMMM", "MMMMM", "MMMMMM", "MMMMMMM", "MMMMMMMM", "MMMMMMMMM", "MMMMMMMMMM"}
	ratusan := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	puluhan := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	satuan := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	if n < 0 {
		return "tidak boleh minus"
	}
	result := ""
	if n > 999 {
		temp := n / 1000
		for i := 1; i <= temp; i++ {
			result += "M"
		}
	}
	if n > 99 {
		result += ratusan[(n%1000)/100]
	}
	if n > 9 {
		result += puluhan[(n%100)/10]
	}
	result += satuan[n%10]
	return result
}
func main() {
	//batasan hanya angka sampai 10000
	fmt.Println(decToRoman(55))
	fmt.Println(decToRoman(999))
	fmt.Println(decToRoman(10000))
	fmt.Println(decToRoman(15431))
}
