package utils

import (
	"encoding/json"
	"net/http"

	"user-service/internal/models"
)

func WriteErrorResponse(code int, err error, message string, w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")

	response := &models.ErrorResponse{
		ErrorCode:    code,
		ErrorTitle:   err.Error(),
		ErrorMessage: message,
	}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		return err
	}

	return nil
}
