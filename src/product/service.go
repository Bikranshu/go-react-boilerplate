package product

import (
	"context"
)

type Service interface {
	FindAll(ctx context.Context) ([]*Product, error)
	FindByID(ctx context.Context, id uint) (*Product, error)

	Store(ctx context.Context, product Product) (*Product, error)
	Update(ctx context.Context, id uint, product Product) (*Product, error)
}
type service struct {
	repo PRepository
}

func NewProductService(r PRepository) *service {
	return &service{repo: r}
}

func (s service) FindAll(ctx context.Context) (p []*Product, err error) {

	return s.repo.FindAll(ctx)
}

func (s service) FindByID(ctx context.Context, id uint) (p *Product, err error) {

	return s.repo.FindByID(ctx, id)
}

func (s service) Store(ctx context.Context, product Product) (p *Product, err error) {

	return s.repo.Store(ctx, product)
}

func (s service) Update(ctx context.Context, id uint, product Product) (p *Product, err error) {

	return s.repo.Update(ctx, id, product)
}
