package domain

import "context"

type BasketRepository interface {
	Save(ctx context.Context, basket *Basket) error
	Get(ctx context.Context, id string) (*Basket, error)
	GetAll(ctx context.Context, customerID string) ([]*Basket, error)
	Update(ctx context.Context, basket *Basket) (error)
}
