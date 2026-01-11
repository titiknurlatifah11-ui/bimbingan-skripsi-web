// //1.ARRAY
// // package main

// // import "fmt"

// // func main() {
// // 	var numbers [5]int
// // 	numbers[0] = 10
// //     numbers[1] = 20
// //     numbers[2] = 30
// // // Mengubah elemen
// //     numbers[2] = 50
// // // Mengiterasi array
// //     for i, num := range numbers {
// //         fmt.Printf("Element %d: %d\n", i, num)
// //     }
// // }

// //modifikasi
// // package main

// // import "fmt"

// // func main() {
// // 	var numbers [5]int

// // 	// Mengisi array dengan input dari pengguna
// // 	for i := 0; i < len(numbers); i++ {
// // 		fmt.Printf("Masukkan elemen ke-%d: ", i+1)
// // 		fmt.Scanln(&numbers[i])
// // 	}

// // 	// Mengubah elemen tertentu (misalnya elemen ke-3)
// // 	numbers[2] = 50

// // 	// Mengiterasi array dan menampilkan elemen
// // 	var total int
// // 	for i, num := range numbers {
// // 		fmt.Printf("Element %d: %d\n", i, num)
// // 		total += num
// // 	}

// // 	// Menampilkan total
// // 	fmt.Printf("Total: %d\n", total)
// // }

// //STUDIKASUS array
// // package main

// // import (
// // 	"fmt"
// // )

// // func main() {
// // 	// Menentukan jumlah nilai ujian yang akan diinput
// // 	var numScores int
// // 	fmt.Print("Berapa banyak nilai ujian yang akan diinput? ")
// // 	fmt.Scanln(&numScores)

// // 	// Membuat array untuk menyimpan nilai-nilai ujian
// // 	scores := make([]float64, numScores)

// // 	// Menginput nilai-nilai ujian dari pengguna
// // 	for i := 0; i < numScores; i++ {
// // 		fmt.Printf("Masukkan nilai ujian ke-%d: ", i+1)
// // 		fmt.Scanln(&scores[i])
// // 	}

// // 	// Menghitung total nilai dan rata-rata
// // 	var total float64
// // 	for _, score := range scores {
// // 		total += score
// // 	}

// // 	average := total / float64(numScores)

// // 	// Menampilkan rata-rata nilai
// // 	fmt.Printf("Rata-rata nilai ujian: %.2f\n", average)
// // }

// //modifikasi STUDIKASUS array
// // package main

// // import (
// // 	"fmt"
// // )

// // func main() {
// // 	// Menentukan jumlah nilai ujian yang akan diinput
// // 	var numScores int
// // 	fmt.Print("Berapa banyak nilai ujian yang akan diinput? ")
// // 	fmt.Scanln(&numScores)

// // 	// Membuat array untuk menyimpan nilai-nilai ujian
// // 	scores := make([]float64, numScores)

// // 	// Menginput nilai-nilai ujian dari pengguna dengan validasi
// // 	for i := 0; i < numScores; i++ {
// // 		for {
// // 			fmt.Printf("Masukkan nilai ujian ke-%d: ", i+1)
// // 			_, err := fmt.Scanln(&scores[i])
// // 			if err == nil && scores[i] >= 0 && scores[i] <= 100 {
// // 				break // Input valid, keluar dari loop
// // 			} else {
// // 				fmt.Println("Nilai tidak valid. Harap masukkan angka antara 0 hingga 100.")
// // 			}
// // 		}
// // 	}

// // 	// Menghitung total nilai dan rata-rata
// // 	var total float64
// // 	for _, score := range scores {
// // 		total += score
// // 	}

// // 	average := total / float64(numScores)

// // 	// Menampilkan total dan rata-rata nilai
// // 	fmt.Printf("Total nilai ujian: %.2f\n", total)
// // 	fmt.Printf("Rata-rata nilai ujian: %.2f\n", average)
// // }


// //2.SLICE
// // package main
// // import "fmt"

// // func main() {
// // 	var numbers []int
// //     numbers = append(numbers, 1, 2, 3, 4, 5)
// //     fmt.Println("Slice sebelum perubahan:", numbers)
// // // Menghapus elemen ketiga
// //     numbers = append(numbers[:2], numbers[3:]...)
// //     fmt.Println("Slice setelah perubahan:", numbers)
// // }

// //modifikasi
// // package main

// // import "fmt"

// // func main() {
// // 	var numbers []int

// // 	// Mengisi slice dengan input dari pengguna
// // 	fmt.Println("Masukkan 5 angka:")
// // 	for i := 0; i < 5; i++ {
// // 		var num int
// // 		fmt.Printf("Angka ke-%d: ", i+1)
// // 		fmt.Scanln(&num)
// // 		numbers = append(numbers, num)
// // 	}

// // 	fmt.Println("Slice sebelum perubahan:", numbers)

// // 	// Meminta pengguna untuk menghapus elemen berdasarkan indeks
// // 	var index int
// // 	fmt.Print("Masukkan indeks elemen yang ingin dihapus (0-4): ")
// // 	fmt.Scanln(&index)

// // 	// Validasi indeks dan menghapus elemen
// // 	if index >= 0 && index < len(numbers) {
// // 		numbers = append(numbers[:index], numbers[index+1:]...)
// // 		fmt.Println("Slice setelah perubahan:", numbers)
// // 	} else {
// // 		fmt.Println("Indeks tidak valid.")
// // 	}
// // }

// //STUDIKASUS slice
// // package main

// // import (
// // 	"fmt"
// // )

