package depot

import (
	"context"

	"github.com/hnamzian/go-mallbots/depot/internal/application"
	"github.com/hnamzian/go-mallbots/depot/internal/grpc"
	"github.com/hnamzian/go-mallbots/depot/internal/handlers"
	"github.com/hnamzian/go-mallbots/depot/internal/logger"
	"github.com/hnamzian/go-mallbots/depot/internal/repository"
	"github.com/hnamzian/go-mallbots/depot/internal/rest"
	"github.com/hnamzian/go-mallbots/internal/ddd"
	"github.com/hnamzian/go-mallbots/internal/module"
)

type Module struct{}

func (m Module) Startup(ctx context.Context, core module.Core) error {
	domainDispatcher := ddd.NewEventDispatcher()

	conn, err := grpc.Dial(ctx, core.Config().Grpc.Address())
	if err != nil {
		return err
	}

	products := grpc.NewProductRepository(conn)
	stores := grpc.NewStoreRepository(conn)
	shoppingLists := repository.NewShoppingListRepository(core.DB())
	orders := grpc.NewOrderRepository(conn)

	var app application.App
	app = application.NewDepotApplication(shoppingLists, orders, products, stores)
	app = logger.NewApplication(core.Logger(), app)

	grpc.RegisterServer(core.RPC(), app)
	if err := rest.RegisterGateway(ctx, core.Mux(), core.Config().Grpc.Address()); err != nil {
		return err
	}
	if err := rest.RegisterSwagger(core.Mux()); err != nil {
		return err
	}

	orderHandlers := logger.NewDomainEventHandlers(application.NewOrderEventHandlers(orders), core.Logger())
	handlers.RegisterOrderHandlers(orderHandlers, domainDispatcher)

	return nil
}
