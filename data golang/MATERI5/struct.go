//4.STRUCT
// package main

// import "fmt"

// // Struct untuk menyimpan informasi orang
// type Person struct {
// 	Name string
// 	Age  int
// }

// // Fungsi untuk memperbarui usia orang menggunakan pointer
// func updateAge(p *Person, newAge int) {
// 	p.Age = newAge
// }

// func main() {
// 	person := Person{Name: "John", Age: 25}
// 	fmt.Println("Sebelum:", person)
// 	updateAge(&person, 30)
// 	fmt.Println("Setelah:", person)
// }

//modifikasi
// package main

// import "fmt"

// // Struct untuk menyimpan informasi orang
// type Person struct {
// 	Name string
// 	Age  int
// }

// // Fungsi untuk memperbarui usia dan nama orang menggunakan pointer
// func updatePerson(p *Person, newName string, newAge int) {
// 	p.Name = newName
// 	p.Age = newAge
// }

// func main() {
// 	// Membuat objek Person dengan data awal
// 	person := Person{Name: "John", Age: 25}
// 	fmt.Println("Sebelum:", person)

// 	// Input dari pengguna untuk nama dan usia baru
// 	var newName string
// 	var newAge int
// 	fmt.Print("Masukkan nama baru: ")
// 	fmt.Scanln(&newName)
// 	fmt.Print("Masukkan usia baru: ")
// 	fmt.Scanln(&newAge)

// 	// Memperbarui informasi orang
// 	updatePerson(&person, newName, newAge)
// 	fmt.Println("Setelah:", person)
// }

//STUDIKASUS struct
// package main

// import (
// 	"fmt"
// )

// // Struct untuk menyimpan informasi buku
// type Buku struct {
// 	Judul     string
// 	Pengarang string
// 	Tahun     int
// }

// // Fungsi utama
// func main() {
// 	// Slice untuk menyimpan daftar buku
// 	var perpustakaan []Buku
// 	for {
// 		// Menampilkan menu pilihan
// 		fmt.Println("\nMenu:")
// 		fmt.Println("1. Tambah Buku")
// 		fmt.Println("2. Tampilkan Semua Buku")
// 		fmt.Println("3. Update Informasi Buku")
// 		fmt.Println("4. Keluar")
// 		fmt.Print("Pilih opsi: ")

// 		var pilihan int
// 		fmt.Scanln(&pilihan)

// 		switch pilihan {
// 		case 1:
// 			// Panggil fungsi untuk menambah buku
// 			perpustakaan = tambahBuku(perpustakaan)
// 		case 2:
// 			// Panggil fungsi untuk menampilkan semua buku
// 			tampilkanBuku(perpustakaan)
// 		case 3:
// 			// Panggil fungsi untuk mengupdate informasi buku
// 			perpustakaan = updateBuku(perpustakaan)
// 		case 4:
// 			// Keluar dari program
// 			fmt.Println("Terima kasih telah menggunakan program ini!")
// 			return
// 		default:
// 			fmt.Println("Pilihan tidak valid, silakan coba lagi.")
// 		}
// 	}
// }

// // Fungsi untuk menambah buku baru
// func tambahBuku(perpustakaan []Buku) []Buku {
// 	var judul, pengarang string
// 	var tahun int

// 	fmt.Print("Masukkan judul buku: ")
// 	fmt.Scanln(&judul)
// 	fmt.Print("Masukkan nama pengarang: ")
// 	fmt.Scanln(&pengarang)
// 	fmt.Print("Masukkan tahun terbit: ")
// 	fmt.Scanln(&tahun)

// 	// Tambahkan buku baru ke slice perpustakaan
// 	perpustakaan = append(perpustakaan, Buku{Judul: judul, Pengarang: pengarang, Tahun: tahun})
// 	fmt.Println("Buku berhasil ditambahkan.")
// 	return perpustakaan
// }

