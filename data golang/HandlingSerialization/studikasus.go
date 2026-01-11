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
	var mahasiswa Mahasiswa

	// Menerima input data dari pengguna
	fmt.Println("Masukkan data mahasiswa:")

	// Nama
	fmt.Print("Nama: ")
	fmt.Scanln(&mahasiswa.Nama)

	// Usia
	fmt.Print("Usia: ")
	fmt.Scanln(&mahasiswa.Usia)

	// Jurusan
	fmt.Print("Jurusan: ")
	fmt.Scanln(&mahasiswa.Jurusan)

	// Menyimpan data dalam file teks (mahasiswa.txt)
	saveToTextFile(mahasiswa)

	// Pilihan format penyimpanan (JSON atau XML)
	var format string
	fmt.Print("Pilih format penyimpanan (json/xml): ")
	fmt.Scanln(&format)

	switch strings.ToLower(format) {
	case "json":
		saveToJSON(mahasiswa)
	case "xml":
		saveToXML(mahasiswa)
	default:
		fmt.Println("Format tidak valid.")
	}

	// Menanyakan apakah ingin membaca file JSON atau XML
	var readFormat string
	fmt.Print("Apakah Anda ingin membaca file JSON atau XML? (json/xml/tidak): ")
	fmt.Scanln(&readFormat)

	switch strings.ToLower(readFormat) {
	case "json":
		readFromFile("mahasiswa.json")
	case "xml":
		readFromFile("mahasiswa.xml")
	case "tidak":
		fmt.Println("Terima kasih, aplikasi selesai.")
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}

// Fungsi untuk menyimpan data mahasiswa ke file teks
func saveToTextFile(mahasiswa Mahasiswa) {
	file, err := os.Create("mahasiswa.txt")
	if err != nil {
		fmt.Println("Error membuat file:", err)
		return
	}
	defer file.Close()

	// Menulis data ke file teks
	file.WriteString(fmt.Sprintf("Nama: %s\nUsia: %d\nJurusan: %s\n", mahasiswa.Nama, mahasiswa.Usia, mahasiswa.Jurusan))
	fmt.Println("Data berhasil disimpan ke mahasiswa.txt")
}

// Fungsi untuk menyimpan data mahasiswa ke file JSON
func saveToJSON(mahasiswa Mahasiswa) {
	jsonData, err := json.MarshalIndent(mahasiswa, "", "  ")
	if err != nil {
		fmt.Println("Error serialisasi ke JSON:", err)
		return
	}

	err = os.WriteFile("mahasiswa.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Error menyimpan JSON:", err)
		return
	}
	fmt.Println("Data berhasil disimpan ke mahasiswa.json")
}

// Fungsi untuk menyimpan data mahasiswa ke file XML
func saveToXML(mahasiswa Mahasiswa) {
	xmlData, err := xml.MarshalIndent(mahasiswa, "", "  ")
	if err != nil {
		fmt.Println("Error serialisasi ke XML:", err)
		return
	}

	err = os.WriteFile("mahasiswa.xml", xmlData, 0644)
	if err != nil {
		fmt.Println("Error menyimpan XML:", err)
		return
	}
	fmt.Println("Data berhasil disimpan ke mahasiswa.xml")
}

// Fungsi untuk membaca dan menampilkan file JSON atau XML
func readFromFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error membuka file:", err)
		return
	}
	defer file.Close()

	// Membaca seluruh isi file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error membaca file:", err)
	}
}
