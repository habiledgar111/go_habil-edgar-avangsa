package main

import (
	"fmt"
)

type mobil struct {
	tipe   string
	fuelIn int
}

func (m mobil) perkiraanJauh() float64 {
	return float64(m.fuelIn) / 1.5
}
func main() {
	car := mobil{
		tipe:   "sedan",
		fuelIn: 50,
	}

	car2 := mobil{
		tipe:   "truck",
		fuelIn: 90,
	}

	fmt.Printf("perkiraan jauh mobil 1 : (mil) %v\n", car.perkiraanJauh())
	fmt.Printf("perkiraan jauh mobil 2 : (mil) %v\n", car2.perkiraanJauh())

}
