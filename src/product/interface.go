package product

import "context"

type Reader interface {
	FindAll(ctx context.Context) (p []*Product, err error)
	FindByID(ctx context.Context, id uint) (p *Product, err error)
}

type Writer interface {
	Store(ctx context.Context, product Product) (p *Product, err error)
	Update(ctx context.Context, id uint, product Product) (p *Product, err error)
}

//Repository interface
type PRepository interface {
	Reader
	Writer
}
