package auth

import "go-jwt-auth/models"

type RegisterResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (res *RegisterResponse) FromUserModel(user *models.User) {
	res.Username = user.Username
	res.Email = user.Email
}
