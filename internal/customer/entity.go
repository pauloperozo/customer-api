package customer

import (
	"time"

	"customer-api/internal/shared"
)

// Entity definition for Customer
type Customer struct {
	ID        string     `gorm:"type:text;primaryKey"`
	FirstName string     `gorm:"size:100;not null"`
	LastName  string     `gorm:"size:100;not null"`
	Email     string     `gorm:"size:100;not null;unique"`
	Language  string     `gorm:"size:10;not null"`
	BirthDate *time.Time `gorm:"type:date"`
	Status    string     `gorm:"size:20;not null"`
}

// Error definitions
var ErrCustomerAlreadyExists = shared.DomainError{Status: 409, Message: "customer already exists"}
var ErrCustomerNotFound = shared.DomainError{Status: 404, Message: "customer not found"}
