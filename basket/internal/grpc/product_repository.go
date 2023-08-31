package grpc

import (
	"context"

	"github.com/hnamzian/go-mallbots/basket/internal/domain"
	"github.com/hnamzian/go-mallbots/stores/storespb"
	"google.golang.org/grpc"
)

type ProductRepository struct {
	client storespb.StoresClient
}

func NewProductRepository(conn *grpc.ClientConn) ProductRepository {
	return ProductRepository{
		client: storespb.NewStoresClient(conn),
	}
}

func (r ProductRepository) Find(ctx context.Context, id string) (*domain.Product, error) {
	resp, err := r.client.GetProduct(ctx, &storespb.GetProductRequest{
		Id: id,
	})
	if err != nil {
		return nil, err
	}

	return productFromProto(resp.Product), nil
}

func productFromProto(product *storespb.Product) *domain.Product {
	return &domain.Product{
		ID:    product.Id,
		Name:  product.Name,
		Price: product.Price,
	}
}