package main

import (
	"fmt"
	"log"
	"net/http"
	"product-bunga/database"
	"product-bunga/handlers"
)

func main() {
	// Inisialisasi koneksi ke database
	err := database.InitDB()
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	// Endpoint utama
	http.HandleFunc("/product-bunga", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Selamat datang di halaman produk bunga!")
	})

	// Routing untuk fitur CRUD
	http.HandleFunc("/products", handlers.GetProducts)       // Get all products
	http.HandleFunc("/products/", handlers.GetProduct)       // Get product by ID
	http.HandleFunc("/products/create", handlers.CreateProduct) // Create a new product
	http.HandleFunc("/products/update", handlers.UpdateProduct) // Update product
	http.HandleFunc("/products/delete", handlers.DeleteProduct) // Delete product

	// Menentukan IP dan Port
	ip := "127.0.0.1" 
	port := "8081"

	// Menjalankan server
	fmt.Printf("Server is running at http://%s:%s...\n", ip, port)
	log.Fatal(http.ListenAndServe(ip+":"+port, nil))
}
