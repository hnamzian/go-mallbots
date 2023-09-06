package grpc

import (
	"context"

	"google.golang.org/grpc"
	"github.com/hnamzian/go-mallbots/oerdering/orderingpb"
)

type OrderRepository struct {
	client orderingpb.OrderingServiceClient
}

func NewOrderRepository(conn *grpc.ClientConn) *OrderRepository {
	return &OrderRepository{
		client: orderingpb.NewOrderingServiceClient(conn),
	}
}

func (r OrderRepository) Ready(ctx context.Context, orderID string) error {
	_, err := r.client.ReadyOrder(ctx, &orderingpb.ReadyOrderRequest{
		Id: orderID,
	})
	return err
}
