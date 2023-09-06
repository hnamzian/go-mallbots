package grpc

import (
	"context"

	"github.com/hnamzian/go-mallbots/ordering/orderingpb"
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

func (r OrderRepository) Complete(ctx context.Context, invoiceId, orderId string) error {
	_, err := r.client.CompletedOrder(ctx, &orderingpb.CompletedOrderRequest{
		Id:        orderId,
		InvoiceId: invoiceId,
	})
	return err
}
