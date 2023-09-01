package domain

import "context"

type PaymentRepository interface {
	Save(ctx context.Context, payment *Payment) error
	Update(ctx context.Context, payment *Payment) error
	Find(ctx context.Context, id string) (*Payment, error)
}
