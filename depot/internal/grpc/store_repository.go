package grpc

import (
	"context"

	"github.com/hnamzian/go-mallbots/depot/internal/domain"
	"github.com/hnamzian/go-mallbots/stores/storespb"
	"google.golang.org/grpc"
)

type StoreRepository struct {
	client storespb.StoresClient
}

func NewStoreRepository(conn *grpc.ClientConn) *StoreRepository {
	return &StoreRepository{
		client: storespb.NewStoresClient(conn),
	}
}

func (r StoreRepository) Find(ctx context.Context, id string) (*domain.Store, error) {
	reposnse, err := r.client.GetStore(ctx, &storespb.GetStoreRequest{
		Id: id,
	})
	if err != nil {
		return nil, err
	}
	return &domain.Store{
		ID:  reposnse.Store.Id,
		Name: reposnse.Store.Name,
		Location: reposnse.Store.Location,
	}, nil
}
