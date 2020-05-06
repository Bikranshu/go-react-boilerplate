package user

import (
	"context"
	"github.com/jinzhu/gorm"
)

type repo struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *repo {
	return &repo{DB: db}
}

func (r *repo) FindAll(ctx context.Context) (u []*User, err error) {

	result := r.DB.Find(&u)

	switch result.Error {
	case nil:
		return u, nil
	default:
		return nil, result.Error
	}
}

func (r *repo) FindByID(ctx context.Context, id uint) (u *User, err error) {

	u = &User{}
	result := r.DB.Where("id = ?", id).First(u)

	switch result.Error {
	case nil:
		return u, nil
	default:
		return nil, result.Error
	}
}

func (r *repo) FindByEmail(ctx context.Context, email string) (u *User, err error) {

	u = &User{}
	result := r.DB.Where("email = ?", email).First(u)

	switch result.Error {
	case nil:
		return u, nil
	default:
		return nil, result.Error
	}
}

func (r *repo) Store(ctx context.Context, user User) (u *User, err error) {

	u = &User{}
	result := r.DB.Create(&user).Find(&u)

	switch result.Error {
	case nil:
		return u, nil
	default:
		return nil, result.Error
	}
}

func (r *repo) Update(ctx context.Context, id uint, user User) (u *User, err error) {

	u = &User{}
	result := r.DB.Table("users").Where("id = ?", id).First(u).Update(map[string]interface{}{
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"email":      user.Email,
		"status":     user.Status,
	})

	switch result.Error {
	case nil:
		return u, nil
	default:
		return nil, result.Error
	}
}

func (r *repo) ChangePassword(ctx context.Context, id uint, email, password string) (err error) {

	u := &User{}
	result := r.DB.Model(u).Where("id = ? AND email = ?", id, email).First(u).UpdateColumn("password", password)

	switch result.Error {
	case nil:
		return nil
	default:
		return result.Error
	}
}
