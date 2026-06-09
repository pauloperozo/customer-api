package main

import (
	"customer-api/internal/customer"
	"customer-api/platform/storage"
	"fmt"
)

func main() {

	fmt.Println("Starting Customer API...")
	db := storage.InitDb()

	repo := customer.NewRepository(db)

	customer_1 := &customer.Customer{
		ID:        "1",
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
	}

	err := repo.Create(customer_1)
	if err != nil {
		fmt.Println("Error creating customer:", err)
		return
	}

	list, err := repo.GetAll()
	if err != nil {
		fmt.Println("Error getting customers:", err)
		return
	}

	fmt.Println("Customers:", list)
	fmt.Println("Customer API is running...")
}
