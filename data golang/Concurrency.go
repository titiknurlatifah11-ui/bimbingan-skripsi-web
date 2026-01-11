// package main

// import (
// 	"fmt"
// 	"time"
// )
// func sayHello() {
// 	fmt.Println("Hello from Goroutine!")
// }
// func main() {
//     go sayHello() // Menjalankan fungsi sebagai goroutine
//     fmt.Println("Hello from Main Function")
// 	time.Sleep(1 * time.Second) //menunggu goroutine selesai
// }

// package main

// import (
// 	"fmt"
// 	"time"
// )
// func printMessages() {
// 	for i := 1; i <= 5; i++ {
// 		fmt.Printf("Message %d from Goroutine\n", i)
// 		time.Sleep(500 * time.Millisecond)	
//     }
// }
// func main() {
//     go  printMessages() // Menjalankan fungsi sebagai goroutine
//     fmt.Println("Main function continues to run concurrently...")
// 	time.Sleep(1 * time.Second) //Memberikan waktu bagi goroutine untuk menyelesaikan
// 	//tugasnya
	
// }

// package main

// import (
// 	"fmt"
// 	"time"
// )
// func calculateSquare(num int) {
// 	square := num * num
//     fmt.Printf("Square of %d is %d\n", num, square)
// }
// func main() {
// 	for i := 1; i <= 5; i++ {
// 		go calculateSquare(i)
//     }
// 	time.Sleep(2 * time.Second) // Memberikan waktu bagi semua goroutine untuk selesai
// }

package main

import (
	"fmt"
	"sync"
)

func factorial(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
    }
	fmt.Printf("Factorial of %d is %d\n", n, result)
}
func main() {
	var wg sync.WaitGroup
    numbers := []int{5, 7, 10, 12}

    for _, num := range numbers {
		wg.Add(1)
		go factorial(num, &wg)
    }
    wg.Wait() // Menunggu semua goroutine selesai
}