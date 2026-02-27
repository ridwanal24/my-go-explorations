package handler

import (
	"authentication-feature/internal/dto"
	"authentication-feature/internal/helpers"
	"authentication-feature/internal/services"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type authHandler struct {
	authSvc services.AuthService
}

func NewAuthHandler(s services.AuthService) *authHandler {
	return &authHandler{authSvc: s}
}

// isisialisasi validator nya
var validate = validator.New()

func (authHandler *authHandler) Login(w http.ResponseWriter, r *http.Request) {
	// ubah json body ke struct
	var request dto.LoginRequestDto

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		helpers.WriteApiResponse(w, 400, "Invalid JSON", nil, nil)
		return
	}

	// validasi
	err = validate.Struct(request)

	if err != nil {
		errorList := helpers.FormatValidationError(err.(validator.ValidationErrors))
		helpers.WriteApiResponse(w, 400, "validasi gagal", nil, errorList)

		return
	}

	// ambil ke repo
	isMatch := authHandler.authSvc.DoLogin(request.Email, request.Password)

	if err != nil {
		helpers.WriteApiResponse(w, 400, "DB bermasalah", nil, nil)
		return
	}

	// return response
	if !isMatch {
		helpers.WriteApiResponse(w, http.StatusUnauthorized, "Email/Password salah", nil, nil)

		return
	}

	helpers.WriteApiResponse(w, 200, "login success", nil, nil)

}
