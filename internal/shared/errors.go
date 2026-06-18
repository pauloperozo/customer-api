package shared

import (
	"encoding/json"
	"errors"
	"net/http"
)

type DomainError struct {
	Status  int
	Message string
}

func (error DomainError) Error() string {
	return error.Message
}

func HandleError(w http.ResponseWriter, err error) {
	var domainErr DomainError
	status := http.StatusInternalServerError
	message := "internal server error"

	if errors.As(err, &domainErr) {
		status = domainErr.Status
		message = domainErr.Message
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(map[string]string{"message": message})
}
