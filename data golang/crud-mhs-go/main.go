package main

import (
	"net/http"

)

func main() {
	// Use http.NewServeMux() instead of http.NewServerMux
	mux := http.NewServeMux()

	// Example route
	Mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
	

	// Start the server
	http.ListenAndServe(":8081", mux)
}
