package main

import (
	"customer-api/platform/storage"
	"fmt"
)

func main() {

	fmt.Println("Starting Customer API...")
	db := storage.InitDb()
	_ = db

	fmt.Println("Customer API is running...")
}
