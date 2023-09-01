package grpc

import (
	"context"

	"github.com/google/uuid"
	"github.com/hnamzian/go-mallbots/payments/internal/application"
	"github.com/hnamzian/go-mallbots/payments/internal/domain"
	"github.com/hnamzian/go-mallbots/payments/paymentspb"
	"google.golang.org/grpc"
)

type PaymentServer struct {
	application.App
	paymentspb.UnsafePaymentsServiceServer
}

func RegisterServer(registrar grpc.ServiceRegistrar, app application.App) {
	paymentspb.RegisterPaymentsServiceServer(registrar, PaymentServer{App: app})
}

func (s PaymentServer) AuthorizePayment(ctx context.Context, request *paymentspb.AuthorizePaymentRequest) (*paymentspb.AuthorizePaymentResponse, error) {
	id := uuid.New().String()
	err := s.App.AuthorizePayment(ctx, application.AuthorizePayment{
		ID:         id,
		CustomerID: request.CustomerId,
		Amount:     request.Amount,
	})
	if err != nil {
		return nil, err
	}
	return &paymentspb.AuthorizePaymentResponse{Id: id}, nil
}
func (s PaymentServer) ConfirmPayment(ctx context.Context, request *paymentspb.ConfirmPaymentRequest) (*paymentspb.ConfirmPaymentResponse, error) {
	err := s.App.ConfirmPayment(ctx, application.ConfirmPayment{
		ID: request.Id,
	})
	if err != nil {
		return nil, err
	}
	return &paymentspb.ConfirmPaymentResponse{}, nil
}
func (s PaymentServer) GetPayment(ctx context.Context, request *paymentspb.GetPaymentRequest) (*paymentspb.GetPaymentResponse, error) {
	payment, err := s.App.GetPayment(ctx, application.GetPayment{ID: request.Id})
	if err != nil {
		return nil, err
	}
	return &paymentspb.GetPaymentResponse{
		Payment: paymentFromDomain(payment),
	}, nil
}
func (s PaymentServer) CreateInvoice(ctx context.Context, request *paymentspb.CreateInvoiceRequest) (*paymentspb.CreateInvoiceResponse, error) {
	id := uuid.New().String()
	err := s.App.CreateInvoice(ctx, application.CreateInvoice{
		ID:      id,
		OrderID: request.OrderId,
		Amount:  request.Amount,
	})
	if err != nil {
		return nil, err
	}
	return &paymentspb.CreateInvoiceResponse{Id: id}, nil
}
func (s PaymentServer) AdjustInvoice(ctx context.Context, request *paymentspb.AdjustInvoiceRequest) (*paymentspb.AdjustInvoiceResponse, error) {
	err := s.App.AdjustInvoice(ctx, application.AdjustInvoice{
		ID:     request.Id,
		Amount: request.Amount,
	})
	return &paymentspb.AdjustInvoiceResponse{}, err
}
func (s PaymentServer) PayInvoice(ctx context.Context, request *paymentspb.PayInvoiceRequest) (*paymentspb.PayInvoiceResponse, error) {
	err := s.App.PayInvoice(ctx, application.PayInvoice{
		ID: request.Id,
	})
	return &paymentspb.PayInvoiceResponse{}, err
}
func (s PaymentServer) CancelInvoice(ctx context.Context, request *paymentspb.CancelInvoiceRequest) (*paymentspb.CancelInvoiceResponse, error) {
	err := s.App.CancelInvoice(ctx, application.CancelInvoice{
		ID: request.Id,
	})
	return &paymentspb.CancelInvoiceResponse{}, err
}
func (s PaymentServer) GetInvoice(ctx context.Context, request *paymentspb.GetInvoiceRequest) (*paymentspb.GetInvoiceResponse, error) {
	invoice, err := s.App.GetInvoice(ctx, application.GetInvoice{ID: request.Id})
	if err != nil {
		return nil, err
	}
	return &paymentspb.GetInvoiceResponse{
		Invoice: invoiceFromDomain(invoice),
	}, nil
}

func paymentFromDomain(payment *domain.Payment) *paymentspb.Payment {
	return &paymentspb.Payment{
		Id:         payment.ID,
		CustomerId: payment.CustomerID,
		Amount:     payment.Amount,
	}
}

func invoiceFromDomain(invoice *domain.Invoice) *paymentspb.Invoice {
	return &paymentspb.Invoice{
		Id:      invoice.ID,
		OrderId: invoice.OrderID,
		Amount:  invoice.Amount,
		Status:  invoice.Status.String(),
	}
}