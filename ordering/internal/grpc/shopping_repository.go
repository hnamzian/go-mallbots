package grpc

import (
	"context"

	"github.com/hnamzian/go-mallbots/depot/depotpb"
	"github.com/hnamzian/go-mallbots/ordering/internal/domain"
	"google.golang.org/grpc"
)

type ShoppingRepository struct {
	client depotpb.DepotServiceClient
}

func NewShoppingRepository(conn *grpc.ClientConn) *ShoppingRepository {
	return &ShoppingRepository{
		client: depotpb.NewDepotServiceClient(conn),
	}
}

func (r ShoppingRepository) Create(ctx context.Context, order *domain.Order) (string, error) {
	var items []*depotpb.OrderItem
	for _, item := range order.Items {
		items = append(items, &depotpb.OrderItem{
			ProductId: item.ProductID,
			StoreId:   item.StoreID,
			Quantity:  int32(item.Quantity),
		})
	}
	resp, err := r.client.CreateShoppingList(ctx, &depotpb.CreateShoppingListRequest{
		OrderId: order.ID,
		Items:   items,
	})

	return resp.GetId(), err
}

func (r ShoppingRepository) Cancel(ctx context.Context, shoppingID string) error {
	_, err := r.client.CancelShoppingList(ctx, &depotpb.CancelShoppingListRequest{
		Id: shoppingID,
	})
	return err
}
