package customer

import (
	"customer-api/platform/storage"

	"gorm.io/gorm"
)

type Repository interface {
	Create(customer *Customer) error
	GetAll() ([]Customer, error)
	GetByEmail(email string) (*Customer, error)
}

type repository struct {
	storage *storage.BaseRepository[Customer]
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{storage: storage.NewBaseRepository[Customer](db)}
}

func (repo *repository) Create(customer *Customer) error {
	_, err := repo.storage.Create(customer)
	return err
}

func (repo *repository) GetAll() ([]Customer, error) {
	return repo.storage.GetAll()
}

func (repo *repository) GetByEmail(email string) (*Customer, error) {
	return repo.storage.GetByField("email", email)
}
