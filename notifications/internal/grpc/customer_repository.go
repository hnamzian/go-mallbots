package grpc

import (
	"context"

	"github.com/hnamzian/go-mallbots/customers/customerspb"
	"github.com/hnamzian/go-mallbots/notifications/internal/models"
	"google.golang.org/grpc"
)

type CustomerRepository struct {
	client customerspb.CustomersClient
}

func NewCustomerRepository(conn *grpc.ClientConn) CustomerRepository {
	return CustomerRepository{
		client: customerspb.NewCustomersClient(conn),
	}
}

func (r CustomerRepository) Find(ctx context.Context, customerID string) (*models.Customer, error) {
	resp, err := r.client.GetCustomer(ctx, &customerspb.GetCustomerRequest{
		Id: customerID,
	})
	if err != nil {
		return nil, err
	}

	return customerFromProto(resp.Customer), nil
}

func customerFromProto(customerProto *customerspb.Customer) *models.Customer {
	return &models.Customer{
		ID:        customerProto.Id,
		Name:      customerProto.Name,
		SmsNumber: customerProto.SmsNumber,
	}
}
