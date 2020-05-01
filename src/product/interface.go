package product

import "context"

type Reader interface {
	FindAll(ctx context.Context, id string) (*[]Product, error)
	FindByID(ctx context.Context, id string) (*Product, error)
}

type Writer interface {
	Store(ctx context.Context, user Product) (*Product, error)
	Update(ctx context.Context, id string, user Product) error
}

//Repository interface
type PRepository interface {
	Reader
	Writer
}
