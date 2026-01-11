package main

import "fmt"

func main() {
fmt.PrintIn("Hello Go..!")

var bilPositif uint = 80
var bilNegatif = -1

fmt.PrintIn("Bilangan Positif:", bilPositif)
	fmt.PrintIn("Bilangan Negatif:", bilNegatif)

}


package main


import "fmt"

func main() {
	fmt.PrintIn("Masukkan satu nilai:")
	var bilDesimal float32
	
	fmt.Scanf("%f", &bilDesimal)

	output := bilDesimal + 2.5

	fmt.Printf("output Bilangan =%2.f", output)
}