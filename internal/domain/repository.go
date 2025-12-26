package domain

import "context"

type ProductRepository interface {
	Save(product *Product, ctx context.Context) error
	FindAll(ctx context.Context) ([]Product, error)
	FindByID(id string, ctx context.Context) (*Product, error)
	UpdateStock(id string, stockQuantity int, ctx context.Context) (*Product, error)
}

type OrderRepository interface {
	Save(order *Order, ctx context.Context) error
	FindAll(ctx context.Context) ([]Order, error)
	FindByID(id string, ctx context.Context) (*Order, error)
}
