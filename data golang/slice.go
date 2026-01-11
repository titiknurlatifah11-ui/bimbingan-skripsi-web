// package main

// import (
// 	"fmt"
// )

// // Struct untuk menyimpan informasi barang
// type Item struct {
// 	Name  string
// 	Price float64
// }

// func main() {
// 	// Slice untuk menyimpan daftar barang yang akan dibeli
// 	var cart []Item

// 	// Fungsi untuk menampilkan menu pilihan
// 	for {
// 		fmt.Println("\nMenu:")
// 		fmt.Println("1. Tambah Barang")
// 		fmt.Println("2. Hapus Barang")
// 		fmt.Println("3. Tampilkan Daftar Barang")
// 		fmt.Println("4. Keluar")
// 		fmt.Print("Pilih opsi: ")

// 		var choice int
// 		fmt.Scanln(&choice)

// 		switch choice {
// 		case 1:
// 			// Menambahkan barang ke keranjang
// 			var name string
// 			var price float64
// 			fmt.Print("Masukkan nama barang: ")
// 			fmt.Scanln(&name)
// 			fmt.Print("Masukkan harga barang: ")
// 			fmt.Scanln(&price)
// 			cart = append(cart, Item{Name: name, Price: price})
// 			fmt.Println("Barang berhasil ditambahkan.")

// 		case 2:
// 			// Menghapus barang dari keranjang
// 			if len(cart) == 0 {
// 				fmt.Println("Keranjang kosong, tidak ada barang untuk dihapus.")
// 			} else {
// 				fmt.Println("Daftar Barang:")
// 				for i, item := range cart {
// 					fmt.Printf("%d. %s - Rp %.2f\n", i+1, item.Name, item.Price)
// 				}
// 				var index int
// 				fmt.Print("Masukkan nomor barang yang akan dihapus: ")
// 				fmt.Scanln(&index)

// 				// Validasi nomor barang yang dimasukkan
// 				if index > 0 && index <= len(cart) {
// 					cart = append(cart[:index-1], cart[index:]...)
// 					fmt.Println("Barang berhasil dihapus.")
// 				} else {
// 					fmt.Println("Nomor barang tidak valid.")
// 				}
// 			}

// 		case 3:
// 			// Menampilkan daftar barang di keranjang
// 			if len(cart) == 0 {
// 				fmt.Println("Keranjang kosong.")
// 			} else {
// 				fmt.Println("Daftar Barang di Keranjang:")
// 				for i, item := range cart {
// 					fmt.Printf("%d. %s - Rp %.2f\n", i+1, item.Name, item.Price)
// 				}
// 			}

// 		case 4:
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

// Struct untuk menyimpan informasi barang
type Item struct {
	Name  string
	Price float64
}

func main() {
	// Slice untuk menyimpan daftar barang yang akan dibeli
	var cart []Item

	// Fungsi untuk menampilkan menu pilihan
	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Tambah Barang")
		fmt.Println("2. Hapus Barang")
		fmt.Println("3. Tampilkan Daftar Barang")
		fmt.Println("4. Keluar")
		fmt.Print("Pilih opsi: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			// Menambahkan barang ke keranjang
			var name string
			var price float64
			fmt.Print("Masukkan nama barang: ")
			fmt.Scanln(&name)
			fmt.Print("Masukkan harga barang: ")
			fmt.Scanln(&price)
			cart = append(cart, Item{Name: name, Price: price})
			fmt.Println("Barang berhasil ditambahkan.")

		case 2:
			// Menghapus barang dari keranjang
			if len(cart) == 0 {
				fmt.Println("Keranjang kosong, tidak ada barang untuk dihapus.")
			} else {
				fmt.Println("Daftar Barang:")
				for i, item := range cart {
					fmt.Printf("%d. %s - Rp %.2f\n", i+1, item.Name, item.Price)
				}
				var index int
				fmt.Print("Masukkan nomor barang yang akan dihapus: ")
				fmt.Scanln(&index)

				// Validasi nomor barang yang dimasukkan
				if index > 0 && index <= len(cart) {
					cart = append(cart[:index-1], cart[index:]...)
					fmt.Println("Barang berhasil dihapus.")
				} else {
					fmt.Println("Nomor barang tidak valid.")
				}
			}

		case 3:
			// Menampilkan daftar barang di keranjang dan total harga
			if len(cart) == 0 {
				fmt.Println("Keranjang kosong.")
			} else {
				fmt.Println("Daftar Barang di Keranjang:")
				var total float64
				for i, item := range cart {
					fmt.Printf("%d. %s - Rp %.2f\n", i+1, item.Name, item.Price)
					total += item.Price
				}
				fmt.Printf("Total harga barang: Rp %.2f\n", total)
			}

		case 4:
			// Konfirmasi keluar dari program
			var confirm string
			fmt.Print("Apakah Anda yakin ingin keluar? (y/n): ")
			fmt.Scanln(&confirm)
			if confirm == "y" || confirm == "Y" {
				fmt.Println("Terima kasih telah menggunakan program ini!")
				return
			}

		default:
			fmt.Println("Pilihan tidak valid, silakan coba lagi.")
		}
	}
}
