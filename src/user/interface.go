package user

import "context"

type Reader interface {
	FindByAll(ctx context.Context) (u []*User, err error)
	FindByID(ctx context.Context, id int64) (u *User, err error)
	FindByEmail(ctx context.Context, email string) (u *User, err error)
}

type Writer interface {
	Insert(ctx context.Context, user User) (u *User, err error)
	Update(ctx context.Context, id int64, user User) (u *User, err error)
	ChangePassword(ctx context.Context, email, password string) (err error)
}

//Repository interface
type URepository interface {
	Reader
	Writer
}
