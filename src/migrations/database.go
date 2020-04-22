package migrations

import (
	"../product"
	"../user"
	"github.com/jinzhu/gorm"
)

func InitMigration(db *gorm.DB) {
	db.AutoMigrate(&user.User{})
	db.AutoMigrate(&product.Product{}).AddForeignKey("created_by", "users(id)", "RESTRICT", "RESTRICT")
}
