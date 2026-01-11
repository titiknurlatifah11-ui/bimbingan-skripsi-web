package main

import (
    "fmt"
    "os"
)

func main() {
	file, err := os.Open("C:/path/to/sample.bin")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer file.Close()

    buffer := make([]byte, 64)
    _, err = file.Read(buffer)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Isi file biner:", buffer)
}
