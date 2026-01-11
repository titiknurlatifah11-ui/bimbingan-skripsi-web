// //problem 1
// package main

// import (
//     "fmt"
//     "strings"
// )

// // Fungsi untuk mengecek apakah sebuah string adalah palindrom
// func isPalindrome(s string) bool {
//     // Menghapus spasi dan mengubah menjadi huruf kecil
//     s = strings.ToLower(strings.ReplaceAll(s, " ", ""))
    
//     // Membalik string
//     reversed := ""
//     for i := len(s) - 1; i >= 0; i-- {
//         reversed += string(s[i])
//     }
    
//     // Mengembalikan true jika s sama dengan reversed
//     return s == reversed
// }

// func main() {
//     var input string
    
//     // Mengambil input dari pengguna
//     fmt.Print("Masukkan kata atau angka: ")
//     fmt.Scanln(&input)
    
//     // Mengecek apakah input adalah palindrom
//     if isPalindrome(input) {
//         fmt.Println(input, "adalah palindrom.")
//     } else {
//         fmt.Println(input, "bukan palindrom.")
//     }
// }

//problem 2
// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// 	"unicode"
// )

// // Fungsi untuk menghitung vokal, konsonan, dan spasi
// func hitungKarakter(kalimat string) (int, int, int) {
// 	vokal := "aeiouAEIOU"
// 	jumlahVokal := 0
// 	jumlahKonsonan := 0
// 	jumlahSpasi := 0

// 	for _, huruf := range kalimat {
// 		if unicode.IsSpace(huruf) { // Memeriksa apakah karakter adalah spasi
// 			jumlahSpasi++
// 		} else if unicode.IsLetter(huruf) { // Memeriksa apakah karakter adalah huruf
// 			if strings.ContainsRune(vokal, huruf) {
// 				jumlahVokal++
// 			} else {
// 				jumlahKonsonan++
// 			}
// 		}
// 	}

// 	return jumlahVokal, jumlahKonsonan, jumlahSpasi
// }

// func main() {
// 	scanner := bufio.NewScanner(os.Stdin)

// 	// Input kalimat dari pengguna
// 	fmt.Print("Masukkan kalimat: ")
// 	scanner.Scan()
// 	kalimat := scanner.Text()

// 	// Menghitung vokal, konsonan, dan spasi
// 	jumlahVokal, jumlahKonsonan, jumlahSpasi := hitungKarakter(kalimat)

// 	// Menampilkan hasil
// 	fmt.Println("Jumlah huruf vokal =", jumlahVokal)
// 	fmt.Println("Jumlah huruf konsonan =", jumlahKonsonan)
// 	fmt.Println("Jumlah spasi =", jumlahSpasi)
// }


//problem 3
// package main

// import (
// 	"fmt"
// )

// // Fungsi untuk mencari dua bilangan yang jika dijumlahkan sama dengan target
// func cariIndeksDuaAngka(nums []int, target int) []int {
// 	for i := 0; i < len(nums); i++ {
// 		for j := i + 1; j < len(nums); j++ {
// 			if nums[i]+nums[j] == target {
// 				// Menambahkan satu ke indeks untuk mengikuti format soal
// 				return []int{i + 1, j + 1}
// 			}
// 		}
// 	}
// 	// Jika tidak ditemukan pasangan yang sesuai, kembalikan array kosong
// 	return []int{}
// }

// func main() {
// 	// Masukan bilangan dan target
// 	nums := []int{2, 7, 11, 15}
// 	target := 9

// 	// Memanggil fungsi dan menampilkan hasil
// 	hasil := cariIndeksDuaAngka(nums, target)
// 	if len(hasil) == 0 {
// 		fmt.Println("Tidak ada pasangan angka yang menghasilkan target")
// 	} else {
// 		fmt.Println("Output:", hasil)
// 	}
// }


//problem 4
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// Struktur untuk menyimpan informasi tiket
type Tiket struct {
	namaPemesan       string
	asalBandara       string
	tujuanBandara     string
	tanggalKeberangkatan string
	tanggalKepulangan string
	jumlahPenumpang   int
	jenisTiket        string
	asuransi          bool
}

// Fungsi untuk menghitung harga tiket berdasarkan jenis
func hitungHargaTiket(jenisTiket string, jumlahPenumpang int) int {
	var hargaPerPenumpang int
	switch strings.ToLower(jenisTiket) {
	case "ekonomi":
		hargaPerPenumpang = 1000000 // contoh harga tiket ekonomi
	case "bisnis":
		hargaPerPenumpang = 3000000 // contoh harga tiket bisnis
	case "firstclass":
		hargaPerPenumpang = 5000000 // contoh harga tiket firstclass
	default:
		fmt.Println("Jenis tiket tidak valid, menggunakan kelas ekonomi sebagai default")
		hargaPerPenumpang = 1000000
	}
	return hargaPerPenumpang * jumlahPenumpang
}

// Fungsi untuk menghitung total harga dengan opsi asuransi dan potongan
func hitungTotalHarga(baseHarga int, asuransi bool, metodePembayaran string) float64 {
	total := float64(baseHarga)
	if asuransi {
		total += total * 0.05 // Tambah 5% jika memilih asuransi
	}

	if strings.ToLower(metodePembayaran) == "kredit" {
		total -= total * 0.02 // Potong 2% jika memilih kredit
	}

	// Tambahkan 3 digit kode unik
	kodeUnik := rand.Intn(1000) // Kode unik antara 0 - 999
	total += float64(kodeUnik)

	return total
}

