package customer

import "customer-api/internal/shared"

var ErrCustomerAlreadyExists = shared.DomainError{Status: 409, Message: "customer already exists"}
var ErrCustomerNotFound = shared.DomainError{Status: 404, Message: "customer not found"}
var ErrCustomerInvalidData = shared.DomainError{Status: 400, Message: "invalid customer data"}
