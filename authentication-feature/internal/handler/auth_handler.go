package handler

import (
	"authentication-feature/internal/dto"
	"authentication-feature/internal/helpers"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

// isisialisasi validator nya
var validate = validator.New()

func LoginHandler(w http.ResponseWriter, r *http.Request) {
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

	// return response
	helpers.WriteApiResponse(w, 200, "login success", nil, nil)

}
