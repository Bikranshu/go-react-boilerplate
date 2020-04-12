package user

import (
	"../utils"
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	FirstName   string    `json:"first_name,omitempty" gorm:"not null"`
	LastName    string    `json:"last_name,omitempty" gorm:"not null"`
	Email       string    `json:"email,omitempty" gorm:"unique;not null"`
	Password    string    `json:"password,omitempty" gorm:"not null"`
	Status      string    `json:"status,omitempty" gorm:"not null"`
	LastLoginAt time.Time `json:"last_login_at,omitempty" sql:"DEFAULT:NULL"`
}

func (user *User) BeforeSave(scope *gorm.Scope) {
	if pw, err := utils.EncryptPassword(user.Password); err == nil {
		scope.SetColumn("password", pw)
	}
}

//func (user *User) BeforeCreate(scope *gorm.Scope) (err error) {
//	if err = scope.DB().Where(&User{Email: user.Email}).First(&user).Error; err == nil {
//		return errors.New("User already exists in system.")
//	}
//	return nil
//}
