package customers

import (
	"context"

	"github.com/hnamzian/go-mallbots/customers/internal/application"
	"github.com/hnamzian/go-mallbots/customers/internal/grpc"
	"github.com/hnamzian/go-mallbots/customers/internal/logger"
	"github.com/hnamzian/go-mallbots/customers/internal/repository"
	"github.com/hnamzian/go-mallbots/customers/internal/rest"
	"github.com/hnamzian/go-mallbots/internal/module"
)

type Module struct{}

func (m Module) Startup(ctx context.Context, core module.Core) error {
	customers := repository.NewCustomersRepository("customers.customers", core.DB())

	var app application.App
	app = application.NewApplication(customers)
	app = logger.NewApplication(app, core.Logger())

	grpc.RegisterServer(core.RPC(), app)

	if err := rest.RegisterGateway(ctx, core.Mux(), core.Config().Grpc.Address()); err != nil {
		return err
	}
	if err := rest.RegisterSwagger(core.Mux()); err != nil {
		return err
	}

	return nil
}
