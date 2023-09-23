package domain

import "context"

type OrderRepository interface {
	Save(context.Context, *Basket) (string, error)
}
