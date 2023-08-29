package domain

import "context"

type ParticipatingStoreRepository interface {
	GetAll(ctx context.Context) ([]*Store, error)
}
