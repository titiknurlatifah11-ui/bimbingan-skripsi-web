//Latihan1
// package main

// import (
// 	"fmt"
// )

// // Struct untuk mendefinisikan Karyawan
// type Karyawan struct {
// 	ID   int
// 	Nama string
// 	Usia int
// 	Gaji float64
// }

// // Deklarasi global slice untuk menyimpan data karyawan
// var daftarKaryawan []Karyawan
// var karyawanMap = make(map[string]*Karyawan) // Map untuk pencarian berdasarkan nama

// // Fungsi untuk menambah karyawan
// func tambahKaryawan(id int, nama string, usia int, gaji float64) {
// 	karyawanBaru := Karyawan{ID: id, Nama: nama, Usia: usia, Gaji: gaji}
// 	daftarKaryawan = append(daftarKaryawan, karyawanBaru)
// 	karyawanMap[nama] = &karyawanBaru // Menyimpan karyawan ke map berdasarkan nama
// 	fmt.Println("Karyawan berhasil ditambahkan!")
// }

// // Fungsi untuk menampilkan semua karyawan
// func tampilkanKaryawan() {
// 	if len(daftarKaryawan) == 0 {
// 		fmt.Println("Tidak ada karyawan yang terdaftar.")
// 		return
// 	}
// 	fmt.Println("Daftar Karyawan:")
// 	for _, karyawan := range daftarKaryawan {
// 		fmt.Printf("ID: %d, Nama: %s, Usia: %d, Gaji: %.2f\n", karyawan.ID, karyawan.Nama, karyawan.Usia, karyawan.Gaji)
// 	}
// }

// // Fungsi untuk mencari karyawan berdasarkan nama
// func cariKaryawan(nama string) {
// 	karyawan, ok := karyawanMap[nama]
// 	if ok {
// 		fmt.Printf("Karyawan ditemukan: ID: %d, Nama: %s, Usia: %d, Gaji: %.2f\n", karyawan.ID, karyawan.Nama, karyawan.Usia, karyawan.Gaji)
// 	} else {
// 		fmt.Println("Karyawan tidak ditemukan.")
// 	}
// }

// // Fungsi untuk mengubah data karyawan menggunakan pointer
// func ubahDataKaryawan(nama string, usiaBaru int, gajiBaru float64) {
// 	karyawan, ok := karyawanMap[nama]
// 	if ok {
// 		karyawan.Usia = usiaBaru
// 		karyawan.Gaji = gajiBaru
// 		fmt.Println("Data karyawan berhasil diubah!")
// 	} else {
// 		fmt.Println("Karyawan tidak ditemukan.")
// 	}
// }

// // Fungsi untuk menghapus karyawan berdasarkan nama
// func hapusKaryawan(nama string) {
// 	index := -1
// 	for i, karyawan := range daftarKaryawan {
// 		if karyawan.Nama == nama {
// 			index = i
// 			break
// 		}
// 	}

// 	if index != -1 {
// 		daftarKaryawan = append(daftarKaryawan[:index], daftarKaryawan[index+1:]...) // Menghapus dari slice
// 		delete(karyawanMap, nama)                                                     // Menghapus dari map
// 		fmt.Println("Karyawan berhasil dihapus!")
// 	} else {
// 		fmt.Println("Karyawan tidak ditemukan.")
// 	}
// }

// // Main function untuk menjalankan program
// func main() {
// 	var pilihan int
// 	for {
// 		fmt.Println("\nMenu:")
// 		fmt.Println("1. Tambah Karyawan")
// 		fmt.Println("2. Tampilkan Karyawan")
// 		fmt.Println("3. Cari Karyawan Berdasarkan Nama")
// 		fmt.Println("4. Ubah Data Karyawan")
// 		fmt.Println("5. Hapus Karyawan")
// 		fmt.Println("6. Keluar")
// 		fmt.Print("Pilih opsi: ")
// 		fmt.Scan(&pilihan)

// 		switch pilihan {
// 		case 1:
// 			var id, usia int
// 			var nama string
// 			var gaji float64
// 			fmt.Print("Masukkan ID: ")
// 			fmt.Scan(&id)
// 			fmt.Print("Masukkan Nama: ")
// 			fmt.Scan(&nama)
// 			fmt.Print("Masukkan Usia: ")
// 			fmt.Scan(&usia)
// 			fmt.Print("Masukkan Gaji: ")
// 			fmt.Scan(&gaji)
// 			tambahKaryawan(id, nama, usia, gaji)
// 		case 2:
// 			tampilkanKaryawan()
// 		case 3:
// 			var nama string
// 			fmt.Print("Masukkan Nama Karyawan yang Dicari: ")
// 			fmt.Scan(&nama)
// 			cariKaryawan(nama)
// 		case 4:
// 			var nama string
// 			var usia int
// 			var gaji float64
// 			fmt.Print("Masukkan Nama Karyawan yang Ingin Diubah: ")
// 			fmt.Scan(&nama)
// 			fmt.Print("Masukkan Usia Baru: ")
// 			fmt.Scan(&usia)
// 			fmt.Print("Masukkan Gaji Baru: ")
// 			fmt.Scan(&gaji)
// 			ubahDataKaryawan(nama, usia, gaji)
// 		case 5:
// 			var nama string
// 			fmt.Print("Masukkan Nama Karyawan yang Ingin Dihapus: ")
// 			fmt.Scan(&nama)
// 			hapusKaryawan(nama)
// 		case 6:
// 			fmt.Println("Keluar dari program.")
// 			return
// 		default:
// 			fmt.Println("Pilihan tidak valid.")
// 		}
// 	}
// }


