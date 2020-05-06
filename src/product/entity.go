package product

import (
	"time"
)

type Product struct {
	ID          uint       `json:"id" gorm:"primary_key"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at" sql:"index"`
	Code        string     `json:"code,omitempty" gorm:"not null"`
	Name        string     `json:"name,omitempty" gorm:"not null"`
	Description string     `json:"description,omitempty"`
	Status      string     `json:"status,omitempty" gorm:"not null"`
	CreatedBy   uint       `json:"created_by,omitempty"`
}
