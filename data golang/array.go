// package main

// import (
// 	"fmt"
// )

// func main() {
// 	// Menentukan jumlah nilai ujian yang akan diinput
// 	var numScores int
// 	fmt.Print("Berapa banyak nilai ujian yang akan diinput? ")
// 	fmt.Scanln(&numScores)

// 	// Membuat array untuk menyimpan nilai-nilai ujian
// 	scores := make([]float64, numScores)

// 	// Menginput nilai-nilai ujian dari pengguna
// 	for i := 0; i < numScores; i++ {
// 		fmt.Printf("Masukkan nilai ujian ke-%d: ", i+1)
// 		fmt.Scanln(&scores[i])
// 	}

// 	// Menghitung total nilai dan rata-rata
// 	var total float64
// 	for _, score := range scores {
// 		total += score
// 	}

// 	average := total / float64(numScores)

// 	// Menampilkan rata-rata nilai
// 	fmt.Printf("Rata-rata nilai ujian: %.2f\n", average)
// }

package main

import (
	"fmt"
)

func main() {
	// Menentukan jumlah nilai ujian yang akan diinput
	var numScores int
	fmt.Print("Berapa banyak nilai ujian yang akan diinput? ")
	fmt.Scanln(&numScores)

	// Membuat array untuk menyimpan nilai-nilai ujian
	scores := make([]float64, numScores)

	// Menginput nilai-nilai ujian dari pengguna dengan validasi
	for i := 0; i < numScores; i++ {
		for {
			fmt.Printf("Masukkan nilai ujian ke-%d: ", i+1)
			_, err := fmt.Scanln(&scores[i])
			if err == nil && scores[i] >= 0 && scores[i] <= 100 {
				break // Input valid, keluar dari loop
			} else {
				fmt.Println("Nilai tidak valid. Harap masukkan angka antara 0 hingga 100.")
			}
		}
	}

	// Menghitung total nilai dan rata-rata
	var total float64
	for _, score := range scores {
		total += score
	}

	average := total / float64(numScores)

	// Menampilkan total dan rata-rata nilai
	fmt.Printf("Total nilai ujian: %.2f\n", total)
	fmt.Printf("Rata-rata nilai ujian: %.2f\n", average)
}

