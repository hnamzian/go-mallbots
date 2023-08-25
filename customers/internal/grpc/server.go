package grpc

import (
	"context"

	"github.com/google/uuid"
	"github.com/hnamzian/go-mallbots/customers/customerspb"
	"github.com/hnamzian/go-mallbots/customers/internal/application"
	"google.golang.org/grpc"
)

type Server struct {
	app application.App
	customerspb.UnimplementedCustomersServer
}

func RegisterServer(registrar grpc.ServiceRegistrar, app application.App) {
	customerspb.RegisterCustomersServer(registrar, &Server{app: app})
}

func (s Server) RegisterCustomer(ctx context.Context, request *customerspb.RegisterCustomerRequest) (*customerspb.RegisterCustomerResponse, error) {
	id := uuid.New().String()
	err := s.app.RegisterCustomer(ctx, application.RegisterCustomer{
		ID:        id,
		Name:      request.Name,
		SmsNumber: request.SmsNumber,
	})
	if err != nil {
		return nil, err
	}
	return &customerspb.RegisterCustomerResponse{Id: id}, nil
}

func (s Server) GetCustomer(ctx context.Context, request *customerspb.GetCustomerRequest) (*customerspb.GetCustomerResponse, error) {
	customer, err := s.app.GetCustomer(ctx, application.GetCustomer{
		ID: request.Id,
	})
	if err != nil {
		return nil, err
	}
	return &customerspb.GetCustomerResponse{Customer: &customerspb.Customer{
		Id:        customer.ID,
		Name:      customer.Name,
		SmsNumber: customer.SmsNumber,
		Enabled:   customer.Enabled,
	}}, nil
}

func (s Server) EnableCustomer(ctx context.Context, request *customerspb.EnableCustomerRequest) (*customerspb.EnableCustomerResponse, error) {
	err := s.app.EnableCustomer(ctx, application.EnableCustomer{
		ID: request.Id,
	})
	if err != nil {
		return nil, err
	}
	return &customerspb.EnableCustomerResponse{}, nil
}

func (s Server) DisableCustomer(ctx context.Context, request *customerspb.DisableCustomerRequest) (*customerspb.DisableCustomerResponse, error) {
	err := s.app.DisableCustomer(ctx, application.DisableCustomer{
		ID: request.Id,
	})
	if err != nil {
		return nil, err
	}
	return &customerspb.DisableCustomerResponse{}, nil
}
