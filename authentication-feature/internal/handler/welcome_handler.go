package handler

import (
	"authentication-feature/internal/helpers"
	"fmt"
	"net/http"
)

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("Welcome to isekai %s", "bang")
	status := 200

	helpers.WriteApiResponse(w, status, message, nil, nil)
}
