package helpers

import (
	"encoding/json"
	"net/http"
)

type ApiResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Errors  any    `json:"errors,omitempty"`
}

func WriteApiResponse(w http.ResponseWriter, status int, message string, data any, errors any) {
	w.Header().Set("Content-Type", "application/json")

	statusString := "success"

	if status >= 400 {
		statusString = "error"
	}

	response := ApiResponse{
		Status:  statusString,
		Message: message,
		Data:    data,
		Errors:  errors,
	}

	json.NewEncoder(w).Encode(response)
}
