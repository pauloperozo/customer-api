package main

import (
	"customer-api/internal/customer"
	"customer-api/platform/storage"
	"fmt"
	"log"
	"net/http"
)

func main() {

	fmt.Println("Starting Customer API...")
	db := storage.InitDb()

	repo := customer.NewRepository(db)
	service := customer.NewService(repo)
	handler := customer.NewHandler(*service)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/v1/customers", handler.Create)
	port := ":8080"
	fmt.Printf("Server is running on http://localhost%s\n", port)
	err := http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatalf("Error al encender el servidor: %v", err)
	}
}