// // // Struct untuk menyimpan informasi barang
// // type Item struct {
// // 	Name  string
// // 	Price float64
// // }

// // func main() {
// // 	// Slice untuk menyimpan daftar barang yang akan dibeli
// // 	var cart []Item

// // 	// Fungsi untuk menampilkan menu pilihan
// // 	for {
// // 		fmt.Println("\nMenu:")
// // 		fmt.Println("1. Tambah Barang")
// // 		fmt.Println("2. Hapus Barang")
// // 		fmt.Println("3. Tampilkan Daftar Barang")
// // 		fmt.Println("4. Keluar")
// // 		fmt.Print("Pilih opsi: ")

// // 		var choice int
// // 		fmt.Scanln(&choice)

// // 		switch choice {
// // 		case 1:
// // 			// Menambahkan barang ke keranjang
// // 			var name string
// // 			var price float64
// // 			fmt.Print("Masukkan nama barang: ")
// // 			fmt.Scanln(&name)
// // 			fmt.Print("Masukkan harga barang: ")
// // 			fmt.Scanln(&price)
// // 			cart = append(cart, Item{Name: name, Price: price})
// // 			fmt.Println("Barang berhasil ditambahkan.")

// // 		case 2:
// // 			// Menghapus barang dari keranjang
// // 			if len(cart) == 0 {
// // 				fmt.Println("Keranjang kosong, tidak ada barang untuk dihapus.")
// // 			} else {
// // 				fmt.Println("Daftar Barang:")
// // 				for i, item := range cart {
// // 					fmt.Printf("%d. %s - Rp %.2f\n", i+1, item.Name, item.Price)
// // 				}
// // 				var index int
// // 				fmt.Print("Masukkan nomor barang yang akan dihapus: ")
// // 				fmt.Scanln(&index)

// // 				// Validasi nomor barang yang dimasukkan
// // 				if index > 0 && index <= len(cart) {
// // 					cart = append(cart[:index-1], cart[index:]...)
// // 					fmt.Println("Barang berhasil dihapus.")
// // 				} else {
// // 					fmt.Println("Nomor barang tidak valid.")
// // 				}
// // 			}

// // 		case 3:
// // 			// Menampilkan daftar barang di keranjang
// // 			if len(cart) == 0 {
// // 				fmt.Println("Keranjang kosong.")
// // 			} else {
// // 				fmt.Println("Daftar Barang di Keranjang:")
// // 				for i, item := range cart {
// // 					fmt.Printf("%d. %s - Rp %.2f\n", i+1, item.Name, item.Price)
// // 				}
// // 			}

// // 		case 4:
// // 			// Keluar dari program
// // 			fmt.Println("Terima kasih telah menggunakan program ini!")
// // 			return

// // 		default:
// // 			fmt.Println("Pilihan tidak valid, silakan coba lagi.")
// // 		}
// // 	}
// // }

// //modifikasi STUDIKASUS slice
// // package main

// // import (
// // 	"fmt"
// // )

// // // Struct untuk menyimpan informasi barang
// // type Item struct {
// // 	Name  string
// // 	Price float64
// // }

// // func main() {
// // 	// Slice untuk menyimpan daftar barang yang akan dibeli
// // 	var cart []Item

// // 	// Fungsi untuk menampilkan menu pilihan
// // 	for {
// // 		fmt.Println("\nMenu:")
// // 		fmt.Println("1. Tambah Barang")
// // 		fmt.Println("2. Hapus Barang")
// // 		fmt.Println("3. Tampilkan Daftar Barang")
// // 		fmt.Println("4. Keluar")
// // 		fmt.Print("Pilih opsi: ")

// // 		var choice int
// // 		fmt.Scanln(&choice)

// // 		switch choice {
// // 		case 1:
// // 			// Menambahkan barang ke keranjang
// // 			var name string
// // 			var price float64
// // 			fmt.Print("Masukkan nama barang: ")
// // 			fmt.Scanln(&name)
// // 			fmt.Print("Masukkan harga barang: ")
// // 			fmt.Scanln(&price)
// // 			cart = append(cart, Item{Name: name, Price: price})
// // 			fmt.Println("Barang berhasil ditambahkan.")

// // 		case 2:
// // 			// Menghapus barang dari keranjang
// // 			if len(cart) == 0 {
// // 				fmt.Println("Keranjang kosong, tidak ada barang untuk dihapus.")
// // 			} else {
// // 				fmt.Println("Daftar Barang:")
// // 				for i, item := range cart {
// // 					fmt.Printf("%d. %s - Rp %.2f\n", i+1, item.Name, item.Price)
// // 				}
// // 				var index int
// // 				fmt.Print("Masukkan nomor barang yang akan dihapus: ")
// // 				fmt.Scanln(&index)

// // 				// Validasi nomor barang yang dimasukkan
// // 				if index > 0 && index <= len(cart) {
// // 					cart = append(cart[:index-1], cart[index:]...)
// // 					fmt.Println("Barang berhasil dihapus.")
// // 				} else {
// // 					fmt.Println("Nomor barang tidak valid.")
// // 				}
// // 			}

// // 		case 3:
// // 			// Menampilkan daftar barang di keranjang dan total harga
// // 			if len(cart) == 0 {
// // 				fmt.Println("Keranjang kosong.")
// // 			} else {
// // 				fmt.Println("Daftar Barang di Keranjang:")
// // 				var total float64
// // 				for i, item := range cart {
// // 					fmt.Printf("%d. %s - Rp %.2f\n", i+1, item.Name, item.Price)
// // 					total += item.Price
// // 				}
// // 				fmt.Printf("Total harga barang: Rp %.2f\n", total)
// // 			}

