package customer

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

// Create handles POST /customers
// @Summary      Crear un nuevo cliente
// @Description  Registra un cliente en el sistema validando que el email sea único.
// @Tags         Customers
// @Accept       json
// @Produce      json
// @Param        customer  body      CreateCustomerRequest  true  "Datos del cliente a crear"
// @Success      200       {object}  ResponseCustomer
// @Failure      400       {object}  map[string]string  "JSON inválido o malformado"
// @Failure      409       {object}  map[string]string  "El email ya está registrado"
// @Failure      500       {object}  map[string]string  "Error interno del servidor"
// @Router       /customers [post]
func (handler *Handler) Create(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")
	var body CreateCustomerRequest
	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	newCustomer := &Customer{
		ID:        uuid.NewString(),
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Email:     body.Email,
		Language:  body.Language,
	}

	err = handler.service.CreateCustomer(newCustomer)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
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

// List handles GET /customers
// @Summary      Obtener clientes
// @Description  Obtiene un cliente por ID o lista todos los clientes.
// @Tags         Customers
// @Accept       json
// @Produce      json
// @Success      200       {array}   ResponseCustomer
// @Failure      400       {object}  map[string]string  "JSON inválido o malformado"
// @Failure      409       {object}  map[string]string  "El email ya está registrado"
// @Failure      500       {object}  map[string]string  "Error interno del servidor"
// @Router       /customers [get]
func (handler *Handler) List(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")
	customers, err := handler.service.GetAllCustomers()
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
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
