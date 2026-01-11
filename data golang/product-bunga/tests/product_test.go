package tests

import (
    "bytes"
    "encoding/json"
    "net/http"
    "product-bunga/models"
    "testing"
)

func TestCreateProduct(t *testing.T) {
    product := models.Product{
        Name:        "Bunga Mawar",
        Description: "Bunga merah",
        Price:       10000,
        Stock:       50,
    }

    body, _ := json.Marshal(product)

    // Membuat permintaan HTTP POST
    req, err := http.NewRequest("POST", "http://127.0.0.1:8081/products/create", bytes.NewBuffer(body))
    if err != nil {
        t.Fatal(err)
    }

    // Menambahkan header Content-Type
    req.Header.Set("Content-Type", "application/json")

    // Mengirim permintaan
    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        t.Fatal(err)
    }
    defer resp.Body.Close()

    // Memeriksa status code
    if resp.StatusCode != http.StatusCreated {
        t.Errorf("Expected status code 201, got %d", resp.StatusCode)
    }
}