func main() {
	rand.Seed(time.Now().UnixNano())
	scanner := bufio.NewScanner(os.Stdin)

	// Input Nama Pemesan
	fmt.Print("Masukkan Nama Pemesan: ")
	scanner.Scan()
	namaPemesan := scanner.Text()

	// Input Asal Bandara
	fmt.Println("Masukkan Asal Bandara:")
	fmt.Println("1. Jakarta (Rp 1.000.000)")
	fmt.Println("2. Surabaya (Rp 1.500.000)")
	fmt.Println("3. Bali (Rp 2.000.000)")
	fmt.Println("4. Medan (Rp 1.800.000)")
	fmt.Print("Pilihan: ")
	scanner.Scan()
	asalBandara := scanner.Text()

	// Input Tujuan Bandara
	fmt.Println("Masukkan Tujuan Bandara:")
	fmt.Print("Pilihan: ")
	scanner.Scan()
	tujuanBandara := scanner.Text()

	// Input Tanggal Keberangkatan
	fmt.Print("Masukkan Tanggal Keberangkatan (YYYY-MM-DD): ")
	scanner.Scan()
	tanggalKeberangkatan := scanner.Text()

	// Input pilihan pulang pergi
	var tanggalKepulangan string
	fmt.Print("Apakah Anda ingin memasukkan tanggal pulang (Ya/Tidak)? ")
	scanner.Scan()
	pulangPergi := strings.ToLower(scanner.Text())
	if pulangPergi == "ya" {
		fmt.Print("Masukkan Tanggal Kepulangan (YYYY-MM-DD): ")
		scanner.Scan()
		tanggalKepulangan = scanner.Text()
	}

	// Input Jumlah Penumpang
	fmt.Print("Masukkan Jumlah Penumpang: ")
	scanner.Scan()
	jumlahPenumpang, _ := strconv.Atoi(scanner.Text())

	// Input Jenis Tiket
	fmt.Print("Masukkan Jenis Tiket (ekonomi/bisnis/firstclass): ")
	scanner.Scan()
	jenisTiket := scanner.Text()

	// Input Opsi Asuransi
	fmt.Print("Pilih Asuransi Perjalanan (Ya/Tidak): ")
	scanner.Scan()
	asuransi := strings.ToLower(scanner.Text()) == "ya"

	// Input Metode Pembayaran
	fmt.Print("Pilih Metode Pembayaran (Transfer/Kredit): ")
	scanner.Scan()
	metodePembayaran := scanner.Text()

	// Hitung harga tiket
	hargaTiket := hitungHargaTiket(jenisTiket, jumlahPenumpang)
	totalHarga := hitungTotalHarga(hargaTiket, asuransi, metodePembayaran)

	// Menampilkan ringkasan pemesanan
	fmt.Println("\n--- Ringkasan Pemesanan ---")
	fmt.Println("Nama Pemesan        :", namaPemesan)
	fmt.Println("Asal Bandara        :", asalBandara)
	fmt.Println("Tujuan Bandara      :", tujuanBandara)
	fmt.Println("Tanggal Keberangkatan:", tanggalKeberangkatan)
	if pulangPergi == "ya" {
		fmt.Println("Tanggal Kepulangan   :", tanggalKepulangan)
	}
	fmt.Println("Jumlah Penumpang    :", jumlahPenumpang)
	fmt.Println("Jenis Tiket         :", jenisTiket)
	fmt.Printf("Total Harga         : Rp%.2f\n", totalHarga)
}

//problem 5
package main

import (
	"fmt"
	"sync"
)

type Menu struct {
	Nama  string
	Harga float64
}

var menuList = []Menu{
	{"Nasi Goreng", 20000},
	{"Mie Ayam", 15000},
	{"Sate Ayam", 25000},
	{"Ayam Penyet", 30000},
}

func tampilkanMenu() {
	fmt.Println("Daftar Menu:")
	for i, menu := range menuList {
		fmt.Printf("%d. %s - Rp. %.2f\n", i+1, menu.Nama, menu.Harga)
	}
}

func hitungTotal(pilihan int, porsi int) float64 {
	return menuList[pilihan].Harga * float64(porsi)
}

func konfirmasiPesanan(total float64, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Total harga pesanan Anda adalah Rp. %.2f\n", total)
	fmt.Println("Terima kasih telah memesan!")
}

func main() {
	tampilkanMenu()
	var pilihan, porsi int

	fmt.Print("Masukkan nomor makanan yang ingin dipesan: ")
	fmt.Scan(&pilihan)
	fmt.Print("Masukkan jumlah porsi: ")
	fmt.Scan(&porsi)

	if pilihan < 1 || pilihan > len(menuList) {
		fmt.Println("Pilihan tidak valid.")
		return
	}

	total := hitungTotal(pilihan-1, porsi)
	
	var wg sync.WaitGroup
	wg.Add(1)
	go konfirmasiPesanan(total, &wg)
	wg.Wait()
}
