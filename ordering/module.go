package ordering

import (
	"context"

	"github.com/hnamzian/go-mallbots/internal/ddd"
	"github.com/hnamzian/go-mallbots/internal/module"
	"github.com/hnamzian/go-mallbots/ordering/internal/application"
	"github.com/hnamzian/go-mallbots/ordering/internal/grpc"
	"github.com/hnamzian/go-mallbots/ordering/internal/handlers"
	"github.com/hnamzian/go-mallbots/ordering/internal/logger"
	"github.com/hnamzian/go-mallbots/ordering/internal/repository"
	"github.com/hnamzian/go-mallbots/ordering/internal/rest"
)

type Module struct{}

func (m Module) Startup(ctx context.Context, core module.Core) error {
	domainDispatcher := ddd.NewEventDispatcher()

	conn, err := grpc.Dial(ctx, core.Config().Grpc.Address())
	if err != nil {
		return err
	}
	orders := repository.NewOrderRepository(core.DB())
	customers := grpc.NewCustomerRepository(conn)
	payments := grpc.NewPaymentRepository(conn)
	shoppings := grpc.NewShoppingRepository(conn)
	invoices := grpc.NewInvoiceRepository(conn)
	notifications := grpc.NewNotificationeRepository(conn)

	var app application.App
	app = application.NewOrderingApplication(
		orders,
		customers,
		shoppings,
		payments,
		domainDispatcher,
	)
	app = logger.NewApplication(app, core.Logger())

	grpc.RegisterServer(app)
	if err = rest.Registergateway(ctx, core.Mux(), core.Config().Grpc.Address()); err != nil {
		return err
	}
	if rest.RegisterSwagger(core.Mux()); err != nil {
		return err
	}

	invoicehandlers := logger.NewDomainEventHandlers(application.NewInvoiceHandlers(invoices), core.Logger())
	notificationHandlers := logger.NewDomainEventHandlers(application.NewNotificationHandlers(notifications), core.Logger())

	handlers.RegisterInvoiceHandlers(invoicehandlers, domainDispatcher)
	handlers.RegisterNotificationHandlers(notificationHandlers, domainDispatcher)

	return nil
}
