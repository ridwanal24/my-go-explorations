package handler

import (
	"authentication-feature/internal/services"
	"encoding/json"
	"net/http"
	"strconv"
)

type UserHandler struct {
	svc services.UserService
}

func NewUserHandler(s services.UserService) *UserHandler {
	return &UserHandler{svc: s}
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idParam)

	user, err := h.svc.GetUser(id)
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})

		return
	}

	json.NewEncoder(w).Encode(user)

}
