package productmodel

import (
	"go-web-native/config"
	"go-web-native/entities"
	"log"
	"errors"
	"database/sql"
)

func GetAll() []entities.Products {
	rows, err := config.DB.Query(`
	SELECT
		products.id,
		products.name,
		categories.name AS category_name,
		products.stock,
		products.description,
		products.created_at,
		products.updated_at
	FROM products
	JOIN categories ON products.category_id = categories.id
	`)

	if err != nil {
		log.Printf("Error querying database: %v", err)
		return nil
	}

	defer rows.Close()

	var products []entities.Products

	for rows.Next() {
		var product entities.Products
		err := rows.Scan(
			&product.Id,
			&product.Name,
			&product.Category.Name,
			&product.Stock,
			&product.Description,
			&product.CreatedAt,
			&product.UpdatedAt,
		)

		if err != nil {
			log.Printf("Error scanning row: %v", err)
			continue
		}

		products = append(products, product)
	}

	if len(products) == 0 {
		log.Println("No products found in database.")
	}

	return products
}

func Insert(name string, categoryID int, stock int64, description string) error {
	query := `
		INSERT INTO products (name, category_id, stock, description, created_at, updated_at)
		VALUES (?, ?, ?, ?, NOW(), NOW())`
	_, err := config.DB.Exec(query, name, categoryID, stock, description)
	return err
}

func Detail(id int) (*entities.Products, error) {
	// Query untuk mengambil detail produk dari database
	query := `SELECT p.id, p.name, c.name AS category_name, p.stock, p.description, p.created_at, p.updated_at
          FROM products p
          JOIN categories c ON p.category_id = c.id
          WHERE p.id = ?`
	var product entities.Products
	err := config.DB.QueryRow(query, id).Scan(&product.Id, &product.Name, &product.Category.Name, &product.Stock, &product.Description, &product.CreatedAt, &product.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("product not found")
		}
		log.Printf("Error querying product: %v", err)
		return nil, err
	}
	return &product, nil
}

func Update(id int, name string, categoryID int, stock int64, description string) error {
	// Query SQL untuk update produk
	query := `
		UPDATE products
		SET name = ?, category_id = ?, stock = ?, description = ?, updated_at = NOW()
		WHERE id = ?`
	_, err := config.DB.Exec(query, name, categoryID, stock, description, id)
	return err
}

