//Studi Kasus: Sistem Pembayaran dengan Diskon Dinamis 
// package main

// import (
// 	"fmt"
// 	"math"
// )

// // // Fungsi untuk menghitung total harga dari barang yang dibeli
// func hitungTotalHarga(hargaBarang []float64, jumlahBarang []int) (float64, error) {
// 	if len(hargaBarang) != len(jumlahBarang) {
// 		return 0, fmt.Errorf("jumlah barang dan harga tidak sesuai")
// 	}
// 	total := 0.0
// 	for i := 0; i < len(hargaBarang); i++ {
// 		total += hargaBarang[i] * float64(jumlahBarang[i])
// 	}
// 	return total, nil
// }

// // // Fungsi untuk menerapkan diskon tetap
// func diskonTetap(total float64, potongan float64) float64 {
// 	return math.Max(total-potongan, 0) // Tidak boleh kurang dari nol
// }

// // // Fungsi untuk menerapkan diskon persentase
// func diskonPersentase(total float64, persentase float64) float64 {
// 	potongan := total * (persentase / 100)
// 	return math.Max(total-potongan, 0) // Tidak boleh kurang dari nol
// }

// // // Closure untuk menerapkan diskon kombinasi
// func diskonKombinasi(diskon1 func(float64) float64, diskon2 func(float64) float64) func(float64) float64 {
// 	return func(total float64) float64 {
// 		totalSetelahDiskon1 := diskon1(total)
// 		totalSetelahDiskon2 := diskon2(totalSetelahDiskon1)
// 		return totalSetelahDiskon2
// 	}
// }

// // func main() {
// // 	// Daftar harga barang dan jumlah yang dibeli
// 	hargaBarang := []float64{200000, 55000, 70000} // Contoh: barang A = 100rb, B = 50rb, C = 75rb
// 	jumlahBarang := []int{2, 1, 3}                 // Jumlah barang A = 2, B = 1, C = 3

// // 	// Menghitung total harga barang
// 	totalHarga, err := hitungTotalHarga(hargaBarang, jumlahBarang)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	fmt.Printf("Total Harga Sebelum Diskon: Rp %.2f\n", totalHarga)

// // 	// Menerapkan diskon tetap Rp 20.000
// 	totalSetelahDiskonTetap := diskonTetap(totalHarga, 20000)
// 	fmt.Printf("Total Setelah Diskon Tetap Rp 20.000: Rp %.2f\n", totalSetelahDiskonTetap)

// // 	// Menerapkan diskon persentase 10%
// 	totalSetelahDiskonPersen := diskonPersentase(totalHarga, 10)
// 	fmt.Printf("Total Setelah Diskon Persen 10%%: Rp %.2f\n", totalSetelahDiskonPersen)

// // 	// Menggabungkan diskon tetap dan diskon persentase (Rp 20.000 + 10%)
// 	diskonGabungan := diskonKombinasi(
// 		func(total float64) float64 { return diskonTetap(total, 20000) },
// 		func(total float64) float64 { return diskonPersentase(total, 10) },
// 	)

// 	totalSetelahDiskonGabungan := diskonGabungan(totalHarga)
// 	fmt.Printf("Total Setelah Diskon Kombinasi (Rp 20.000 + 10%%): Rp %.2f\n", totalSetelahDiskonGabungan)
// }

//TANTANGAN LANJUTAN STUDI KASUS 1
// package main

// import (
// 	"fmt"
// 	"math"
// )

// // Struct untuk menyimpan data pelanggan
// type Pelanggan struct {
// 	Nama        string
// 	StatusVIP   bool
// 	DiskonVIP   float64 // Diskon persentase khusus untuk pelanggan VIP
// }

// // Fungsi untuk menghitung total harga dari barang yang dibeli
// func hitungTotalHarga(hargaBarang []float64, jumlahBarang []int) (float64, error) {
// 	if len(hargaBarang) != len(jumlahBarang) {
// 		return 0, fmt.Errorf("jumlah barang dan harga tidak sesuai")
// 	}
// 	total := 0.0
// 	for i := 0; i < len(hargaBarang); i++ {
// 		total += hargaBarang[i] * float64(jumlahBarang[i])
// 	}
// 	return total, nil
// }

// // Fungsi untuk menerapkan diskon tetap
// func diskonTetap(total float64, potongan float64) float64 {
// 	return math.Max(total-potongan, 0) // Tidak boleh kurang dari nol
// }

// // Fungsi untuk menerapkan diskon persentase
// func diskonPersentase(total float64, persentase float64) float64 {
// 	potongan := total * (persentase / 100)
// 	return math.Max(total-potongan, 0) // Tidak boleh kurang dari nol
// }

