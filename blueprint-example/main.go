package main //pkgm

import (
	"fmt"
	"geometry-lib/shape"
)

func main() {
	// mau menghitung luas gedung 2 lantai yang berbentuk persegi panjang
	firstFloor := shape.Rectangle{Width: 7.5, Length: 6.5}
	secondFloor := shape.Rectangle{Width: 4.5, Length: 5.5}

	totalArea := firstFloor.Area() + secondFloor.Area()

	fmt.Println("Total Area :", totalArea)
}