// // 		case 4:
// // 			// Konfirmasi keluar dari program
// // 			var confirm string
// // 			fmt.Print("Apakah Anda yakin ingin keluar? (y/n): ")
// // 			fmt.Scanln(&confirm)
// // 			if confirm == "y" || confirm == "Y" {
// // 				fmt.Println("Terima kasih telah menggunakan program ini!")
// // 				return
// // 			}

// // 		default:
// // 			fmt.Println("Pilihan tidak valid, silakan coba lagi.")
// // 		}
// // 	}
// // }


// //3.MAP
// // package main

// // import "fmt"

// // func main() {
// // 	ages := map[string]int{
// // 		"John": 25,
// // 		"Jane": 30,
// // 	}

// // 	// Menambahkan elemen
// // 	ages["Doe"] = 22

// // 	// Menghapus elemen
// // 	delete(ages, "John")

// // 	// Memeriksa keberadaan kunci
// // 	age, exists := ages["Jane"]
// // 	if exists {
// // 		fmt.Println("Jane's age:", age)
// // 	} else {
// // 		fmt.Println("Jane tidak ditemukan.")
// // 	}
// // }

// //modifikasi
// // package main

// // import "fmt"

// // func main() {
// // 	ages := map[string]int{
// // 		"John": 25,
// // 		"Jane": 30,
// // 	}

// // 	// Menampilkan usia awal
// // 	fmt.Println("Usia awal:")
// // 	for name, age := range ages {
// // 		fmt.Printf("%s: %d\n", name, age)
// // 	}

// // 	// Menambahkan elemen
// // 	var newName string
// // 	var newAge int
// // 	fmt.Print("Masukkan nama untuk menambahkan usia: ")
// // 	fmt.Scanln(&newName)
// // 	fmt.Print("Masukkan usia: ")
// // 	fmt.Scanln(&newAge)
// // 	ages[newName] = newAge

// // 	// Menghapus elemen
// // 	var deleteName string
// // 	fmt.Print("Masukkan nama untuk dihapus: ")
// // 	fmt.Scanln(&deleteName)
// // 	delete(ages, deleteName)

// // 	// Menampilkan usia setelah perubahan
// // 	fmt.Println("\nUsia setelah perubahan:")
// // 	for name, age := range ages {
// // 		fmt.Printf("%s: %d\n", name, age)
// // 	}
// // }

// //STUDIKASUS map
// // package main

// // import (
// // 	"fmt"
// // )

// // func main() {
// // 	// Map untuk menyimpan data mahasiswa (nama sebagai kunci, nilai sebagai nilai)
// // 	mahasiswa := make(map[string]float64)

// // 	for {
// // 		// Menampilkan menu pilihan
// // 		fmt.Println("\nMenu:")
// // 		fmt.Println("1. Tambah Mahasiswa")
// // 		fmt.Println("2. Hapus Mahasiswa")
// // 		fmt.Println("3. Update Nilai Mahasiswa")
// // 		fmt.Println("4. Tampilkan Semua Mahasiswa")
// // 		fmt.Println("5. Keluar")
// // 		fmt.Print("Pilih opsi: ")

// // 		var choice int
// // 		fmt.Scanln(&choice)

// // 		switch choice {
// // 		case 1:
// // 			// Menambahkan mahasiswa baru
// // 			var nama string
// // 			var nilai float64
// // 			fmt.Print("Masukkan nama mahasiswa: ")
// // 			fmt.Scanln(&nama)
// // 			fmt.Print("Masukkan nilai mahasiswa: ")
// // 			fmt.Scanln(&nilai)
// // 			mahasiswa[nama] = nilai
// // 			fmt.Println("Mahasiswa berhasil ditambahkan.")

// // 		case 2:
// // 			// Menghapus data mahasiswa
// // 			var nama string
// // 			fmt.Print("Masukkan nama mahasiswa yang akan dihapus: ")
// // 			fmt.Scanln(&nama)
// // 			if _, exists := mahasiswa[nama]; exists {
// // 				delete(mahasiswa, nama)
// // 				fmt.Println("Data mahasiswa berhasil dihapus.")
// // 			} else {
// // 				fmt.Println("Mahasiswa tidak ditemukan.")
// // 			}

// // 		case 3:
// // 			// Mengupdate nilai mahasiswa
// // 			var nama string
// // 			var nilaiBaru float64
// // 			fmt.Print("Masukkan nama mahasiswa yang akan diupdate nilainya: ")
// // 			fmt.Scanln(&nama)
// // 			if _, exists := mahasiswa[nama]; exists {
// // 				fmt.Print("Masukkan nilai baru: ")
// // 				fmt.Scanln(&nilaiBaru)
// // 				mahasiswa[nama] = nilaiBaru
// // 				fmt.Println("Nilai mahasiswa berhasil diupdate.")
// // 			} else {
// // 				fmt.Println("Mahasiswa tidak ditemukan.")
// // 			}

// // 		case 4:
// // 			// Menampilkan semua data mahasiswa
// // 			if len(mahasiswa) == 0 {
// // 				fmt.Println("Tidak ada data mahasiswa.")
// // 			} else {
// // 				fmt.Println("Daftar Mahasiswa:")
// // 				for nama, nilai := range mahasiswa {
// // 					fmt.Printf("Nama: %s, Nilai: %.2f\n", nama, nilai)
// // 				}
// // 			}

// // 		case 5:
// // 			// Keluar dari program
// // 			fmt.Println("Terima kasih telah menggunakan program ini!")
// // 			return

