package productcontroller

import (
	"go-web-native/models/categorymodel"
	"go-web-native/models/productmodel"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func Index(w http.ResponseWriter, r *http.Request) {
	// Ambil data produk dari model
	products := productmodel.GetAll()

	// Parsing template
	temp, err := template.ParseFiles("views/product/index.html")
	if err != nil {
		log.Printf("Error loading template: %v", err)
		http.Error(w, "Failed to load template: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Kirim data ke template
	data := map[string]interface{}{
		"products": products,
	}

	temp.Execute(w, data)
}

func Detail(w http.ResponseWriter, r *http.Request) {
	// Mengambil ID produk dari query parameter
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	// Mengambil data produk dari model
	products, err := productmodel.Detail(id)
	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	// Menyusun data yang akan diteruskan ke template
	data := map[string]interface{}{
		"product": products,
	}

	// Parsing file template HTML
	temp, err := template.ParseFiles("views/product/detail.html")
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Menjalankan template dengan data produk
	if err := temp.Execute(w, data); err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		temp, err := template.ParseFiles("views/product/create.html")
		if err != nil {
			log.Printf("Error loading template: %v", err)
			http.Error(w, "Failed to load template: "+err.Error(), http.StatusInternalServerError)
			return
		}

		categories := categorymodel.GetAll()
		data := map[string]interface{}{
			"categories": categories,
		}

		temp.Execute(w, data)
		return
	}

	if r.Method == http.MethodPost {
		// Ambil data dari form
		name := r.FormValue("name")
		categoryID, _ := strconv.Atoi(r.FormValue("category_id"))
		stock, _ := strconv.ParseInt(r.FormValue("stock"), 10, 64)
		description := r.FormValue("description")


		// Simpan ke database melalui model
		err := productmodel.Insert(name, categoryID, stock, description)
		if err != nil {
			log.Printf("Error inserting product: %v", err)
			http.Error(w, "Failed to insert product: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Redirect ke halaman product
		http.Redirect(w, r, "/product", http.StatusSeeOther)
	}
}

func Edit(w http.ResponseWriter, r *http.Request) {
	// Ambil template
	temp, err := template.ParseFiles("views/product/edit.html")
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Ambil ID produk
	idString := r.URL.Query().Get("id")
	if idString == "" {
		http.Error(w, "Product ID is missing", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	// Ambil data produk
	products, err := productmodel.Detail(id)
	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	// Ambil data kategori
	categories := categorymodel.GetAll()

	// Data untuk template
	data := map[string]interface{}{
		"categories": categories,
		"product":    products,
	}

	// Jika metode adalah POST, lakukan update
	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		categoryID, _ := strconv.Atoi(r.FormValue("category_id"))
		stock, _ := strconv.ParseInt(r.FormValue("stock"), 10, 64)
		description := r.FormValue("description")

		// Update produk
		err := productmodel.Update(id, name, categoryID, stock, description)
		if err != nil {
			http.Error(w, "Failed to update product", http.StatusInternalServerError)
			return
		}
		
		// Redirect ke halaman produk
		http.Redirect(w, r, "/product", http.StatusSeeOther)
		return
		}

		// Jalankan template (GET request)
		if err := temp.Execute(w, data); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
}