// // Fungsi untuk menampilkan semua buku
// func tampilkanBuku(perpustakaan []Buku) {
// 	if len(perpustakaan) == 0 {
// 		fmt.Println("Tidak ada buku di perpustakaan.")
// 	} else {
// 		fmt.Println("Daftar Buku:")
// 		for i, buku := range perpustakaan {
// 			fmt.Printf("%d. Judul: %s, Pengarang: %s, Tahun Terbit: %d\n", i+1, buku.Judul, buku.Pengarang, buku.Tahun)
// 		}
// 	}
// }

// // Fungsi untuk mengupdate informasi buku
// func updateBuku(perpustakaan []Buku) []Buku {
// 	if len(perpustakaan) == 0 {
// 		fmt.Println("Tidak ada buku untuk diupdate.")
// 	} else {
// 		// Tampilkan daftar buku
// 		tampilkanBuku(perpustakaan)

// 		// Memilih buku yang akan diupdate
// 		var index int
// 		fmt.Print("Masukkan nomor buku yang akan diupdate: ")
// 		fmt.Scanln(&index)

// 		// Validasi nomor buku
// 		if index > 0 && index <= len(perpustakaan) {
// 			var judulBaru, pengarangBaru string
// 			var tahunBaru int

// 			fmt.Print("Masukkan judul baru (tekan Enter jika tidak ingin mengubah): ")
// 			fmt.Scanln(&judulBaru)
// 			fmt.Print("Masukkan nama pengarang baru (tekan Enter jika tidak ingin mengubah): ")
// 			fmt.Scanln(&pengarangBaru)
// 			fmt.Print("Masukkan tahun terbit baru (tekan 0 jika tidak ingin mengubah): ")
// 			fmt.Scanln(&tahunBaru)

// 			// Update informasi buku jika input diberikan
// 			if judulBaru != "" {
// 				perpustakaan[index-1].Judul = judulBaru
// 			}
// 			if pengarangBaru != "" {
// 				perpustakaan[index-1].Pengarang = pengarangBaru
// 			}
// 			if tahunBaru != 0 {
// 				perpustakaan[index-1].Tahun = tahunBaru
// 			}

// 			fmt.Println("Informasi buku berhasil diupdate.")
// 		} else {
// 			fmt.Println("Nomor buku tidak valid.")
// 		}
// 	}
// 	return perpustakaan
// }

//modifikasi STUDIKASUS struct
// package main

// import (
// 	"fmt"
// )

// // Struct untuk menyimpan informasi buku
// type Buku struct {
// 	Judul     string
// 	Pengarang string
// 	Tahun     int
// }

// // Fungsi utama
// func main() {
// 	// Slice untuk menyimpan daftar buku
// 	var perpustakaan []Buku
// 	for {
// 		// Menampilkan menu pilihan
// 		fmt.Println("\nMenu:")
// 		fmt.Println("1. Tambah Buku")
// 		fmt.Println("2. Tampilkan Semua Buku")
// 		fmt.Println("3. Update Informasi Buku")
// 		fmt.Println("4. Hapus Buku")
// 		fmt.Println("5. Keluar")
// 		fmt.Print("Pilih opsi: ")

// 		var pilihan int
// 		fmt.Scanln(&pilihan)

// 		switch pilihan {
// 		case 1:
// 			// Panggil fungsi untuk menambah buku
// 			perpustakaan = tambahBuku(perpustakaan)
// 		case 2:
// 			// Panggil fungsi untuk menampilkan semua buku
// 			tampilkanBuku(perpustakaan)
// 		case 3:
// 			// Panggil fungsi untuk mengupdate informasi buku
// 			perpustakaan = updateBuku(perpustakaan)
// 		case 4:
// 			// Panggil fungsi untuk menghapus buku
// 			perpustakaan = hapusBuku(perpustakaan)
// 		case 5:
// 			// Keluar dari program
// 			fmt.Println("Terima kasih telah menggunakan program ini!")
// 			return
// 		default:
// 			fmt.Println("Pilihan tidak valid, silakan coba lagi.")
// 		}
// 	}
// }

