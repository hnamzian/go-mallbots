package grpc

import (
	"context"

	"github.com/hnamzian/go-mallbots/customers/customerspb"
	"google.golang.org/grpc"
)

type Server struct {
	customerspb.UnimplementedCustomersServer
}

func RegisterServer(registrar grpc.ServiceRegistrar, srv Server) {
	customerspb.RegisterCustomersServer(registrar, srv)
}

func (s Server) RegisterCustomer(ctx context.Context, request *customerspb.RegisterCustomerRequest) (*customerspb.RegisterCustomerResponse, error) {
	return nil, nil
}

func (s Server) GetCustomers(ctx context.Context, request *customerspb.GetCustomerRequest) (*customerspb.GetCustomerResponse, error) {
	return nil, nil
}

func (s Server) EnableCustomer(ctx context.Context, request *customerspb.EnableCustomerRequest) (*customerspb.EnableCustomerResponse, error) {
	return nil, nil
}

func (s Server) DisableCustomer(ctx context.Context, request *customerspb.DisableCustomerRequest) (*customerspb.DisableCustomerResponse, error) {
	return nil, nil
}