// // Fungsi untuk membatasi diskon maksimum
// func batasMaksimumDiskon(total float64, diskon float64, batasMaks float64) float64 {
// 	if diskon > batasMaks {
// 		return total - batasMaks
// 	}
// 	return total - diskon
// }

// // Closure untuk menerapkan diskon kombinasi
// func diskonKombinasi(diskon1 func(float64) float64, diskon2 func(float64) float64) func(float64) float64 {
// 	return func(total float64) float64 {
// 		totalSetelahDiskon1 := diskon1(total)
// 		totalSetelahDiskon2 := diskon2(totalSetelahDiskon1)
// 		return totalSetelahDiskon2
// 	}
// }

// // Fungsi untuk menerapkan diskon berbasis kuantitas (contoh: beli 3 gratis 1)
// func diskonKuantitas(hargaBarang float64, jumlahBarang int, minimalPembelian int) float64 {
// 	if jumlahBarang >= minimalPembelian {
// 		barangGratis := jumlahBarang / minimalPembelian
// 		return float64(jumlahBarang-barangGratis) * hargaBarang
// 	}
// 	return float64(jumlahBarang) * hargaBarang
// }

// // Fungsi untuk menghitung diskon pelanggan VIP
// func diskonPelanggan(total float64, pelanggan Pelanggan) float64 {
// 	if pelanggan.StatusVIP {
// 		fmt.Printf("Pelanggan VIP: %s mendapat diskon tambahan sebesar %.2f%%\n", pelanggan.Nama, pelanggan.DiskonVIP)
// 		return diskonPersentase(total, pelanggan.DiskonVIP)
// 	}
// 	return total
// }

// func main() {
// 	// Daftar harga barang dan jumlah yang dibeli
// 	hargaBarang := []float64{200000, 55000, 70000} // Contoh: barang A = 200rb, B = 55rb, C = 70rb
// 	jumlahBarang := []int{3, 1, 3}                 // Jumlah barang A = 3, B = 1, C = 3

// 	// Menghitung total harga barang
// 	totalHarga, err := hitungTotalHarga(hargaBarang, jumlahBarang)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	fmt.Printf("Total Harga Sebelum Diskon: Rp %.2f\n", totalHarga)

// 	// Menerapkan diskon tetap Rp 20.000
// 	totalSetelahDiskonTetap := diskonTetap(totalHarga, 20000)
// 	fmt.Printf("Total Setelah Diskon Tetap Rp 20.000: Rp %.2f\n", totalSetelahDiskonTetap)

// 	// Menerapkan diskon persentase 10%
// 	totalSetelahDiskonPersen := diskonPersentase(totalHarga, 10)
// 	fmt.Printf("Total Setelah Diskon Persen 10%%: Rp %.2f\n", totalSetelahDiskonPersen)

// 	// Menggabungkan diskon tetap dan diskon persentase (Rp 20.000 + 10%)
// 	diskonGabungan := diskonKombinasi(
// 		func(total float64) float64 { return diskonTetap(total, 20000) },
// 		func(total float64) float64 { return diskonPersentase(total, 10) },
// 	)

// 	totalSetelahDiskonGabungan := diskonGabungan(totalHarga)
// 	fmt.Printf("Total Setelah Diskon Kombinasi (Rp 20.000 + 10%%): Rp %.2f\n", totalSetelahDiskonGabungan)

// 	// Menerapkan diskon berbasis kuantitas (beli 3 gratis 1 untuk barang pertama)
// 	totalSetelahDiskonKuantitas := diskonKuantitas(hargaBarang[0], jumlahBarang[0], 3)
// 	fmt.Printf("Total Harga Barang Pertama Setelah Diskon Kuantitas (Beli 3 Gratis 1): Rp %.2f\n", totalSetelahDiskonKuantitas)

//     // Membatasi diskon maksimum (maksimum diskon 50.000)
//     diskon := 70000 // Diskon sebesar 70.000
//     totalSetelahBatasDiskon := batasMaksimumDiskon(totalHarga, float64(diskon), 50000) // Konversi diskon ke float64
//     fmt.Printf("Total Setelah Diskon Dengan Batas Maksimum Rp 50.000: Rp %.2f\n", totalSetelahBatasDiskon)


// 	// Data pelanggan
// 	pelanggan1 := Pelanggan{
// 		Nama:      "Titik Nur Latifah",
// 		StatusVIP: true,
// 		DiskonVIP: 5, // Diskon 5% untuk VIP
// 	}

