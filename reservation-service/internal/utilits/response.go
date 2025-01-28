package utilits

import (
	"encoding/json"
	"net/http"
	"reservation-service/internal/models"
)

func WriteErrorResponse(code int, title string, message error, w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")

	errMessage := models.ErrorResponse{
		Code:    code,
		Title:   title,
		Message: message.Error(),
	}

	err := json.NewEncoder(w).Encode(&errMessage)
	if err != nil {
		return err
	}

	return nil
}