// // 		default:
// // 			fmt.Println("Pilihan tidak valid, silakan coba lagi.")
// // 		}
// // 	}
// // }

// //modifikasi STUDIKASUS map
// // package main

// // import (
// // 	"fmt"
// // )

// // func main() {
// // 	// Map untuk menyimpan data mahasiswa (nama sebagai kunci, nilai sebagai nilai)
// // 	mahasiswa := make(map[string]float64)

// // 	for {
// // 		// Menampilkan menu pilihan
// // 		fmt.Println("\nMenu:")
// // 		fmt.Println("1. Tambah Mahasiswa")
// // 		fmt.Println("2. Hapus Mahasiswa")
// // 		fmt.Println("3. Update Nilai Mahasiswa")
// // 		fmt.Println("4. Tampilkan Semua Mahasiswa")
// // 		fmt.Println("5. Keluar")
// // 		fmt.Print("Pilih opsi: ")

// // 		var choice int
// // 		fmt.Scanln(&choice)

// // 		switch choice {
// // 		case 1:
// // 			// Menambahkan mahasiswa baru
// // 			var nama string
// // 			var nilai float64
// // 			fmt.Print("Masukkan nama mahasiswa: ")
// // 			fmt.Scanln(&nama)
// // 			fmt.Print("Masukkan nilai mahasiswa: ")
// // 			fmt.Scanln(&nilai)
// // 			mahasiswa[nama] = nilai
// // 			fmt.Println("Mahasiswa berhasil ditambahkan.")

// // 		case 2:
// // 			// Menghapus data mahasiswa dengan konfirmasi
// // 			var nama string
// // 			fmt.Print("Masukkan nama mahasiswa yang akan dihapus: ")
// // 			fmt.Scanln(&nama)
// // 			if _, exists := mahasiswa[nama]; exists {
// // 				var confirm string
// // 				fmt.Print("Apakah Anda yakin ingin menghapus data mahasiswa ini? (y/n): ")
// // 				fmt.Scanln(&confirm)
// // 				if confirm == "y" || confirm == "Y" {
// // 					delete(mahasiswa, nama)
// // 					fmt.Println("Data mahasiswa berhasil dihapus.")
// // 				} else {
// // 					fmt.Println("Penghapusan dibatalkan.")
// // 				}
// // 			} else {
// // 				fmt.Println("Mahasiswa tidak ditemukan.")
// // 			}

// // 		case 3:
// // 			// Mengupdate nilai mahasiswa
// // 			var nama string
// // 			var nilaiBaru float64
// // 			fmt.Print("Masukkan nama mahasiswa yang akan diupdate nilainya: ")
// // 			fmt.Scanln(&nama)
// // 			if _, exists := mahasiswa[nama]; exists {
// // 				fmt.Print("Masukkan nilai baru: ")
// // 				fmt.Scanln(&nilaiBaru)
// // 				mahasiswa[nama] = nilaiBaru
// // 				fmt.Println("Nilai mahasiswa berhasil diupdate.")
// // 			} else {
// // 				fmt.Println("Mahasiswa tidak ditemukan.")
// // 			}

// // 		case 4:
// // 			// Menampilkan semua data mahasiswa dan jumlah total mahasiswa
// // 			if len(mahasiswa) == 0 {
// // 				fmt.Println("Tidak ada data mahasiswa.")
// // 			} else {
// // 				fmt.Println("Daftar Mahasiswa:")
// // 				for nama, nilai := range mahasiswa {
// // 					fmt.Printf("Nama: %s, Nilai: %.2f\n", nama, nilai)
// // 				}
// // 				fmt.Printf("Total mahasiswa: %d\n", len(mahasiswa))
// // 			}

// // 		case 5:
// // 			// Keluar dari program
// // 			fmt.Println("Terima kasih telah menggunakan program ini!")
// // 			return

// // 		default:
// // 			fmt.Println("Pilihan tidak valid, silakan coba lagi.")
// // 		}
// // 	}
// // }


// //4.STRUCT
// // package main

// // import "fmt"

// // // Struct untuk menyimpan informasi orang
// // type Person struct {
// // 	Name string
// // 	Age  int
// // }

// // // Fungsi untuk memperbarui usia orang menggunakan pointer
// // func updateAge(p *Person, newAge int) {
// // 	p.Age = newAge
// // }

// // func main() {
// // 	person := Person{Name: "John", Age: 25}
// // 	fmt.Println("Sebelum:", person)
// // 	updateAge(&person, 30)
// // 	fmt.Println("Setelah:", person)
// // }

// //modifikasi
// // package main

// // import "fmt"

// // // Struct untuk menyimpan informasi orang
// // type Person struct {
// // 	Name string
// // 	Age  int
// // }

// // // Fungsi untuk memperbarui usia dan nama orang menggunakan pointer
// // func updatePerson(p *Person, newName string, newAge int) {
// // 	p.Name = newName
// // 	p.Age = newAge
// // }

// // func main() {
// // 	// Membuat objek Person dengan data awal
// // 	person := Person{Name: "John", Age: 25}
// // 	fmt.Println("Sebelum:", person)

// // 	// Input dari pengguna untuk nama dan usia baru
// // 	var newName string
// // 	var newAge int
// // 	fmt.Print("Masukkan nama baru: ")
// // 	fmt.Scanln(&newName)
// // 	fmt.Print("Masukkan usia baru: ")
// // 	fmt.Scanln(&newAge)

// // 	// Memperbarui informasi orang
// // 	updatePerson(&person, newName, newAge)
// // 	fmt.Println("Setelah:", person)
// // }

