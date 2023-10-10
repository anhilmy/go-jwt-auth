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

type LoginResponse struct {
	JWT string `json:"jwt"`
}

type ProfileResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (res *ProfileResponse) FromUserModel(user *models.User) {
	res.Username = user.Username
	res.Email = user.Email
	res.ID = user.ID
}