// 	// Menerapkan diskon VIP
// 	totalSetelahDiskonVIP := diskonPelanggan(totalHarga, pelanggan1)
// 	fmt.Printf("Total Setelah Diskon Pelanggan VIP: Rp %.2f\n", totalSetelahDiskonVIP)
// }


//Studi Kasus: Aplikasi Pengolahan Data Mahasiswa Menggunakan Fungsi di Go 
// package main

// import (
// 	"fmt"
// )

// type Mahasiswa struct {
// 	NIM     string
// 	Nama    string
// 	Usia    int
// 	Jurusan string
// }

// func tambahMahasiswa(data *[]Mahasiswa, mhs Mahasiswa) {
// 	*data = append(*data, mhs)
// }

// func tampilkanMahasiswa(data []Mahasiswa) {
// 	for _, mhs := range data {
// 		fmt.Printf("NIM: %s, Nama: %s, Usia: %d, Jurusan: %s\n", mhs.NIM, mhs.Nama, mhs.Usia, mhs.Jurusan)
// 	}
// }

// func cariMahasiswa(data []Mahasiswa, nim string) *Mahasiswa {
// 	for i, mhs := range data {
// 		if mhs.NIM == nim {
// 			return &data[i]
// 		}
// 	}
// 	return nil
// }

// func hapusMahasiswa(data *[]Mahasiswa, nim string) bool {
// 	for i, mhs := range *data {
// 		if mhs.NIM == nim {
// 			*data = append((*data)[:i], (*data)[i+1:]...)
// 			return true
// 		}
// 	}
// 	return false
// }

// func filterMahasiswa(data []Mahasiswa, filterFunc func(Mahasiswa) bool) []Mahasiswa {
// 	var filtered []Mahasiswa
// 	for _, mhs := range data {
// 		if filterFunc(mhs) {
// 			filtered = append(filtered, mhs)
// 		}
// 	}
// 	return filtered
// }

// func main() {
// 	var dataMahasiswa []Mahasiswa

// 	// Menambahkan data mahasiswa
// 	tambahMahasiswa(&dataMahasiswa, Mahasiswa{
// 		NIM:     "1152700006",
// 		Nama:    "Alayha",
// 		Usia:    19,
// 		Jurusan: "Informatika",
// 	})

// 	tambahMahasiswa(&dataMahasiswa, Mahasiswa{
// 		NIM:     "11152700049",
// 		Nama:    "Ratu",
// 		Usia:    20,
// 		Jurusan: "Sistem Informasi",
// 	})

// 	tambahMahasiswa(&dataMahasiswa, Mahasiswa{
// 		NIM:     "1152700011",
// 		Nama:    "Vina",
// 		Usia:    19,
// 		Jurusan: "Informatika",
// 	})

// 	// Menampilkan semua data mahasiswa
// 	fmt.Println("=== Data Mahasiswa ===")
// 	tampilkanMahasiswa(dataMahasiswa)

// 	// Mencari mahasiswa berdasarkan NIM
// 	nimDicari := "1152700011"
// 	mhs := cariMahasiswa(dataMahasiswa, nimDicari)
// 	if mhs != nil {
// 		fmt.Printf("\nMahasiswa dengan NIM %s ditemukan:\n", nimDicari)
// 		fmt.Printf("Nama: %s, Usia: %d, Jurusan: %s\n", mhs.Nama, mhs.Usia, mhs.Jurusan)
// 	} else {
// 		fmt.Printf("\nMahasiswa dengan NIM %s tidak ditemukan.\n", nimDicari)
// 	}

// 	// Menghapus mahasiswa berdasarkan NIM
// 	nimDihapus := "1152700006"
// 	if hapusMahasiswa(&dataMahasiswa, nimDihapus) {
// 		fmt.Printf("\nMahasiswa dengan NIM %s berhasil dihapus.\n", nimDihapus)
// 	} else {
// 		fmt.Printf("\nMahasiswa dengan NIM %s tidak ditemukan.\n", nimDihapus)
// 	}

// 	// Menampilkan data mahasiswa setelah penghapusan
// 	fmt.Println("\n=== Data Mahasiswa Setelah Penghapusan ===")
// 	tampilkanMahasiswa(dataMahasiswa)

// 	// Memfilter mahasiswa jurusan Informatika menggunakan fungsi anonim
// 	fmt.Println("\n=== Mahasiswa Jurusan Informatika ===")
// 	mhsInformatika := filterMahasiswa(dataMahasiswa, func(mhs Mahasiswa) bool {
// 		return mhs.Jurusan == "Informatika"
// 	})
// 	tampilkanMahasiswa(mhsInformatika)
// }
