package main

import (
	"bufio"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"strings"
)

// Struct untuk menyimpan data mahasiswa
type Mahasiswa struct {
	Nama   string `json:"nama" xml:"nama"`
	Usia   int    `json:"usia" xml:"usia"`
	Jurusan string `json:"jurusan" xml:"jurusan"`
}

// Fungsi untuk mengambil input data mahasiswa dari pengguna
func inputDataMahasiswa() Mahasiswa {
	var mhs Mahasiswa
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan nama mahasiswa: ")
	nama, _ := reader.ReadString('\n')
	mhs.Nama = strings.TrimSpace(nama)

	fmt.Print("Masukkan usia mahasiswa: ")
	fmt.Scan(&mhs.Usia)

	fmt.Print("Masukkan jurusan mahasiswa: ")
	jurusan, _ := reader.ReadString('\n')
	mhs.Jurusan = strings.TrimSpace(jurusan)

	return mhs
}

// Fungsi untuk menyimpan data dalam format JSON
func simpanDataJSON(data Mahasiswa) {
	file, err := os.OpenFile("mahasiswa.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Gagal membuka file JSON:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(data); err != nil {
		fmt.Println("Gagal menyimpan data JSON:", err)
	} else {
		fmt.Println("Data berhasil disimpan dalam format JSON (mahasiswa.json).")
	}
}

// Fungsi untuk menyimpan data dalam format XML
func simpanDataXML(data Mahasiswa) {
	file, err := os.OpenFile("mahasiswa.xml", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Gagal membuka file XML:", err)
		return
	}
	defer file.Close()

	encoder := xml.NewEncoder(file)
	file.WriteString(xml.Header) // Menulis header XML
	if err := encoder.Encode(data); err != nil {
		fmt.Println("Gagal menyimpan data XML:", err)
	} else {
		fmt.Println("Data berhasil disimpan dalam format XML (mahasiswa.xml).")
	}
}

// Fungsi untuk membaca dan menampilkan data dari file JSON
func bacaDataJSON() {
	file, err := os.Open("mahasiswa.json")
	if err != nil {
		fmt.Println("Gagal membuka file JSON:", err)
		return
	}
	defer file.Close()

	var data Mahasiswa
	decoder := json.NewDecoder(file)
	for decoder.More() {
		if err := decoder.Decode(&data); err == nil {
			fmt.Printf("Nama: %s, Usia: %d, Jurusan: %s\n", data.Nama, data.Usia, data.Jurusan)
		}
	}
}

// Fungsi untuk membaca dan menampilkan data dari file XML
func bacaDataXML() {
	file, err := os.Open("mahasiswa.xml")
	if err != nil {
		fmt.Println("Gagal membuka file XML:", err)
		return
	}
	defer file.Close()

	var data Mahasiswa
	decoder := xml.NewDecoder(file)
	for {
		if err := decoder.Decode(&data); err != nil {
			break
		}
		fmt.Printf("Nama: %s, Usia: %d, Jurusan: %s\n", data.Nama, data.Usia, data.Jurusan)
	}
}

func main() {
	// Pilihan menu untuk pengguna
	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Input data mahasiswa")
		fmt.Println("2. Tampilkan data dari file JSON")
		fmt.Println("3. Tampilkan data dari file XML")
		fmt.Println("4. Keluar")
		fmt.Print("Pilih opsi: ")

		var pilihan int
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			data := inputDataMahasiswa()
			fmt.Print("Pilih format penyimpanan (json/xml): ")
			var format string
			fmt.Scan(&format)
			if format == "json" {
				simpanDataJSON(data)
			} else if format == "xml" {
				simpanDataXML(data)
			} else {
				fmt.Println("Format tidak dikenal, data tidak disimpan.")
			}

		case 2:
			fmt.Println("Menampilkan data dari file JSON:")
			bacaDataJSON()

		case 3:
			fmt.Println("Menampilkan data dari file XML:")
			bacaDataXML()

		case 4:
			fmt.Println("Keluar dari program.")
			return

		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}
