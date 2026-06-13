package main

import (
	_ "customer-api/docs"
	"customer-api/internal/customer"
	"customer-api/platform/storage"
	"fmt"
	"log"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger/v2"
)

// @title           Customer API
// @version         1.0
// @description     Este es un servidor de ejemplo para la gestión de clientes en Go.
// @host            localhost:8080
// @BasePath        /api/v1
func main() {

	fmt.Println("Starting Customer API...")
	db := storage.InitDb()

	repo := customer.NewRepository(db)
	service := customer.NewService(repo)
	handler := customer.NewHandler(*service)

	mux := http.NewServeMux()
	mux.Handle("/swagger/", httpSwagger.WrapHandler)
	mux.HandleFunc("POST /api/v1/customers", handler.Create)
	port := ":8080"
	fmt.Printf("Server is running on http://localhost%s\n", port)
	err := http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatalf("Error al encender el servidor: %v", err)
	}
}
