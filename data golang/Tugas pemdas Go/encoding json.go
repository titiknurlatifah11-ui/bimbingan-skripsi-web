// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// type Mahasiswa struct {
// 	Nama    string `json:"nama"`
// 	Usia    int    `json:"usia"`
// 	Jurusan string `json:"jurusan"`
// }

// func main() {
// 	mhs := Mahasiswa{"Himma", 19, "Informatika"}
// 	jsonData, err := json.Marshal(mhs)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	fmt.Println("Data JSON:", string(jsonData))
// }
