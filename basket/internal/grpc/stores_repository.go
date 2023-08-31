package grpc

import (
	"context"

	"github.com/hnamzian/go-mallbots/basket/internal/domain"
	"github.com/hnamzian/go-mallbots/stores/storespb"
	"google.golang.org/grpc"
)

type StoreRepository struct {
	client storespb.StoresClient
}

func NewStoreRepository(conn *grpc.ClientConn) StoreRepository {
	return StoreRepository{
		client: storespb.NewStoresClient(conn),
	}
}

func (r StoreRepository) Find(ctx context.Context, id string) (*domain.Store, error) {
	resp, err := r.client.GetStore(ctx, &storespb.GetStoreRequest{
		Id: id,
	})
	if err != nil {
		return nil, err
	}

	return storeFromProto(resp.Store), nil
}

func storeFromProto(store *storespb.Store) *domain.Store {
	return &domain.Store{
		ID:       store.Id,
		Name:     store.Name,
		Location: store.Location,
	}
}
