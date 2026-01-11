package main

import "fmt"

func main() {
    for i := 1; i <= 10; i++ {
        if i == 5 {
            continue // Lewati angka 5
        }
        fmt.Println(i)
    }
}
