package dto

type LoginRequestDto struct {
	Email    string `json:"email" validate:"required,email,min=3,max=50"`
	Password string `json:"password" validate:"required,min=6"`
}
