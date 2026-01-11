//1.ARRAY
// package main

// import "fmt"

// func main() {
// 	var numbers [5]int
// 	numbers[0] = 10
//     numbers[1] = 20
//     numbers[2] = 30
// // Mengubah elemen
//     numbers[2] = 50
// // Mengiterasi array
//     for i, num := range numbers {
//         fmt.Printf("Element %d: %d\n", i, num)
//     }
// }

//modifikasi
// package main

// import "fmt"

// func main() {
// 	var numbers [5]int

// 	// Mengisi array dengan input dari pengguna
// 	for i := 0; i < len(numbers); i++ {
// 		fmt.Printf("Masukkan elemen ke-%d: ", i+1)
// 		fmt.Scanln(&numbers[i])
// 	}

// 	// Mengubah elemen tertentu (misalnya elemen ke-3)
// 	numbers[2] = 50

// 	// Mengiterasi array dan menampilkan elemen
// 	var total int
// 	for i, num := range numbers {
// 		fmt.Printf("Element %d: %d\n", i, num)
// 		total += num
// 	}

// 	// Menampilkan total
// 	fmt.Printf("Total: %d\n", total)
// }

//STUDIKASUS array
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

//modifikasi STUDIKASUS array
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

// 	// Menginput nilai-nilai ujian dari pengguna dengan validasi
// 	for i := 0; i < numScores; i++ {
// 		for {
// 			fmt.Printf("Masukkan nilai ujian ke-%d: ", i+1)
// 			_, err := fmt.Scanln(&scores[i])
// 			if err == nil && scores[i] >= 0 && scores[i] <= 100 {
// 				break // Input valid, keluar dari loop
// 			} else {
// 				fmt.Println("Nilai tidak valid. Harap masukkan angka antara 0 hingga 100.")
// 			}
// 		}
// 	}

// 	// Menghitung total nilai dan rata-rata
// 	var total float64
// 	for _, score := range scores {
// 		total += score
// 	}

// 	average := total / float64(numScores)

// 	// Menampilkan total dan rata-rata nilai
// 	fmt.Printf("Total nilai ujian: %.2f\n", total)
// 	fmt.Printf("Rata-rata nilai ujian: %.2f\n", average)
// }