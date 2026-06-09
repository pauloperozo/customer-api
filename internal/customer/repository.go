package customer

import (
	"errors"

	"gorm.io/gorm"
)

// Contract definition for Customer Repository
type Repository interface {
	Create(customer *Customer) error
	GetAll() ([]Customer, error)
	GetByEmail(email string) (*Customer, error)
}

// Implementation of the Customer Repository using GORM
type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (repo *repository) Create(customer *Customer) error {
	return repo.db.Create(customer).Error
}

func (repo *repository) GetAll() ([]Customer, error) {
	var customers []Customer
	err := repo.db.Find(&customers).Error
	return customers, err
}

func (repo *repository) GetByEmail(email string) (*Customer, error) {
	var customer Customer
	err := repo.db.Where("email=?", email).First(&customer).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &customer, nil
}
