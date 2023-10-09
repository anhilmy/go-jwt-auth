package auth

import (
	"context"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(context.Context, RegisterRequest) (*RegisterResponse, error)
}

type service struct {
	UserRepo Repository
}

func NewService(userRepo Repository) Service {
	return service{userRepo}
}

func (s service) Register(ctx context.Context, req RegisterRequest) (*RegisterResponse, error) {

	newUser, err := req.ToUserModel()
	if err != nil {
		return nil, err
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	newUser.Password = string(hashed)
	if err := s.UserRepo.Create(ctx, newUser); err != nil {
		return nil, err
	}

	res := RegisterResponse{}
	res.FromUserModel(newUser)

	return &res, nil
}
