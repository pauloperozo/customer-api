package customer

import (
	"fmt"

	"github.com/google/uuid"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (service *Service) CreateCustomer(customer *Customer) error {
	existingCustomer, err := service.repo.GetByEmail(customer.Email)
	if err != nil {
		return fmt.Errorf("error checking existing customer: %w", err)
	}

	if existingCustomer != nil {
		return ErrCustomerAlreadyExists
	}

	if customer.ID == "" {
		customer.ID = uuid.NewString()
	}

	return service.repo.Create(customer)
}

func (service *Service) GetAllCustomers() ([]Customer, error) {
	return service.repo.GetAll()
}
