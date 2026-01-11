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
// 	jsonString := `{"nama": "Nazwa", "usia": 19, "jurusan": "kebidanan"}`
// 	var mhs Mahasiswa
// 	err := json.Unmarshal([]byte(jsonString), &mhs)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	fmt.Println("Struct dari JSON:", mhs)
// }
