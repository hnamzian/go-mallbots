package grpc

import (
	"context"

	"github.com/hnamzian/go-mallbots/payments/paymentspb"
	"google.golang.org/grpc"
)

type InvoiceRepository struct {
	client paymentspb.PaymentsServiceClient
}

func NewInvoiceRepository(conn *grpc.ClientConn) *InvoiceRepository {
	return &InvoiceRepository{
		client: paymentspb.NewPaymentsServiceClient(conn),
	}
}

func (r InvoiceRepository) Save(ctx context.Context, orderID, paymentID string, amount float64) error {
	_, err := r.client.CreateInvoice(ctx, &paymentspb.CreateInvoiceRequest{
		OrderId:  orderID,
		PaymentId: paymentID,
		Amount:   amount,
	})
	return err
}

func (r InvoiceRepository) Delete(ctx context.Context, invoiceID string) error {
	_, err := r.client.CancelInvoice(ctx, &paymentspb.CancelInvoiceRequest{
		Id: invoiceID,
	})
	return err
}
