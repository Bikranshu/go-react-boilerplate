package product
import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Code   string `json:"code,omitempty" gorm:"not null"`
	Name    string `json:"name,omitempty" gorm:"not null"`
	Description       string `json:"description,omitempty"`
	Status  string `json:"status,omitempty" gorm:"not null"`
	CreatedBy  uint `json:"created_by,omitempty"`
}