// //STUDIKASUS struct
// // package main

// // import (
// // 	"fmt"
// // )

// // // Struct untuk menyimpan informasi buku
// // type Buku struct {
// // 	Judul     string
// // 	Pengarang string
// // 	Tahun     int
// // }

// // // Fungsi utama
// // func main() {
// // 	// Slice untuk menyimpan daftar buku
// // 	var perpustakaan []Buku
// // 	for {
// // 		// Menampilkan menu pilihan
// // 		fmt.Println("\nMenu:")
// // 		fmt.Println("1. Tambah Buku")
// // 		fmt.Println("2. Tampilkan Semua Buku")
// // 		fmt.Println("3. Update Informasi Buku")
// // 		fmt.Println("4. Keluar")
// // 		fmt.Print("Pilih opsi: ")

// // 		var pilihan int
// // 		fmt.Scanln(&pilihan)

// // 		switch pilihan {
// // 		case 1:
// // 			// Panggil fungsi untuk menambah buku
// // 			perpustakaan = tambahBuku(perpustakaan)
// // 		case 2:
// // 			// Panggil fungsi untuk menampilkan semua buku
// // 			tampilkanBuku(perpustakaan)
// // 		case 3:
// // 			// Panggil fungsi untuk mengupdate informasi buku
// // 			perpustakaan = updateBuku(perpustakaan)
// // 		case 4:
// // 			// Keluar dari program
// // 			fmt.Println("Terima kasih telah menggunakan program ini!")
// // 			return
// // 		default:
// // 			fmt.Println("Pilihan tidak valid, silakan coba lagi.")
// // 		}
// // 	}
// // }

// // // Fungsi untuk menambah buku baru
// // func tambahBuku(perpustakaan []Buku) []Buku {
// // 	var judul, pengarang string
// // 	var tahun int

// // 	fmt.Print("Masukkan judul buku: ")
// // 	fmt.Scanln(&judul)
// // 	fmt.Print("Masukkan nama pengarang: ")
// // 	fmt.Scanln(&pengarang)
// // 	fmt.Print("Masukkan tahun terbit: ")
// // 	fmt.Scanln(&tahun)

// // 	// Tambahkan buku baru ke slice perpustakaan
// // 	perpustakaan = append(perpustakaan, Buku{Judul: judul, Pengarang: pengarang, Tahun: tahun})
// // 	fmt.Println("Buku berhasil ditambahkan.")
// // 	return perpustakaan
// // }

// // // Fungsi untuk menampilkan semua buku
// // func tampilkanBuku(perpustakaan []Buku) {
// // 	if len(perpustakaan) == 0 {
// // 		fmt.Println("Tidak ada buku di perpustakaan.")
// // 	} else {
// // 		fmt.Println("Daftar Buku:")
// // 		for i, buku := range perpustakaan {
// // 			fmt.Printf("%d. Judul: %s, Pengarang: %s, Tahun Terbit: %d\n", i+1, buku.Judul, buku.Pengarang, buku.Tahun)
// // 		}
// // 	}
// // }

// // // Fungsi untuk mengupdate informasi buku
// // func updateBuku(perpustakaan []Buku) []Buku {
// // 	if len(perpustakaan) == 0 {
// // 		fmt.Println("Tidak ada buku untuk diupdate.")
// // 	} else {
// // 		// Tampilkan daftar buku
// // 		tampilkanBuku(perpustakaan)

// // 		// Memilih buku yang akan diupdate
// // 		var index int
// // 		fmt.Print("Masukkan nomor buku yang akan diupdate: ")
// // 		fmt.Scanln(&index)

// // 		// Validasi nomor buku
// // 		if index > 0 && index <= len(perpustakaan) {
// // 			var judulBaru, pengarangBaru string
// // 			var tahunBaru int

// // 			fmt.Print("Masukkan judul baru (tekan Enter jika tidak ingin mengubah): ")
// // 			fmt.Scanln(&judulBaru)
// // 			fmt.Print("Masukkan nama pengarang baru (tekan Enter jika tidak ingin mengubah): ")
// // 			fmt.Scanln(&pengarangBaru)
// // 			fmt.Print("Masukkan tahun terbit baru (tekan 0 jika tidak ingin mengubah): ")
// // 			fmt.Scanln(&tahunBaru)

// // 			// Update informasi buku jika input diberikan
// // 			if judulBaru != "" {
// // 				perpustakaan[index-1].Judul = judulBaru
// // 			}
// // 			if pengarangBaru != "" {
// // 				perpustakaan[index-1].Pengarang = pengarangBaru
// // 			}
// // 			if tahunBaru != 0 {
// // 				perpustakaan[index-1].Tahun = tahunBaru
// // 			}

// // 			fmt.Println("Informasi buku berhasil diupdate.")
// // 		} else {
// // 			fmt.Println("Nomor buku tidak valid.")
// // 		}
// // 	}
// // 	return perpustakaan
// // }

// //modifikasi STUDIKASUS struct
// // package main

// // import (
// // 	"fmt"
// // )

// // // Struct untuk menyimpan informasi buku
// // type Buku struct {
// // 	Judul     string
// // 	Pengarang string
// // 	Tahun     int
// // }

// // // Fungsi utama
// // func main() {
// // 	// Slice untuk menyimpan daftar buku
// // 	var perpustakaan []Buku
// // 	for {
// // 		// Menampilkan menu pilihan
// // 		fmt.Println("\nMenu:")
// // 		fmt.Println("1. Tambah Buku")
// // 		fmt.Println("2. Tampilkan Semua Buku")
// // 		fmt.Println("3. Update Informasi Buku")
// // 		fmt.Println("4. Hapus Buku")
// // 		fmt.Println("5. Keluar")
// // 		fmt.Print("Pilih opsi: ")

