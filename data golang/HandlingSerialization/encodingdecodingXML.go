// package main

// import (
// 	"encoding/xml"
// 	"fmt"
// )

// type Mahasiswa struct {
// 	Nama    string `xml:"nama"`
// 	Usia    int    `xml:"usia"`
// 	Jurusan string `xml:"jurusan"`
// }

// func main() {
// 	mhs := Mahasiswa{"Rara", 19, "Seni musik"}
// 	xmlData, err := xml.Marshal(mhs)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	fmt.Println("Data XML:", string(xmlData))
// 	var mhsDecode Mahasiswa
// 	err = xml.Unmarshal(xmlData, &mhsDecode)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	fmt.Println("Struct dari XML:", mhsDecode)
// }
