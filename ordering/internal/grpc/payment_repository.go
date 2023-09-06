package grpc

import (
	"context"

	"github.com/hnamzian/go-mallbots/payments/paymentspb"
	"google.golang.org/grpc"
)

type PaymentRepository struct {
	client paymentspb.PaymentsServiceClient
}

func NewPaymentRepository(conn *grpc.ClientConn) *PaymentRepository {
	return &PaymentRepository{
		client: paymentspb.NewPaymentsServiceClient(conn),
	}
}

func (r PaymentRepository) Confirm(ctx context.Context, paymentID string) error {
	_, err := r.client.ConfirmPayment(ctx, &paymentspb.ConfirmPaymentRequest{
		Id: paymentID,
	})
	return err
}
