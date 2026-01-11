// package main

// import (
// 	"fmt"
// )

// // Struct untuk menyimpan informasi karyawan
// type Karyawan struct {
// 	Nama   string
// 	Gaji   float64
// 	Posisi string
// }

// // Fungsi untuk mengubah gaji karyawan menggunakan pointer
// func ubahGaji(k *Karyawan, gajiBaru float64) {
// 	k.Gaji = gajiBaru
// }

// func main() {
// 	// Membuat data karyawan
// 	karyawan1 := Karyawan{
// 		Nama:   "Budi",
// 		Gaji:   5000000,
// 		Posisi: "Software Engineer",
// 	}

// 	// Menampilkan informasi karyawan sebelum perubahan
// 	fmt.Println("Sebelum perubahan:")
// 	fmt.Printf("Nama: %s, Gaji: %.2f, Posisi: %s\n", karyawan1.Nama, karyawan1.Gaji, karyawan1.Posisi)

// 	// Mengubah gaji karyawan menggunakan fungsi ubahGaji
// 	ubahGaji(&karyawan1, 7000000)

// 	// Menampilkan informasi karyawan setelah perubahan
// 	fmt.Println("\nSetelah perubahan:")
// 	fmt.Printf("Nama: %s, Gaji: %.2f, Posisi: %s\n", karyawan1.Nama, karyawan1.Gaji, karyawan1.Posisi)
// }


package main

import (
	"fmt"
)

// Struct untuk menyimpan informasi karyawan
type Karyawan struct {
	Nama   string
	Gaji   float64
	Posisi string
}

// Fungsi untuk mengubah gaji karyawan menggunakan pointer
func ubahGaji(k *Karyawan, gajiBaru float64) {
	k.Gaji = gajiBaru
}

// Fungsi untuk mengubah posisi karyawan menggunakan pointer
func ubahPosisi(k *Karyawan, posisiBaru string) {
	k.Posisi = posisiBaru
}

func main() {
	// Membuat data karyawan dengan input dari pengguna
	var nama string
	var gaji float64
	var posisi string

	fmt.Print("Masukkan nama karyawan: ")
	fmt.Scanln(&nama)
	fmt.Print("Masukkan gaji karyawan: ")
	fmt.Scanln(&gaji)
	fmt.Print("Masukkan posisi karyawan: ")
	fmt.Scanln(&posisi)

	karyawan1 := Karyawan{
		Nama:   nama,
		Gaji:   gaji,
		Posisi: posisi,
	}

	// Menampilkan informasi karyawan sebelum perubahan
	fmt.Println("\nSebelum perubahan:")
	fmt.Printf("Nama: %s, Gaji: %.2f, Posisi: %s\n", karyawan1.Nama, karyawan1.Gaji, karyawan1.Posisi)

	// Mengubah gaji karyawan
	var gajiBaru float64
	fmt.Print("Masukkan gaji baru: ")
	fmt.Scanln(&gajiBaru)
	ubahGaji(&karyawan1, gajiBaru)

	// Mengubah posisi karyawan
	var posisiBaru string
	fmt.Print("Masukkan posisi baru: ")
	fmt.Scanln(&posisiBaru)
	ubahPosisi(&karyawan1, posisiBaru)

	// Menampilkan informasi karyawan setelah perubahan
	fmt.Println("\nSetelah perubahan:")
	fmt.Printf("Nama: %s, Gaji: %.2f, Posisi: %s\n", karyawan1.Nama, karyawan1.Gaji, karyawan1.Posisi)
}
