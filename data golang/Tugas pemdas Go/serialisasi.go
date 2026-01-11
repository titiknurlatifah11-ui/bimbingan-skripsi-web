// package main

// import (
// 	"encoding/gob"
// 	"fmt"
// 	"os"
// )

// type Mahasiswa struct {
// 	Nama    string
// 	Usia    int
// 	Jurusan string
// }

// func main() {
// 	mhs := Mahasiswa{"Faisal", 24, "Matematika"}

// 	// Membuat file untuk menyimpan data dalam format gob
// 	file, err := os.Create("mahasiswa.gob")
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	defer file.Close()

// 	// Membuat encoder untuk menulis ke file dalam format gob
// 	encoder := gob.NewEncoder(file)
// 	err = encoder.Encode(mhs)
// 	if err != nil {
// 		fmt.Println("Error encoding:", err)
// 		return
// 	}

// 	fmt.Println("Data berhasil disimpan dalam format gob")
// }
