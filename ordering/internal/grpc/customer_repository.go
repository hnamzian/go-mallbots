package grpc

import (
	"context"

	"github.com/hnamzian/go-mallbots/customers/customerspb"
	"google.golang.org/grpc"
)

type CustomerRepository struct {
	client customerspb.CustomersClient
}

func NewCustomerRepository(conn *grpc.ClientConn) *CustomerRepository {
	return &CustomerRepository{
		client: customerspb.NewCustomersClient(conn),
	}
}

func (r CustomerRepository) Authorize(ctx context.Context, customerID string) error {
	_, err := r.client.AuthorizeCustomer(ctx, &customerspb.AuthorizeCustomerRequest{
		Id: customerID,
	})
	return err
}
