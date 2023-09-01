package domain

import (
	"context"

	"github.com/hnamzian/go-mallbots/notifications/internal/models"

)

type CustomerRepository interface {
	Find(ctx context.Context, customerID string) (*models.Customer, error)
}
