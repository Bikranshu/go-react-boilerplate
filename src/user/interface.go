package user

import "context"

type Reader interface {
	FindAll(ctx context.Context) (u *Users, err error)
	FindByID(ctx context.Context, id uint) (u *User, err error)
	FindByEmail(ctx context.Context, email string) (u *User, err error)
}

type Writer interface {
	Store(ctx context.Context, user User) (u *User, err error)
	Update(ctx context.Context, id uint, user User) (u *User, err error)
	ChangePassword(ctx context.Context, id uint, email, password string) (err error)
}

//Repository interface
type URepository interface {
	Reader
	Writer
}
