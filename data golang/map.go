// package main

// import (
// 	"fmt"
// )

// func main() {
// 	// Map untuk menyimpan data mahasiswa (nama sebagai kunci, nilai sebagai nilai)
// 	mahasiswa := make(map[string]float64)

// 	for {
// 		// Menampilkan menu pilihan
// 		fmt.Println("\nMenu:")
// 		fmt.Println("1. Tambah Mahasiswa")
// 		fmt.Println("2. Hapus Mahasiswa")
// 		fmt.Println("3. Update Nilai Mahasiswa")
// 		fmt.Println("4. Tampilkan Semua Mahasiswa")
// 		fmt.Println("5. Keluar")
// 		fmt.Print("Pilih opsi: ")

// 		var choice int
// 		fmt.Scanln(&choice)

// 		switch choice {
// 		case 1:
// 			// Menambahkan mahasiswa baru
// 			var nama string
// 			var nilai float64
// 			fmt.Print("Masukkan nama mahasiswa: ")
// 			fmt.Scanln(&nama)
// 			fmt.Print("Masukkan nilai mahasiswa: ")
// 			fmt.Scanln(&nilai)
// 			mahasiswa[nama] = nilai
// 			fmt.Println("Mahasiswa berhasil ditambahkan.")

// 		case 2:
// 			// Menghapus data mahasiswa
// 			var nama string
// 			fmt.Print("Masukkan nama mahasiswa yang akan dihapus: ")
// 			fmt.Scanln(&nama)
// 			if _, exists := mahasiswa[nama]; exists {
// 				delete(mahasiswa, nama)
// 				fmt.Println("Data mahasiswa berhasil dihapus.")
// 			} else {
// 				fmt.Println("Mahasiswa tidak ditemukan.")
// 			}

// 		case 3:
// 			// Mengupdate nilai mahasiswa
// 			var nama string
// 			var nilaiBaru float64
// 			fmt.Print("Masukkan nama mahasiswa yang akan diupdate nilainya: ")
// 			fmt.Scanln(&nama)
// 			if _, exists := mahasiswa[nama]; exists {
// 				fmt.Print("Masukkan nilai baru: ")
// 				fmt.Scanln(&nilaiBaru)
// 				mahasiswa[nama] = nilaiBaru
// 				fmt.Println("Nilai mahasiswa berhasil diupdate.")
// 			} else {
// 				fmt.Println("Mahasiswa tidak ditemukan.")
// 			}

// 		case 4:
// 			// Menampilkan semua data mahasiswa
// 			if len(mahasiswa) == 0 {
// 				fmt.Println("Tidak ada data mahasiswa.")
// 			} else {
// 				fmt.Println("Daftar Mahasiswa:")
// 				for nama, nilai := range mahasiswa {
// 					fmt.Printf("Nama: %s, Nilai: %.2f\n", nama, nilai)
// 				}
// 			}

// 		case 5:
// 			// Keluar dari program
// 			fmt.Println("Terima kasih telah menggunakan program ini!")
// 			return

// 		default:
// 			fmt.Println("Pilihan tidak valid, silakan coba lagi.")
// 		}
// 	}
// }

package main

import (
	"fmt"
)

func main() {
	// Map untuk menyimpan data mahasiswa (nama sebagai kunci, nilai sebagai nilai)
	mahasiswa := make(map[string]float64)

	for {
		// Menampilkan menu pilihan
		fmt.Println("\nMenu:")
		fmt.Println("1. Tambah Mahasiswa")
		fmt.Println("2. Hapus Mahasiswa")
		fmt.Println("3. Update Nilai Mahasiswa")
		fmt.Println("4. Tampilkan Semua Mahasiswa")
		fmt.Println("5. Keluar")
		fmt.Print("Pilih opsi: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			// Menambahkan mahasiswa baru
			var nama string
			var nilai float64
			fmt.Print("Masukkan nama mahasiswa: ")
			fmt.Scanln(&nama)
			fmt.Print("Masukkan nilai mahasiswa: ")
			fmt.Scanln(&nilai)
			mahasiswa[nama] = nilai
			fmt.Println("Mahasiswa berhasil ditambahkan.")

		case 2:
			// Menghapus data mahasiswa
			var nama string
			fmt.Print("Masukkan nama mahasiswa yang akan dihapus: ")
			fmt.Scanln(&nama)
			if _, exists := mahasiswa[nama]; exists {
				delete(mahasiswa, nama)
				fmt.Println("Data mahasiswa berhasil dihapus.")
			} else {
				fmt.Println("Mahasiswa tidak ditemukan.")
			}

		case 3:
			// Mengupdate nilai mahasiswa
			var nama string
			var nilaiBaru float64
			fmt.Print("Masukkan nama mahasiswa yang akan diupdate nilainya: ")
			fmt.Scanln(&nama)
			if _, exists := mahasiswa[nama]; exists {
				fmt.Print("Masukkan nilai baru: ")
				fmt.Scanln(&nilaiBaru)
				mahasiswa[nama] = nilaiBaru
				fmt.Println("Nilai mahasiswa berhasil diupdate.")
			} else {
				fmt.Println("Mahasiswa tidak ditemukan.")
			}

		case 4:
			// Menampilkan semua data mahasiswa
			if len(mahasiswa) == 0 {
				fmt.Println("Tidak ada data mahasiswa.")
			} else {
				fmt.Println("Daftar Mahasiswa:")
				for nama, nilai := range mahasiswa {
					fmt.Printf("Nama: %s, Nilai: %.2f\n", nama, nilai)
				}
			}

		case 5:
			// Keluar dari program
			fmt.Println("Terima kasih telah menggunakan program ini!")
			return

		default:
			fmt.Println("Pilihan tidak valid, silakan coba lagi.")
		}
	}
}
