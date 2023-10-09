package auth

import (
	"go-jwt-auth/models"
	"html"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

func (r RegisterRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(r.Email, is.Email, validation.Required),
		validation.Field(r.Username, validation.Required),
		validation.Field(r.Password, validation.Required),
	)
}

func (r RegisterRequest) ToUserModel() (*models.User, error) {
	return &models.User{
		Username: html.EscapeString(strings.TrimSpace(r.Username)),
		Password: r.Password,
		Email:    r.Email,
	}, nil
}
