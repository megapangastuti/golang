package main

import (
	"fmt"
	"geometry-lib/shape"
)

func main() {
	// menghitung hasil panen dari ladang yang berbentuk persegi panjang

	field1 := shape.Rectangle{Width: 15.0, Length: 5.0}
	field2 := shape.Rectangle{Width: 17.5, Length: 15.5}

	harvestField1 := field1.Area() / 100
	harvestField2 := field2.Area() / 100

	fmt.Println("Harvest Field 1 :", harvestField1)
	fmt.Println("Harvest Field 2 :", harvestField2)
}