// // 		var pilihan int
// // 		fmt.Scanln(&pilihan)

// // 		switch pilihan {
// // 		case 1:
// // 			// Panggil fungsi untuk menambah buku
// // 			perpustakaan = tambahBuku(perpustakaan)
// // 		case 2:
// // 			// Panggil fungsi untuk menampilkan semua buku
// // 			tampilkanBuku(perpustakaan)
// // 		case 3:
// // 			// Panggil fungsi untuk mengupdate informasi buku
// // 			perpustakaan = updateBuku(perpustakaan)
// // 		case 4:
// // 			// Panggil fungsi untuk menghapus buku
// // 			perpustakaan = hapusBuku(perpustakaan)
// // 		case 5:
// // 			// Keluar dari program
// // 			fmt.Println("Terima kasih telah menggunakan program ini!")
// // 			return
// // 		default:
// // 			fmt.Println("Pilihan tidak valid, silakan coba lagi.")
// // 		}
// // 	}
// // }

// // // Fungsi untuk menambah buku baru
// // func tambahBuku(perpustakaan []Buku) []Buku {
// // 	var judul, pengarang string
// // 	var tahun int

// // 	fmt.Print("Masukkan judul buku: ")
// // 	fmt.Scanln(&judul)
// // 	fmt.Print("Masukkan nama pengarang: ")
// // 	fmt.Scanln(&pengarang)
// // 	fmt.Print("Masukkan tahun terbit: ")
// // 	fmt.Scanln(&tahun)

// // 	// Tambahkan buku baru ke slice perpustakaan
// // 	perpustakaan = append(perpustakaan, Buku{Judul: judul, Pengarang: pengarang, Tahun: tahun})
// // 	fmt.Println("Buku berhasil ditambahkan.")
// // 	return perpustakaan
// // }

// // // Fungsi untuk menampilkan semua buku
// // func tampilkanBuku(perpustakaan []Buku) {
// // 	if len(perpustakaan) == 0 {
// // 		fmt.Println("Tidak ada buku di perpustakaan.")
// // 	} else {
// // 		fmt.Println("Daftar Buku:")
// // 		for i, buku := range perpustakaan {
// // 			fmt.Printf("%d. Judul: %s, Pengarang: %s, Tahun Terbit: %d\n", i+1, buku.Judul, buku.Pengarang, buku.Tahun)
// // 		}
// // 	}
// // }

// // // Fungsi untuk mengupdate informasi buku
// // func updateBuku(perpustakaan []Buku) []Buku {
// // 	if len(perpustakaan) == 0 {
// // 		fmt.Println("Tidak ada buku untuk diupdate.")
// // 	} else {
// // 		// Tampilkan daftar buku
// // 		tampilkanBuku(perpustakaan)

// // 		// Memilih buku yang akan diupdate
// // 		var index int
// // 		fmt.Print("Masukkan nomor buku yang akan diupdate: ")
// // 		fmt.Scanln(&index)

// // 		// Validasi nomor buku
// // 		if index > 0 && index <= len(perpustakaan) {
// // 			var judulBaru, pengarangBaru string
// // 			var tahunBaru int

// // 			fmt.Print("Masukkan judul baru (tekan Enter jika tidak ingin mengubah): ")
// // 			fmt.Scanln(&judulBaru)
// // 			fmt.Print("Masukkan nama pengarang baru (tekan Enter jika tidak ingin mengubah): ")
// // 			fmt.Scanln(&pengarangBaru)
// // 			fmt.Print("Masukkan tahun terbit baru (tekan 0 jika tidak ingin mengubah): ")
// // 			fmt.Scanln(&tahunBaru)

// // 			// Update informasi buku jika input diberikan
// // 			if judulBaru != "" {
// // 				perpustakaan[index-1].Judul = judulBaru
// // 			}
// // 			if pengarangBaru != "" {
// // 				perpustakaan[index-1].Pengarang = pengarangBaru
// // 			}
// // 			if tahunBaru != 0 {
// // 				perpustakaan[index-1].Tahun = tahunBaru
// // 			}

// // 			fmt.Println("Informasi buku berhasil diupdate.")
// // 		} else {
// // 			fmt.Println("Nomor buku tidak valid.")
// // 		}
// // 	}
// // 	return perpustakaan
// // }

// // // Fungsi untuk menghapus buku
// // func hapusBuku(perpustakaan []Buku) []Buku {
// // 	if len(perpustakaan) == 0 {
// // 		fmt.Println("Tidak ada buku untuk dihapus.")
// // 	} else {
// // 		// Tampilkan daftar buku
// // 		tampilkanBuku(perpustakaan)

// // 		// Memilih buku yang akan dihapus
// // 		var index int
// // 		fmt.Print("Masukkan nomor buku yang akan dihapus: ")
// // 		fmt.Scanln(&index)

// // 		// Validasi nomor buku
// // 		if index > 0 && index <= len(perpustakaan) {
// // 			perpustakaan = append(perpustakaan[:index-1], perpustakaan[index:]...)
// // 			fmt.Println("Buku berhasil dihapus.")
// // 		} else {
// // 			fmt.Println("Nomor buku tidak valid.")
// // 		}
// // 	}
// // 	return perpustakaan
// // }

// //5. POINTER
// // package main

// // import "fmt"

