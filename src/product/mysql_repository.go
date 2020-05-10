package product

import (
	"context"
	"github.com/jinzhu/gorm"
)

type repo struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *repo {
	return &repo{DB: db}
}

func (r *repo) FindAll(ctx context.Context) (p []*Product, err error) {

	result := r.DB.Find(&p)

	switch result.Error {
	case nil:
		return p, nil
	default:
		return nil, result.Error
	}
}

func (r *repo) FindByID(ctx context.Context, id uint) (p *Product, err error) {

	p = &Product{}
	result := r.DB.Where("id = ?", id).First(p)

	switch result.Error {
	case nil:
		return p, nil
	default:
		return nil, result.Error
	}
}

func (r *repo) Store(ctx context.Context, product Product) (p *Product, err error) {

	p = &Product{}
	result := r.DB.Create(&product).Find(&p)

	switch result.Error {
	case nil:
		return p, nil
	default:
		return nil, result.Error
	}
}

func (r *repo) Update(ctx context.Context, id uint, product Product) (p *Product, err error) {

	p = &Product{}
	result := r.DB.Table("products").Where("id = ?", id).First(p).Update(map[string]interface{}{
		"code":        product.Code,
		"name":        product.Name,
		"description": product.Description,
		"status":      product.Status,
	})

	switch result.Error {
	case nil:
		return p, nil
	default:
		return nil, result.Error
	}
}
