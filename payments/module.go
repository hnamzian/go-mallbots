package payments

import (
	"context"

	"github.com/hnamzian/go-mallbots/payments/internal/application"
	"github.com/hnamzian/go-mallbots/payments/internal/repository"
	"github.com/hnamzian/go-mallbots/payments/internal/logger"
	"github.com/hnamzian/go-mallbots/payments/internal/grpc"
	"github.com/hnamzian/go-mallbots/payments/internal/rest"
	"github.com/hnamzian/go-mallbots/internal/module"
)

type Module struct {}

func (m Module) Startup(ctx context.Context, core module.Core) error {
	invoices := repository.NewInvoiceRepository(core.DB())
	payments := repository.NewPaymentRepository(core.DB())

	conn, err := grpc.Dial(ctx, core.Config().Grpc.Address())
	if err != nil {
		return err
	}
	orders := grpc.NewOrderRepository(conn)
	
	var app application.App
	app = application.NewApplication(invoices, payments, orders)
	app = logger.NewApplication(core.Logger(), app)

	grpc.RegisterServer(core.RPC(), app)

	if err = rest.RegisterGateway(ctx, core.Mux(), core.Config().Grpc.Address()); err != nil {
		return err
	}
	if err = rest.RegisterSwagger(core.Mux()); err != nil {
		return err
	}
	
	return nil
}