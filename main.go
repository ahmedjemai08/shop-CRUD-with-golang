package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Product struct {
	ID       string    `json:"id"`
	Type     string    `json:"type"`
	Price    int       `json:"price"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var products []Product

func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)

}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	for index, item := range products {

		if item.ID == param["ID"] {
			products = append(products[:index], products[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(products)
}
func getProductById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	for _, item := range products {
		if item.ID == param["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

}

func createProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product Product
	_ = json.NewDecoder(r.Body).Decode(&product)
	product.ID = strconv.Itoa(rand.Intn(10000))
	products = append(products, product)
	json.NewEncoder(w).Encode(products)

}

func updateProduct(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)

	for index, item := range products {
		if item.ID == param["id"] {
			products = append(products[:index], products[index+1:]...)
			var product Product
			_ = json.NewDecoder(r.Body).Decode(&product)
			product.ID = param["id"]
			products = append(products, product)
			json.NewEncoder(w).Encode(product)
		}
	}

}

func main() {

	r := mux.NewRouter()

	products = append(products, Product{ID: "1", Type: "T-shirt", Price: 80, Director: &Director{Firstname: "Black", Lastname: "M"}})
	products = append(products, Product{ID: "2", Type: "Sneakers", Price: 90, Director: &Director{Firstname: "RED", Lastname: "41"}})
	r.HandleFunc("/product", getProducts).Methods("GET")
	r.HandleFunc("/product/{id}", getProductById).Methods("GET")
	r.HandleFunc("/product", createProduct).Methods("POST")
	r.HandleFunc("/product/{id}", updateProduct).Methods("PUT")
	r.HandleFunc("/product/{id}", deleteProduct).Methods("DELETE")

	fmt.Printf("Starting Server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))

}
