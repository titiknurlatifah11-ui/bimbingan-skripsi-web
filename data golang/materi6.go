//1. Concurrency
//Contoh Sederhana Concurrency di Go
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


//2. Goroutine
//2.2  Mencetak Pesan Secara Concurrent
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

//2.3 Menghitung Kuadrat Secara Concurrent
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


//2.4 Menjalankan Fungsi Anonymous Concurrently
// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	for i := 1; i <= 3; i++ {
// 		go func(n int) {
// 			fmt.Printf("Anonymous Goroutine with value %d\n", n)
// 		}(i)
// 	}
	
// 	time.Sleep(1 * time.Second)
// }


//2.5 Contoh Tingkat Menengah: Menghitung Faktorial secara Concurrent
// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func factorial(n int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	result := 1
// 	for i := 1; i <= n; i++ {
// 		result *= i
//     }
// 	fmt.Printf("Factorial of %d is %d\n", n, result)
// }
// func main() {
// 	var wg sync.WaitGroup
//     numbers := []int{5, 7, 10, 12}

//     for _, num := range numbers {
// 		wg.Add(1)
// 		go factorial(num, &wg)
//     }
//     wg.Wait() // Menunggu semua goroutine selesai
// }


//2.6 Mengambil Data dari Beberapa URL Secara Bersamaan
// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"sync"
// )

// func fetchURL(url string, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		fmt.Printf("Failed to fetch %s: %v\n", url, err)
// 		return
// 	}
// 	fmt.Printf("Fetched %s: %d\n", url, resp.StatusCode)
// }

// func main() {
// 	var wg sync.WaitGroup
// 	urls := []string{
// 		"https://www.google.com",
// 		"https://www.facebook.com",
// 		"https://www.twitter.com",
// 		"https://www.github.com",
// 	}

// 	for _, url := range urls {
// 		wg.Add(1)
// 		go fetchURL(url, &wg)
// 	}

// 	wg.Wait()
// }

//2.7 Studi Kasus: Pengelolaan Transaksi Bank Secara Concurrent
// package main

// import (
// 	"fmt"
// 	"sync"
// )

// var balance = 1000
// var mu sync.Mutex

// func withdraw(amount int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	mu.Lock()
// 	defer mu.Unlock()

// 	if balance >= amount {
// 		balance -= amount
// 		fmt.Printf("Withdrawal of %d successful. Remaining balance: %d\n", amount, balance)
// 	} else {
// 		fmt.Printf("Withdrawal of %d failed. Insufficient balance.\n", amount)
// 	}
// }

// func main() {
// 	var wg sync.WaitGroup
// 	transactions := []int{200, 300, 150, 600}

// 	for _, amount := range transactions {
// 		wg.Add(1)
// 		go withdraw(amount, &wg)
// 	}

// 	wg.Wait()
// 	fmt.Printf("Final Balance: %d\n", balance)
// }


//Channel
//3.2 Membuat dan Menggunakan Channel
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	message := make(chan string)

// 	go func() {
// 		message <- "Hello, Channel!"
// 	}()

// 	msg := <-message
// 	fmt.Println(msg)
// }

//3.3  Buffered Channel
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	message := make(chan string, 3)
// 	message <- "Buffer 1"
// 	message <- "Buffer 2"
// 	message <- "Buffer 3"

// 	fmt.Println(<-message)
// 	fmt.Println(<-message)
// 	fmt.Println(<-message)
// }


//3.4 . Sinkronisasi Goroutine dengan Channel
// package main

// import (
// 	"fmt"
// 	"time"
// )

// func worker(done chan bool) {
// 	fmt.Println("Starting work...")
// 	time.Sleep(2 * time.Second)
// 	fmt.Println("Work done!")
// 	done <- true
// }

// func main() {
// 	done := make(chan bool)
// 	go worker(done)
// 	<-done
// 	fmt.Println("Main program exits")
// }


//3.5 . Select Statement untuk Mengelola Multiple Channel
// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	ch1 := make(chan string)
// 	ch2 := make(chan string)

// 	go func() {
// 		time.Sleep(1 * time.Second)
// 		ch1 <- "Message from Channel 1"
// 	}()

// 	go func() {
// 		time.Sleep(2 * time.Second)
// 		ch2 <- "Message from Channel 2"
// 	}()

// 	select {
// 	case msg1 := <-ch1:
// 		fmt.Println(msg1)
// 	case msg2 := <-ch2:
// 		fmt.Println(msg2)
// 	}
// }


//3.6 Buffered vs Unbuffered Channel
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	message := make(chan string, 2)

// 	message <- "Message 1"
// 	fmt.Println("Sent Message 1")

// 	message <- "Message 2"
// 	fmt.Println("Sent Message 2")

// 	fmt.Println(<-message)
// 	fmt.Println(<-message)
// }


//3.7 Studi Kasus: Sistem Antrian Pemrosesan Pesanan
package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job)
		time.Sleep(1 * time.Second)
		fmt.Printf("Worker %d finished job %d\n", id, job)
		results <- job * 2
	}
}

func main() {
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= 5; j++ {
		jobs <- j
	}

	close(jobs)

	for a := 1; a <= 5; a++ {
		fmt.Printf("Result: %d\n", <-results)
	}
}
