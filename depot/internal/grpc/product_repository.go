package grpc

import (
	"context"

	"google.golang.org/grpc"

	"github.com/hnamzian/go-mallbots/depot/internal/domain"
	"github.com/hnamzian/go-mallbots/stores/storespb"
)

type ProductRepository struct {
	client storespb.StoresClient
}

func NewProductRepository(conn *grpc.ClientConn) *ProductRepository {
	return &ProductRepository{
		client: storespb.NewStoresClient(conn),
	}
}

func (r *ProductRepository) Find(ctx context.Context, productID string) (*domain.Product, error) {
	response, err := r.client.GetProduct(ctx, &storespb.GetProductRequest{
		Id: productID,	
	})
	if err != nil {
		return nil, err
	}
	return &domain.Product{
		ID: response.Product.Id,
		StoreID: response.Product.StoreId,
		Name: response.Product.Name,
	}, nil
}
