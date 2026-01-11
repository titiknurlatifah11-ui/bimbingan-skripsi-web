package categorycontroller

import (  
	"go-web-native/entities"
	"go-web-native/models/categorymodel"
	"html/template"
	"log"
	"net/http"
	"time"
)

// Fungsi untuk menampilkan semua kategori
func Index(w http.ResponseWriter, r *http.Request) {
	// Mengambil semua kategori dari model
	categories := categorymodel.GetAll()

	// Menyiapkan data untuk dikirim ke template
	data := map[string]interface{}{
		"categories": categories,
	}

	// Memuat dan mengeksekusi template
	temp, err := template.ParseFiles("views/category/index.html")
	if err != nil {
		log.Printf("Error loading template: %v", err)
		http.Error(w, "Failed to load template: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Menyusun template dengan data
	err = temp.Execute(w, data)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Failed to execute template: "+err.Error(), http.StatusInternalServerError)
	}
}


// Fungsi untuk menambahkan kategori baru
func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/category/create.html")
		if err != nil {
			panic(err)
		}

		temp.Execute(w, nil)
	}
	if r.Method == "POST" {
		var category entities.Category

		category.Name = r.FormValue("name")
		category.CreatedAt = time.Now()
		category.UpdatedAt = time.Now()
	}

}

// Fungsi Edit untuk mengedit kategori (belum diimplementasikan)
func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// Ambil ID dari query parameter
		id := r.URL.Query().Get("id")
		if id == "" {
			log.Println("No ID provided in URL query")
			http.Error(w, "ID is required", http.StatusBadRequest)
			return
		}

		// Fetch kategori dari database
		category, err := categorymodel.GetByID(id)
		if err != nil {
			log.Printf("Error fetching category by Id: %v", err)
			http.Error(w, "Failed to fetch category", http.StatusInternalServerError)
			return
		}

		// Memeriksa apakah kategori ditemukan
		if category.Name == "" {
			log.Println("Category not found")
			http.Error(w, "Category not found", http.StatusNotFound)
			return
		}

		// Kirim data ke template
		data := map[string]interface{}{
			"category": category,
		}

		// Memuat template
		temp, err := template.ParseFiles("views/category/edit.html")
		if err != nil {
			log.Printf("Error loading template: %v", err)
			http.Error(w, "Failed to load template", http.StatusInternalServerError)
			return
		}

		// Menjalankan template dengan data kategori
		err = temp.Execute(w, data)
		if err != nil {
			log.Printf("Error executing template: %v", err)
			http.Error(w, "Failed to render template", http.StatusInternalServerError)
		}
	} else if r.Method == "POST" {
		// Ambil nilai dari form
		id := r.FormValue("id")
		name := r.FormValue("name")

		// Debug log
		log.Printf("Received Id: %s, Name: %s", id, name)

		// Validasi data
		if id == "" || name == "" {
			log.Println("Id or Name is empty")
			http.Error(w, "Id and Name are required", http.StatusBadRequest)
			return
		}

		// Update kategori di database
		err := categorymodel.Update(id, name)
		if err != nil {
			log.Printf("Error updating category: %v", err)
			http.Error(w, "Failed to update category", http.StatusInternalServerError)
			return
		}

		// Redirect ke halaman index setelah berhasil
		http.Redirect(w, r, "/categories", http.StatusSeeOther)
	}
}


// Fungsi Delete untuk menghapus kategori menggunakan metode GET
func Delete(w http.ResponseWriter, r *http.Request) {
    // Pastikan menggunakan GET untuk penghapusan
    if r.Method == "GET" {
        // Ambil ID kategori dari query parameter
        id := r.URL.Query().Get("id")
        if id == "" {
            log.Println("No ID provided in query")
            http.Error(w, "ID is required", http.StatusBadRequest)
            return
        }

        // Panggil fungsi model untuk menghapus kategori berdasarkan ID
        err := categorymodel.Delete(id)
        if err != nil {
            log.Printf("Error deleting category: %v", err)
            http.Error(w, "Failed to delete category", http.StatusInternalServerError)
            return
        }

        // Redirect kembali ke halaman kategori setelah berhasil menghapus
        http.Redirect(w, r, "/categories", http.StatusSeeOther)
    } else {
        // Jika bukan metode GET, kirimkan pesan error
        log.Println("Invalid method for deleting category")
        http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
    }
}
