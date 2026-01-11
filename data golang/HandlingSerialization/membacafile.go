package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("sample.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Isi file:", string(data))
}
