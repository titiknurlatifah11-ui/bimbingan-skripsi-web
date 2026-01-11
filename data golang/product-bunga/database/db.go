package database

import (
	"database/sql"
	"errors"
	"product-bunga/models"

	_ "modernc.org/sqlite"
)

var db *sql.DB

func InitDB() error {
	var err error
	// Koneksi ke SQLite
	db, err = sql.Open("sqlite", "./product-bunga.db") // Gunakan nama database .db
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}

	// Membuat tabel produk jika belum ada
	createTableQuery := `
		CREATE TABLE IF NOT EXISTS products (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT,
			description TEXT,
			price REAL,
			stock INTEGER
		);`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		return err
	}

	return nil
}

func GetDB() *sql.DB {
	return db
}

// Fungsi untuk membuat produk baru
func CreateProduct(product *models.Product) error {
	if db == nil {
		return errors.New("database connection is not initialized")
	}

	// Query untuk menambahkan produk
	query := `INSERT INTO products (name, description, price, stock) VALUES (?, ?, ?, ?)`
	_, err := db.Exec(query, product.Name, product.Description, product.Price, product.Stock)
	if err != nil {
		return err
	}

	return nil
}
