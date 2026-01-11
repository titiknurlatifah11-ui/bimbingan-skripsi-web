package main

import (
	"go-web-native/config"
    "go-web-native/controller/homecontroller"
	"go-web-native/controller/categorycontroller" 
	"go-web-native/controller/productcontroller"  
	"log"      
	"net/http"           
)

func main() {
    config.ConnectDB() // Pastikan ConnectDB dipanggil
    http.HandleFunc("/", homecontroller.Welcome)
	http.HandleFunc("/categories", categorycontroller.Index)
	http.HandleFunc("/categories/add", categorycontroller.Add)
	http.HandleFunc("/categories/edit", categorycontroller.Edit)
	http.HandleFunc("/categories/delete", categorycontroller.Delete)

	http.HandleFunc("/product", productcontroller.Index)
	http.HandleFunc("/product/add", productcontroller.Add)
	http.HandleFunc("/product/detail", productcontroller.Detail)
	http.HandleFunc("/product/edit", productcontroller.Edit)

	log.Println("Server running on port 8000")
    http.ListenAndServe(":8000", nil)
}
