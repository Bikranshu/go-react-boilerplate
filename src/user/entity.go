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
	Password    string    `json:"password,omitempty" gorm:"not null"`
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
