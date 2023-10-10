package auth

import (
	"context"
	"go-jwt-auth/internal/errors"
	"go-jwt-auth/internal/token"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(context.Context, RegisterRequest) (*RegisterResponse, error)
	Login(context.Context, LoginRequest) (*LoginResponse, error)
	Profile(context.Context, string) (*ProfileResponse, error)
}

type service struct {
	UserRepo Repository

	tokenServ token.Service
}

func NewService(userRepo Repository, tokenServ token.Service) Service {
	return service{userRepo, tokenServ}
}

func (s service) Register(ctx context.Context, req RegisterRequest) (*RegisterResponse, error) {

	newUser, err := req.ToUserModel()
	if err != nil {
		return nil, err
	}

	if err := s.UserRepo.Create(ctx, newUser); err != nil {
		return nil, err
	}

	res := RegisterResponse{}
	res.FromUserModel(newUser)

	return &res, nil
}

func (s service) Login(ctx context.Context, req LoginRequest) (*LoginResponse, error) {
	user, err := s.UserRepo.GetByUsername(ctx, req.Username)
	if user == nil {
		return nil, errors.LoginFailed{}
	} else if err != nil {
		return nil, err
	}
	err = user.VerifyPassword(req.Password)
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return nil, errors.LoginFailed{}
	} else if err != nil {
		return nil, err
	}

	token, err := s.tokenServ.GenerateToken(user.ID)
	if err != nil {
		return nil, err
	}

	res := LoginResponse{
		JWT: token,
	}

	return &res, nil
}

func (s service) Profile(ctx context.Context, token string) (*ProfileResponse, error) {

	userId, err := s.tokenServ.TokenValid(token)
	if err != nil {
		return nil, err
	}

	user, err := s.UserRepo.GetById(ctx, userId)
	if err != nil {
		return nil, err
	} else if user == nil {
		return nil, errors.TokenInvalid{}
	}

	res := ProfileResponse{}
	res.FromUserModel(user)

	return &res, nil
}
