package routes

import (
	handler "authentication-feature/internal/handler"
	"net/http"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handler.WelcomeHandler)
	mux.HandleFunc("POST /auth/login", handler.LoginHandler)

	return mux
}