// // // Fungsi untuk mengubah nilai menggunakan pointer
// // func changeValue(x *int) {
// // 	*x = 50
// // }

// // func main() {
// // 	a := 10
// // 	fmt.Println("Nilai a sebelum:", a)
// // 	changeValue(&a)
// // 	fmt.Println("Nilai a setelah:", a)
// // }

// //modifikasi
// // package main

// // import "fmt"

// // // Fungsi untuk mengubah nilai menggunakan pointer
// // func changeValue(x *int, newValue int) {
// // 	*x = newValue
// // }

// // func main() {
// // 	var a int
// // 	fmt.Print("Masukkan nilai awal a: ")
// // 	fmt.Scanln(&a)
// // 	fmt.Println("Nilai a sebelum:", a)

// // 	var newValue int
// // 	fmt.Print("Masukkan nilai baru untuk a: ")
// // 	fmt.Scanln(&newValue)

// // 	changeValue(&a, newValue)
// // 	fmt.Println("Nilai a setelah:", a)
// // }

// //STUDIKASUS pointer
// // package main

// // import (
// // 	"fmt"
// // )

// // // Struct untuk menyimpan informasi karyawan
// // type Karyawan struct {
// // 	Nama   string
// // 	Gaji   float64
// // 	Posisi string
// // }

// // // Fungsi untuk mengubah gaji karyawan menggunakan pointer
// // func ubahGaji(k *Karyawan, gajiBaru float64) {
// // 	k.Gaji = gajiBaru
// // }

// // func main() {
// // 	// Membuat data karyawan
// // 	karyawan1 := Karyawan{
// // 		Nama:   "Budi",
// // 		Gaji:   5000000,
// // 		Posisi: "Software Engineer",
// // 	}

// // 	// Menampilkan informasi karyawan sebelum perubahan
// // 	fmt.Println("Sebelum perubahan:")
// // 	fmt.Printf("Nama: %s, Gaji: %.2f, Posisi: %s\n", karyawan1.Nama, karyawan1.Gaji, karyawan1.Posisi)

// // 	// Mengubah gaji karyawan menggunakan fungsi ubahGaji
// // 	ubahGaji(&karyawan1, 7000000)

// // 	// Menampilkan informasi karyawan setelah perubahan
// // 	fmt.Println("\nSetelah perubahan:")
// // 	fmt.Printf("Nama: %s, Gaji: %.2f, Posisi: %s\n", karyawan1.Nama, karyawan1.Gaji, karyawan1.Posisi)
// // }

// //modifikasi STUDIKASUS pointer
// // package main

// // import (
// // 	"fmt"
// // )

// // // Struct untuk menyimpan informasi karyawan
// // type Karyawan struct {
// // 	Nama   string
// // 	Gaji   float64
// // 	Posisi string
// // }

// // // Fungsi untuk mengubah gaji karyawan menggunakan pointer
// // func ubahGaji(k *Karyawan, gajiBaru float64) {
// // 	k.Gaji = gajiBaru
// // }

// // // Fungsi untuk mengubah posisi karyawan menggunakan pointer
// // func ubahPosisi(k *Karyawan, posisiBaru string) {
// // 	k.Posisi = posisiBaru
// // }

// // func main() {
// // 	// Membuat data karyawan dengan input dari pengguna
// // 	var nama string
// // 	var gaji float64
// // 	var posisi string

// // 	fmt.Print("Masukkan nama karyawan: ")
// // 	fmt.Scanln(&nama)
// // 	fmt.Print("Masukkan gaji karyawan: ")
// // 	fmt.Scanln(&gaji)
// // 	fmt.Print("Masukkan posisi karyawan: ")
// // 	fmt.Scanln(&posisi)

// // 	karyawan1 := Karyawan{
// // 		Nama:   nama,
// // 		Gaji:   gaji,
// // 		Posisi: posisi,
// // 	}

// // 	// Menampilkan informasi karyawan sebelum perubahan
// // 	fmt.Println("\nSebelum perubahan:")
// // 	fmt.Printf("Nama: %s, Gaji: %.2f, Posisi: %s\n", karyawan1.Nama, karyawan1.Gaji, karyawan1.Posisi)

// // 	// Mengubah gaji karyawan
// // 	var gajiBaru float64
// // 	fmt.Print("Masukkan gaji baru: ")
// // 	fmt.Scanln(&gajiBaru)
// // 	ubahGaji(&karyawan1, gajiBaru)

// // 	// Mengubah posisi karyawan
// // 	var posisiBaru string
// // 	fmt.Print("Masukkan posisi baru: ")
// // 	fmt.Scanln(&posisiBaru)
// // 	ubahPosisi(&karyawan1, posisiBaru)

// // 	// Menampilkan informasi karyawan setelah perubahan
// // 	fmt.Println("\nSetelah perubahan:")
// // 	fmt.Printf("Nama: %s, Gaji: %.2f, Posisi: %s\n", karyawan1.Nama, karyawan1.Gaji, karyawan1.Posisi)
// // }

// //Latihan1
// // package main

// // import (
// // 	"fmt"
// // )

// // // Struct untuk mendefinisikan Karyawan
// // type Karyawan struct {
// // 	ID   int
// // 	Nama string
// // 	Usia int
// // 	Gaji float64
// // }

// // // Deklarasi global slice untuk menyimpan data karyawan
// // var daftarKaryawan []Karyawan
// // var karyawanMap = make(map[string]*Karyawan) // Map untuk pencarian berdasarkan nama

