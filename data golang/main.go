package main

import "fmt"

func main() {

	for i := 0; i < 5; i++{
		fmt.Println(i)
	}


// //infinite loop
i := 0
for {
	fmt.Println(i)
	i++
	if i == 5 {
		break
	}
}

// //while loop
y := 0
for y < 5 {
	fmt.Println(y)
	y++ 
}

// //nested loops
for i := 1; i <= 3; i++ {
	for j := 1; j <= 3; j++ {
		fmt.Printf("i=%d, j=%d\n", i, j)
	}
}
}