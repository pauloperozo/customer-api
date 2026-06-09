package customer

import "gorm.io/gorm"

type Repository interface {
	Create(customer *Customer) error
	GetAll() ([]Customer, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(customer *Customer) error {
	return r.db.Create(customer).Error
}

func (r *repository) GetAll() ([]Customer, error) {
	var customers []Customer
	err := r.db.Find(&customers).Error
	return customers, err
}
