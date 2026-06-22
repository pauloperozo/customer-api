package customer

import (
	"encoding/json"
	"net/http"

	"customer-api/internal/shared"
)

// Create handles POST /customers
// @Summary      Crear un nuevo cliente
// @Description  Registra un cliente en el sistema validando que el email sea único.
// @Tags         Customers
// @Accept       json
// @Produce      json
// @Param        customer  body      CreateCustomerRequest  true  "Datos del cliente a crear"
// @Success      200       {object}  ResponseCustomer
// @Failure      400       {object}  shared.ErrorResponse "JSON inválido o malformado"
// @Failure      409       {object}  shared.ErrorResponse "El email ya está registrado"
// @Failure      500       {object}  shared.ErrorResponse "Error interno del servidor"
// @Router       /customers [post]
func (handler *Handler) CreateCustomer(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")
	var body CreateCustomerRequest
	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
		shared.HandleError(res, shared.DomainError{
			Status:  http.StatusBadRequest,
			Message: "JSON inválido o malformado",
		})
		return
	}

	if err := handler.validate.Struct(body); err != nil {
		shared.HandleError(res, ErrCustomerInvalidData)
		return
	}

	newCustomer := &Customer{
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Email:     body.Email,
		Language:  body.Language,
	}

	err = handler.service.CreateCustomer(newCustomer)
	if err != nil {
		shared.HandleError(res, err)
		return
	}

	response := CreateCustomerResponse{
		ID:        newCustomer.ID,
		FirstName: newCustomer.FirstName,
		LastName:  newCustomer.LastName,
		Email:     newCustomer.Email,
		Language:  newCustomer.Language,
	}

	json.NewEncoder(res).Encode(response)

}
