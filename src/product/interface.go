package product

import "context"

type Reader interface {
	FindByAll(ctx context.Context, id string) (*[]Product, error)
	FindByID(ctx context.Context, id string) (*Product, error)
}

type Writer interface {
	Insert(ctx context.Context, user Product) (*Product, error)
	Update(ctx context.Context, id string, user Product) error
}

//Repository interface
type PRepository interface {
	Reader
	Writer
}
