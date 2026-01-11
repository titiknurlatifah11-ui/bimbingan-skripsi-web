package main

import "fmt"

func main() {
    for i := 1; i <= 10; i++ {
        fmt.Printf("5 x %d = %d\n", i, 5*i)
    }
}
// package main

// import "fmt"

// func main() {
//     fmt.Println("Bilangan prima dari 1 hingga 50:")
//     for num := 2; num <= 50; num++ {
//         isPrime := true
//         for i := 2; i*i <= num; i++ {
//             if num % i == 0 {
//                 isPrime = false
//                 break
//             }
//         }
//         if isPrime {
//             fmt.Println(num)
//         }
//     }
// }

