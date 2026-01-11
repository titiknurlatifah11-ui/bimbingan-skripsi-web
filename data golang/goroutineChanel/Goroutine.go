//Mencetak Pesan Secara Concurrent
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

//Menghitung Kuadrat Secara Concurrent
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

//Menjalankan Fungsi Anonymous Concurrently
package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 1; i <= 3; i++ {
		go func(n int) {
			fmt.Printf("Anonymous Goroutine with value %d\n", n)
		}(i)
	}
	
	time.Sleep(1 * time.Second)
}

