// package main

// import (
// 	"encoding/json"
// 	"encoding/xml"
// 	"fmt"
// 	"io/ioutil"
// 	"os"
// 	"strings"
// )

// // Struct untuk menyimpan data mahasiswa
// type Mahasiswa struct {
// 	Nama    string `json:"nama" xml:"nama"`
// 	Usia    int    `json:"usia" xml:"usia"`
// 	Jurusan string `json:"jurusan" xml:"jurusan"`
// }

// // Fungsi untuk membaca data dari file teks
// func bacaData(filePath string) (Mahasiswa, error) {
// 	var mahasiswa Mahasiswa

// 	data, err := ioutil.ReadFile(filePath)
// 	if err != nil {
// 		return mahasiswa, err
// 	}

// 	lines := strings.Split(string(data), "\n")
// 	for _, line := range lines {
// 		if strings.TrimSpace(line) == "" {
// 			continue
// 		}
// 		parts := strings.Split(line, ": ")
// 		if len(parts) != 2 {
// 			continue
// 		}
// 		key := strings.TrimSpace(parts[0])
// 		value := strings.TrimSpace(parts[1])

// 		switch key {
// 		case "Nama":
// 			mahasiswa.Nama = value
// 		case "Usia":
// 			fmt.Sscanf(value, "%d", &mahasiswa.Usia)
// 		case "Jurusan":
// 			mahasiswa.Jurusan = value
// 		}
// 	}

// 	return mahasiswa, nil
// }

// // Fungsi untuk menyimpan data dalam format JSON
// func simpanJSON(mahasiswa Mahasiswa, filePath string) error {
// 	file, err := os.Create(filePath)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	encoder := json.NewEncoder(file)
// 	encoder.SetIndent("", "  ")
// 	return encoder.Encode(mahasiswa)
// }

// // Fungsi untuk menyimpan data dalam format XML
// func simpanXML(mahasiswa Mahasiswa, filePath string) error {
// 	file, err := os.Create(filePath)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	encoder := xml.NewEncoder(file)
// 	encoder.Indent("", "  ")
// 	return encoder.Encode(mahasiswa)
// }

// // Main program
// func main() {
// 	// Baca data dari file teks
// 	mahasiswa, err := bacaData("mahasiswa.txt")
// 	if err != nil {
// 		fmt.Println("Error membaca data:", err)
// 		return
// 	}

// 	// Simpan dalam format JSON
// 	err = simpanJSON(mahasiswa, "mahasiswa.json")
// 	if err != nil {
// 		fmt.Println("Error menyimpan JSON:", err)
// 		return
// 	}

// 	// Simpan dalam format XML
// 	err = simpanXML(mahasiswa, "mahasiswa.xml")
// 	if err != nil {
// 		fmt.Println("Error menyimpan XML:", err)
// 		return
// 	}

// 	fmt.Println("Data mahasiswa telah disimpan dalam format JSON dan XML.")
// }
