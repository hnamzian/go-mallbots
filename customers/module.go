package customers

import (
	"context"

	"github.com/hnamzian/go-mallbots/customers/internal/application"
	"github.com/hnamzian/go-mallbots/customers/internal/grpc"
	"github.com/hnamzian/go-mallbots/customers/internal/repository"
	"github.com/hnamzian/go-mallbots/internal/module"
)

type Module struct {}

func (m Module) Startup(ctx context.Context, app module.Core) error{
	customers := repository.NewCustomersRepository("customers.customers", app.DB())

	customersApp := application.NewApplication(customers)

	grpc.RegisterServer(app.RPC(), customersApp)

	return nil
}