//Latihan2
// package main

// import (
// 	"fmt"
// )

// // Struct untuk mendefinisikan Barang
// type Barang struct {
// 	Nama   string
// 	Harga  float64
// 	Stok   int
// }

// // Slice untuk menyimpan daftar barang
// var daftarBarang []Barang

// // Map untuk pencarian barang berdasarkan nama
// var barangMap = make(map[string]*Barang)

// // Fungsi untuk menambah barang ke inventaris
// func tambahBarang(nama string, harga float64, stok int) {
// 	barangBaru := Barang{Nama: nama, Harga: harga, Stok: stok}
// 	daftarBarang = append(daftarBarang, barangBaru)
// 	barangMap[nama] = &barangBaru // Simpan dalam map berdasarkan nama
// 	fmt.Println("Barang berhasil ditambahkan!")
// }

// // Fungsi untuk menampilkan semua barang di inventaris
// func tampilkanBarang() {
// 	if len(daftarBarang) == 0 {
// 		fmt.Println("Tidak ada barang yang tersedia.")
// 		return
// 	}
// 	fmt.Println("Daftar Barang di Inventaris:")
// 	for _, barang := range daftarBarang {
// 		fmt.Printf("Nama: %s, Harga: %.2f, Stok: %d\n", barang.Nama, barang.Harga, barang.Stok)
// 	}
// }

// // Fungsi untuk mencari barang berdasarkan nama
// func cariBarang(nama string) {
// 	barang, ok := barangMap[nama]
// 	if ok {
// 		fmt.Printf("Barang ditemukan: Nama: %s, Harga: %.2f, Stok: %d\n", barang.Nama, barang.Harga, barang.Stok)
// 	} else {
// 		fmt.Println("Barang tidak ditemukan.")
// 	}
// }

// // Fungsi untuk mengubah jumlah stok barang menggunakan pointer
// func ubahStokBarang(nama string, stokBaru int) {
// 	barang, ok := barangMap[nama]
// 	if ok {
// 		barang.Stok = stokBaru
// 		fmt.Println("Stok barang berhasil diubah!")
// 	} else {
// 		fmt.Println("Barang tidak ditemukan.")
// 	}
// }

// // Fungsi untuk menghapus barang dari inventaris
// func hapusBarang(nama string) {
// 	index := -1
// 	for i, barang := range daftarBarang {
// 		if barang.Nama == nama {
// 			index = i
// 			break
// 		}
// 	}

// 	if index != -1 {
// 		daftarBarang = append(daftarBarang[:index], daftarBarang[index+1:]...) // Menghapus dari slice
// 		delete(barangMap, nama)                                                // Menghapus dari map
// 		fmt.Println("Barang berhasil dihapus!")
// 	} else {
// 		fmt.Println("Barang tidak ditemukan.")
// 	}
// }

// // Main function untuk menjalankan program
// func main() {
// 	var pilihan int
// 	for {
// 		fmt.Println("\nMenu:")
// 		fmt.Println("1. Tambah Barang")
// 		fmt.Println("2. Tampilkan Semua Barang")
// 		fmt.Println("3. Cari Barang Berdasarkan Nama")
// 		fmt.Println("4. Ubah Stok Barang")
// 		fmt.Println("5. Hapus Barang")
// 		fmt.Println("6. Keluar")
// 		fmt.Print("Pilih opsi: ")
// 		fmt.Scan(&pilihan)

// 		switch pilihan {
// 		case 1:
// 			var nama string
// 			var harga float64
// 			var stok int
// 			fmt.Print("Masukkan Nama Barang: ")
// 			fmt.Scan(&nama)
// 			fmt.Print("Masukkan Harga Barang: ")
// 			fmt.Scan(&harga)
// 			fmt.Print("Masukkan Jumlah Stok Barang: ")
// 			fmt.Scan(&stok)
// 			tambahBarang(nama, harga, stok)
// 		case 2:
// 			tampilkanBarang()
// 		case 3:
// 			var nama string
// 			fmt.Print("Masukkan Nama Barang yang Dicari: ")
// 			fmt.Scan(&nama)
// 			cariBarang(nama)
// 		case 4:
// 			var nama string
// 			var stok int
// 			fmt.Print("Masukkan Nama Barang yang Ingin Diubah Stoknya: ")
// 			fmt.Scan(&nama)
// 			fmt.Print("Masukkan Jumlah Stok Baru: ")
// 			fmt.Scan(&stok)
// 			ubahStokBarang(nama, stok)
// 		case 5:
// 			var nama string
// 			fmt.Print("Masukkan Nama Barang yang Ingin Dihapus: ")
// 			fmt.Scan(&nama)
// 			hapusBarang(nama)
// 		case 6:
// 			fmt.Println("Keluar dari program.")
// 			return
// 		default:
// 			fmt.Println("Pilihan tidak valid.")
// 		}
// 	}
// }

