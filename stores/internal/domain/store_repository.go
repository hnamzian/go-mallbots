package domain

import "context"

type StoreRepository interface {
	Save(ctx context.Context, store *Store) error
	GetAll(ctx context.Context) ([]*Store, error)
	Get(ctx context.Context, id string) (*Store, error)
	Update(ctx context.Context, store *Store) error
	Delete(ctx context.Context, id string) error
}
