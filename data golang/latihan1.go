package main

import (
	"fmt"
)

func main() {
	var feet float64
	fmt.Print("Masukkan panjang dalam kaki: ")
	fmt.Scan(&feet)

	meters := feet * 0.3048
	fmt.Printf("%.2f kaki adalah %.2f meter\n", feet, meters)
}

