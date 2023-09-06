package ordering

import (
	"context"

	"github.com/hnamzian/go-mallbots/internal/module"
	"github.com/hnamzian/go-mallbots/ordering/internal/application"
	"github.com/hnamzian/go-mallbots/ordering/internal/grpc"
	"github.com/hnamzian/go-mallbots/ordering/internal/rest"
	"github.com/hnamzian/go-mallbots/ordering/internal/logger"
	"github.com/hnamzian/go-mallbots/ordering/internal/repository"
)

type Module struct {}

func (m Module) Startup(ctx context.Context, core module.Core) error {
	conn, err := grpc.Dial(ctx, core.Config().Grpc.Address())
	if err != nil {
		return err
	}
	orders := repository.NewOrderRepository(core.DB())
	customers := grpc.NewCustomerRepository(conn)
	invoices := grpc.NewInvoiceRepository(conn)
	notifications := grpc.NewNotificationeRepository(conn)
	payments := grpc.NewPaymentRepository(conn)
	shoppings := grpc.NewShoppingRepository(conn)

	var app application.App
	app = application.NewOrderingApplication(
		orders,
		customers,
		invoices,
		shoppings,
		payments,
		notifications,
	)
	app = logger.NewApplication(app, core.Logger())

	grpc.RegisterServer(app)
	if err = rest.Registergateway(ctx, core.Mux(), core.Config().Grpc.Address()); err != nil {
		return err
	}
	if rest.RegisterSwagger(core.Mux()); err != nil {
		return err
	}

	return nil
}