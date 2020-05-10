package user

import (
	"../utils"
	"errors"
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	ID          uint       `json:"id" gorm:"primary_key"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at" sql:"index"`
	FirstName   string     `json:"first_name,omitempty" gorm:"not null"`
	LastName    string     `json:"last_name,omitempty" gorm:"not null"`
	Email       string     `json:"email,omitempty" gorm:"unique;not null"`
	Password    string     `json:"password,omitempty" gorm:"not null"`
	Status      string     `json:"status,omitempty" gorm:"not null"`
	LastLoginAt *time.Time `json:"last_login_at" sql:"DEFAULT:NULL"`
}

func (u *User) BeforeSave(scope *gorm.Scope) {
	if pw, err := utils.EncryptPassword(u.Password); err == nil {
		scope.SetColumn("password", pw)
	}
}

func (u *User) BeforeCreate(scope *gorm.Scope) (err error) {
	if err = scope.DB().Where(&User{Email: u.Email}).First(&u).Error; err == nil {
		return errors.New("record already exists")
	}
	return nil
}

//func (u *User) AfterUpdate(scope *gorm.Scope)  {
//	utils.OmitHiddenFields(u)
//}
//
//func (u *User) AfterFind(scope *gorm.Scope) {
//	utils.OmitHiddenFields(u)
//}

type PublicUser struct {
	ID        uint      `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Users []User

//So that we dont expose the user's password to the world
func (users Users) PublicUsers() []interface{} {
	result := make([]interface{}, len(users))
	for index, user := range users {
		result[index] = user.PublicUser()
	}
	return result
}

//So that we dont expose the user's password to the world
func (u *User) PublicUser() interface{} {
	return &PublicUser{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Status:    u.Status,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
