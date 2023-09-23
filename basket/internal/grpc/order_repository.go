package grpc

import (
	"context"

	"github.com/hnamzian/go-mallbots/basket/internal/domain"
	"github.com/hnamzian/go-mallbots/ordering/orderingpb"
	"github.com/stackus/errors"
	"google.golang.org/grpc"
)

type OrderRepository struct {
	client orderingpb.OrderingServiceClient
}

func NewOrderRepository(conn *grpc.ClientConn) *OrderRepository {
	return &OrderRepository{
		client: orderingpb.NewOrderingServiceClient(conn),
	}
}

func (r *OrderRepository) Save(ctx context.Context, basket *domain.Basket) (string, error) {
	items := make([]*orderingpb.Item, len(basket.Items))
	for i, item := range basket.Items {
		items[i] = &orderingpb.Item{
			ProductId:   item.ProductID,
			StoreId:     item.StoreID,
			ProductName: item.ProductName,
			StoreName:   item.StoreName,
			Price:       item.ProductPrice,
			Quantity:    int32(item.Quantity),
		}
	}
	resp, err := r.client.CreateOrder(ctx, &orderingpb.CreateOrderRequest{
		Items:      items,
		CustomerId: basket.CustomerID,
		PaymentId:  basket.PaymentID,
	})
	if err != nil {
		return "", errors.Wrap(err, "failed to create order")
	}
	return resp.GetId(), nil
}
