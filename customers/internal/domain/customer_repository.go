package domain

import "context"

type CustomersRepository interface {
	Save(ctx context.Context, customer *Customer) error
	Get(ctx context.Context, customerID string) (*Customer, error)
	Update(ctx context.Context, customer *Customer) error
}
