package main

import (
	"bufio"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"strings"
)

// Struct Mahasiswa
type Mahasiswa struct {
	Nama    string `json:"nama" xml:"nama"`
	Usia    int    `json:"usia" xml:"usia"`
	Jurusan string `json:"jurusan" xml:"jurusan"`
}

func main() {
	// Buka file teks mahasiswa.txt
	file, err := os.Open("mahasiswa.txt")
	if err != nil {
		fmt.Println("Error membuka file:", err)
		return
	}
	defer file.Close()

	// Parsing file teks ke struct Mahasiswa
	scanner := bufio.NewScanner(file)
	mahasiswa := Mahasiswa{}

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, ": ", 2)
		if len(parts) != 2 {
			fmt.Println("Format baris tidak valid:", line)
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// Masukkan data ke struct
		switch key {
		case "Nama":
			mahasiswa.Nama = value
		case "Usia":
			fmt.Sscanf(value, "%d", &mahasiswa.Usia)
		case "Jurusan":
			mahasiswa.Jurusan = value
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error membaca file:", err)
		return
	}

	// Serialisasi ke JSON
	jsonData, err := json.MarshalIndent(mahasiswa, "", "  ")
	if err != nil {
		fmt.Println("Error serialisasi ke JSON:", err)
		return
	}

	// Simpan JSON ke file mahasiswa.json
	err = os.WriteFile("mahasiswa.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Error menyimpan JSON:", err)
		return
	}
	fmt.Println("Data berhasil disimpan ke mahasiswa.json")

	// Serialisasi ke XML
	xmlData, err := xml.MarshalIndent(mahasiswa, "", "  ")
	if err != nil {
		fmt.Println("Error serialisasi ke XML:", err)
		return
	}

	// Simpan XML ke file mahasiswa.xml
	err = os.WriteFile("mahasiswa.xml", xmlData, 0644)
	if err != nil {
		fmt.Println("Error menyimpan XML:", err)
		return
	}
	fmt.Println("Data berhasil disimpan ke mahasiswa.xml")
}
