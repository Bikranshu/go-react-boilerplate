package auth

import (
	"../user"
	"../utils"
	"context"
	"errors"
)

type Service interface {
	Login(ctx context.Context, email, password string) (string string, err error)
}
type service struct {
	repo user.URepository
}

func NewAuthService(r user.URepository) Service {
	return &service{repo: r}
}

func (s service) Login(ctx context.Context, email, password string) (string string, err error) {
	res, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	if utils.ValidatePassword(res.Password, password) {
		token, err := utils.GenerateToken(email)
		if err != nil {
			return "", err
		}
		return token, nil
	}
	return "", errors.New("Invalid username or password.")
}
