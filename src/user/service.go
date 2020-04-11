package user

import (
	"context"
)

type Service interface {
	FindByAll(ctx context.Context) ([]*User, error)
	FindByID(ctx context.Context, id int64) (*User, error)

	Insert(ctx context.Context, user User) (*User, error)
	Update(ctx context.Context, id int64, user User) (*User, error)
	ChangePassword(ctx context.Context, email, password string) error
}
type service struct {
	repo URepository
}

func NewUserService(r URepository) *service {
	return &service{repo: r}
}

func (s service) FindByAll(ctx context.Context) (u []*User, err error) {

	return s.repo.FindByAll(ctx)
}

func (s service) FindByID(ctx context.Context, id int64) (u *User, err error) {

	return s.repo.FindByID(ctx, id)
}

func (s service) Insert(ctx context.Context, user User) (u *User, err error) {

	return s.repo.Insert(ctx, user)
}

func (s service) Update(ctx context.Context, id int64, user User) (u *User, err error) {

	return s.repo.Update(ctx, id, user)
}

func (s service) ChangePassword(ctx context.Context, email, password string) (err error) {

	return s.repo.ChangePassword(ctx, email, password)
}