// // Fungsi untuk menambah buku baru
// func tambahBuku(perpustakaan []Buku) []Buku {
// 	var judul, pengarang string
// 	var tahun int

// 	fmt.Print("Masukkan judul buku: ")
// 	fmt.Scanln(&judul)
// 	fmt.Print("Masukkan nama pengarang: ")
// 	fmt.Scanln(&pengarang)
// 	fmt.Print("Masukkan tahun terbit: ")
// 	fmt.Scanln(&tahun)

// 	// Tambahkan buku baru ke slice perpustakaan
// 	perpustakaan = append(perpustakaan, Buku{Judul: judul, Pengarang: pengarang, Tahun: tahun})
// 	fmt.Println("Buku berhasil ditambahkan.")
// 	return perpustakaan
// }

// // Fungsi untuk menampilkan semua buku
// func tampilkanBuku(perpustakaan []Buku) {
// 	if len(perpustakaan) == 0 {
// 		fmt.Println("Tidak ada buku di perpustakaan.")
// 	} else {
// 		fmt.Println("Daftar Buku:")
// 		for i, buku := range perpustakaan {
// 			fmt.Printf("%d. Judul: %s, Pengarang: %s, Tahun Terbit: %d\n", i+1, buku.Judul, buku.Pengarang, buku.Tahun)
// 		}
// 	}
// }

// // Fungsi untuk mengupdate informasi buku
// func updateBuku(perpustakaan []Buku) []Buku {
// 	if len(perpustakaan) == 0 {
// 		fmt.Println("Tidak ada buku untuk diupdate.")
// 	} else {
// 		// Tampilkan daftar buku
// 		tampilkanBuku(perpustakaan)

// 		// Memilih buku yang akan diupdate
// 		var index int
// 		fmt.Print("Masukkan nomor buku yang akan diupdate: ")
// 		fmt.Scanln(&index)

// 		// Validasi nomor buku
// 		if index > 0 && index <= len(perpustakaan) {
// 			var judulBaru, pengarangBaru string
// 			var tahunBaru int

// 			fmt.Print("Masukkan judul baru (tekan Enter jika tidak ingin mengubah): ")
// 			fmt.Scanln(&judulBaru)
// 			fmt.Print("Masukkan nama pengarang baru (tekan Enter jika tidak ingin mengubah): ")
// 			fmt.Scanln(&pengarangBaru)
// 			fmt.Print("Masukkan tahun terbit baru (tekan 0 jika tidak ingin mengubah): ")
// 			fmt.Scanln(&tahunBaru)

// 			// Update informasi buku jika input diberikan
// 			if judulBaru != "" {
// 				perpustakaan[index-1].Judul = judulBaru
// 			}
// 			if pengarangBaru != "" {
// 				perpustakaan[index-1].Pengarang = pengarangBaru
// 			}
// 			if tahunBaru != 0 {
// 				perpustakaan[index-1].Tahun = tahunBaru
// 			}

// 			fmt.Println("Informasi buku berhasil diupdate.")
// 		} else {
// 			fmt.Println("Nomor buku tidak valid.")
// 		}
// 	}
// 	return perpustakaan
// }

// // Fungsi untuk menghapus buku
// func hapusBuku(perpustakaan []Buku) []Buku {
// 	if len(perpustakaan) == 0 {
// 		fmt.Println("Tidak ada buku untuk dihapus.")
// 	} else {
// 		// Tampilkan daftar buku
// 		tampilkanBuku(perpustakaan)

// 		// Memilih buku yang akan dihapus
// 		var index int
// 		fmt.Print("Masukkan nomor buku yang akan dihapus: ")
// 		fmt.Scanln(&index)

// 		// Validasi nomor buku
// 		if index > 0 && index <= len(perpustakaan) {
// 			perpustakaan = append(perpustakaan[:index-1], perpustakaan[index:]...)
// 			fmt.Println("Buku berhasil dihapus.")
// 		} else {
// 			fmt.Println("Nomor buku tidak valid.")
// 		}
// 	}
// 	return perpustakaan
// }