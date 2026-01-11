package categorymodel

import (
	"go-web-native/config"
	"go-web-native/entities"
	"log"
)

func GetAll() []entities.Category {
	// Query the categories from the database
	rows, err := config.DB.Query("SELECT * FROM categories")
	if err != nil {
		log.Println("Error querying categories:", err)
		return nil // Return nil if there's an error
	}
	defer rows.Close()

	var categories []entities.Category

	// Iterate through the rows and scan the data
	for rows.Next() {
		var category entities.Category
		if err := rows.Scan(&category.Id, &category.Name, &category.CreatedAt, &category.UpdatedAt); err != nil {
			log.Println("Error scanning category:", err)
			continue // Skip this row if scanning fails
		}
		categories = append(categories, category)
	}

	// Check if there were any errors while iterating over the rows
	if err := rows.Err(); err != nil {
		log.Println("Error iterating over rows:", err)
		return nil
	}

	// Log the fetched categories for debugging purposes
	log.Printf("Fetched categories: %+v", categories)

	// Return the list of categories
	return categories
}


// Fungsi untuk menambahkan kategori baru ke dalam database
func Add(name string) error {
	_, err := config.DB.Exec("INSERT INTO categories (name) VALUES (?)", name)
	if err != nil {
		log.Printf("Error inserting category: %v", err)
		return err
	}
	return nil
}

func GetByID(id string) (entities.Category, error) {
	var category entities.Category

	// Query database
	err := config.DB.QueryRow("SELECT id, name, created_at, updated_at FROM categories WHERE id = ?", id).Scan(
		&category.Id, &category.Name, &category.CreatedAt, &category.UpdatedAt,
	)
	if err != nil {
		log.Printf("Error fetching category by Id: %v", err)
		return category, err
	}

	// Debug log
	log.Printf("Fetched Category: %+v", category)

	return category, nil
}


func Update(id string, name string) error {
	query := "UPDATE categories SET name = ?, updated_at = NOW() WHERE id = ?"
	stmt, err := config.DB.Prepare(query)
	if err != nil {
		log.Printf("Error preparing query: %v", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(name, id)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		return err
	}

	log.Printf("Category updated successfully: ID=%s, Name=%s", id, name)
	return nil
}

// Fungsi Delete di model  untuk menghapus kategori berdasarkan ID
func Delete(id string) error {
    query := "DELETE FROM categories WHERE id = ?"
    _, err := config.DB.Exec(query, id)
    if err != nil {
        log.Printf("Error deleting category from DB: %v", err)
        return err
    }
    return nil
}

// func Create(category entities.Category) bool {
// 	result, err := config.DB.exec('
// 	INSERT INTO categories (name, created_at, updated_at)
// 	VALUE (? , ?, ?)',
// 	category.Name, category.CreatedAt, category.UpdatedAt,
// 	)

// 	if err != nil {
// 		panic(err)
// 	}

// 	lastInsertId, err := result.lastInsertId()
// 	if err != nil {
// 		panic(err)
// 	}

// 	return lastInsertId > 0
// }

