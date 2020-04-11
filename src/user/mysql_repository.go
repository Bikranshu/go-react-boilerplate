package user

import (
	"../pkg"
	"context"
	"github.com/jinzhu/gorm"
)

type repo struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *repo {
	return &repo{DB: db,}
}

func (r *repo) FindByAll(ctx context.Context) (u []*User, err error) {

	result := r.DB.Select("first_name, last_name, email, status").Find(&u)

	switch result.Error {
	case nil:
		return u, nil
	case gorm.ErrRecordNotFound:
		return nil, pkg.ErrNotFound
	default:
		return nil, pkg.ErrDatabase
	}
}

func (r *repo) FindByID(ctx context.Context, id int64) (u *User, err error) {

	u = &User{}
	result := r.DB.Select("first_name, last_name, email, status").Where("id = ?", id).First(u)

	switch result.Error {
	case nil:
		return u, nil
	case gorm.ErrRecordNotFound:
		return nil, pkg.ErrNotFound
	default:
		return nil, pkg.ErrDatabase
	}
}

func (r *repo) FindByEmail(ctx context.Context, email string) (u *User, err error) {

	u = &User{}
	result := r.DB.Where("email = ?", email).First(u)

	switch result.Error {
	case nil:
		return u, nil
	case gorm.ErrRecordNotFound:
		return nil, pkg.ErrNotFound
	default:
		return nil, pkg.ErrDatabase
	}
}

func (r *repo) Insert(ctx context.Context, user User) (u *User, err error) {

	u = &User{}
	result := r.DB.Create(&user)

	switch result.Error {
	case nil:
		return u, nil
	default:
		return nil, pkg.ErrDatabase
	}
}

func (r *repo) Update(ctx context.Context, id int64, user User) (u *User, err error) {

	result := r.DB.Table("users").Where("id = ?", id).Update(map[string]interface{}{
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"email":      user.Email,
		"status":     user.Status,
	})

	switch result.Error {
	case nil:
		return u, nil
	case gorm.ErrRecordNotFound:
		return nil, pkg.ErrNotFound
	default:
		return nil, pkg.ErrDatabase
	}
}

func (r *repo) ChangePassword(ctx context.Context, email, password string) (err error) {

	result := r.DB.Table("users").Where("email = ?", email).Update(map[string]interface{}{
		"password": password,
	})

	switch result.Error {
	case nil:
		return nil
	case gorm.ErrRecordNotFound:
		return pkg.ErrNotFound
	default:
		return pkg.ErrDatabase
	}
}