// // // Fungsi untuk menambah karyawan
// // func tambahKaryawan(id int, nama string, usia int, gaji float64) {
// // 	karyawanBaru := Karyawan{ID: id, Nama: nama, Usia: usia, Gaji: gaji}
// // 	daftarKaryawan = append(daftarKaryawan, karyawanBaru)
// // 	karyawanMap[nama] = &karyawanBaru // Menyimpan karyawan ke map berdasarkan nama
// // 	fmt.Println("Karyawan berhasil ditambahkan!")
// // }

// // // Fungsi untuk menampilkan semua karyawan
// // func tampilkanKaryawan() {
// // 	if len(daftarKaryawan) == 0 {
// // 		fmt.Println("Tidak ada karyawan yang terdaftar.")
// // 		return
// // 	}
// // 	fmt.Println("Daftar Karyawan:")
// // 	for _, karyawan := range daftarKaryawan {
// // 		fmt.Printf("ID: %d, Nama: %s, Usia: %d, Gaji: %.2f\n", karyawan.ID, karyawan.Nama, karyawan.Usia, karyawan.Gaji)
// // 	}
// // }

// // // Fungsi untuk mencari karyawan berdasarkan nama
// // func cariKaryawan(nama string) {
// // 	karyawan, ok := karyawanMap[nama]
// // 	if ok {
// // 		fmt.Printf("Karyawan ditemukan: ID: %d, Nama: %s, Usia: %d, Gaji: %.2f\n", karyawan.ID, karyawan.Nama, karyawan.Usia, karyawan.Gaji)
// // 	} else {
// // 		fmt.Println("Karyawan tidak ditemukan.")
// // 	}
// // }

// // // Fungsi untuk mengubah data karyawan menggunakan pointer
// // func ubahDataKaryawan(nama string, usiaBaru int, gajiBaru float64) {
// // 	karyawan, ok := karyawanMap[nama]
// // 	if ok {
// // 		karyawan.Usia = usiaBaru
// // 		karyawan.Gaji = gajiBaru
// // 		fmt.Println("Data karyawan berhasil diubah!")
// // 	} else {
// // 		fmt.Println("Karyawan tidak ditemukan.")
// // 	}
// // }

// // // Fungsi untuk menghapus karyawan berdasarkan nama
// // func hapusKaryawan(nama string) {
// // 	index := -1
// // 	for i, karyawan := range daftarKaryawan {
// // 		if karyawan.Nama == nama {
// // 			index = i
// // 			break
// // 		}
// // 	}

// // 	if index != -1 {
// // 		daftarKaryawan = append(daftarKaryawan[:index], daftarKaryawan[index+1:]...) // Menghapus dari slice
// // 		delete(karyawanMap, nama)                                                     // Menghapus dari map
// // 		fmt.Println("Karyawan berhasil dihapus!")
// // 	} else {
// // 		fmt.Println("Karyawan tidak ditemukan.")
// // 	}
// // }

// // // Main function untuk menjalankan program
// // func main() {
// // 	var pilihan int
// // 	for {
// // 		fmt.Println("\nMenu:")
// // 		fmt.Println("1. Tambah Karyawan")
// // 		fmt.Println("2. Tampilkan Karyawan")
// // 		fmt.Println("3. Cari Karyawan Berdasarkan Nama")
// // 		fmt.Println("4. Ubah Data Karyawan")
// // 		fmt.Println("5. Hapus Karyawan")
// // 		fmt.Println("6. Keluar")
// // 		fmt.Print("Pilih opsi: ")
// // 		fmt.Scan(&pilihan)

// // 		switch pilihan {
// // 		case 1:
// // 			var id, usia int
// // 			var nama string
// // 			var gaji float64
// // 			fmt.Print("Masukkan ID: ")
// // 			fmt.Scan(&id)
// // 			fmt.Print("Masukkan Nama: ")
// // 			fmt.Scan(&nama)
// // 			fmt.Print("Masukkan Usia: ")
// // 			fmt.Scan(&usia)
// // 			fmt.Print("Masukkan Gaji: ")
// // 			fmt.Scan(&gaji)
// // 			tambahKaryawan(id, nama, usia, gaji)
// // 		case 2:
// // 			tampilkanKaryawan()
// // 		case 3:
// // 			var nama string
// // 			fmt.Print("Masukkan Nama Karyawan yang Dicari: ")
// // 			fmt.Scan(&nama)
// // 			cariKaryawan(nama)
// // 		case 4:
// // 			var nama string
// // 			var usia int
// // 			var gaji float64
// // 			fmt.Print("Masukkan Nama Karyawan yang Ingin Diubah: ")
// // 			fmt.Scan(&nama)
// // 			fmt.Print("Masukkan Usia Baru: ")
// // 			fmt.Scan(&usia)
// // 			fmt.Print("Masukkan Gaji Baru: ")
// // 			fmt.Scan(&gaji)
// // 			ubahDataKaryawan(nama, usia, gaji)
// // 		case 5:
// // 			var nama string
// // 			fmt.Print("Masukkan Nama Karyawan yang Ingin Dihapus: ")
// // 			fmt.Scan(&nama)
// // 			hapusKaryawan(nama)
// // 		case 6:
// // 			fmt.Println("Keluar dari program.")
// // 			return
// // 		default:
// // 			fmt.Println("Pilihan tidak valid.")
// // 		}
// // 	}
// // }


// //Latihan2
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

package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 1; i <= 3; i++ {
		go func(n int) {
			fmt.Printf("Anonymous Goroutine with value %d\n", n)
		}(i)
	}
	
	time.Sleep(1 * time.Second)
}
