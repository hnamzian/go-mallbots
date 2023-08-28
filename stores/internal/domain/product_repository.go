package domain

import "context"

type ProductRepository interface {
	Save(ctx context.Context, product *Product) error
	Get(ctx context.Context, id string) (*Product, error)
	GetCatalog(ctx context.Context, storeID string) ([]*Product, error)
	Delete(ctx context.Context, id string) error
}
