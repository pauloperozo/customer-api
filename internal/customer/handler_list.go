package customer

import (
	"encoding/json"
	"net/http"

	"customer-api/internal/shared"
)

// List handles GET /customers
// @Summary      Obtener clientes
// @Description  Obtiene un cliente por ID o lista todos los clientes.
// @Tags         Customers
// @Accept       json
// @Produce      json
// @Success      200       {array}   ResponseCustomer
// @Failure      500       {object}  shared.ErrorResponse "Error interno del servidor"
// @Router       /customers [get]
func (handler *Handler) ListCustomers(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")
	customers, err := handler.service.GetAllCustomers()
	if err != nil {
		shared.HandleError(res, err)
		return
	}

	response := make([]ResponseCustomer, 0, len(customers))
	for _, customer := range customers {
		response = append(response, ResponseCustomer{
			ID:        customer.ID,
			FirstName: customer.FirstName,
			LastName:  customer.LastName,
			Email:     customer.Email,
			Language:  customer.Language,
		})
	}

	json.NewEncoder(res).Encode(response)

}
