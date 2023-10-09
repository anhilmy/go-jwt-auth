package auth

import "context"

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

	if err := s.UserRepo.Create(ctx, newUser); err != nil {
		return nil, err
	}

	res := RegisterResponse{}
	res.FromUserModel(newUser)

	return &res, nil
}
