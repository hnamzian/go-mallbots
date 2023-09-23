package baskets

import (
	"context"

	"github.com/hnamzian/go-mallbots/basket/internal/application"
	"github.com/hnamzian/go-mallbots/basket/internal/grpc"
	"github.com/hnamzian/go-mallbots/basket/internal/handlers"
	"github.com/hnamzian/go-mallbots/basket/internal/logger"
	"github.com/hnamzian/go-mallbots/basket/internal/repository"
	"github.com/hnamzian/go-mallbots/basket/internal/rest"
	"github.com/hnamzian/go-mallbots/internal/ddd"
	"github.com/hnamzian/go-mallbots/internal/module"
)

type Module struct{}

func (m Module) Startup(ctx context.Context, core module.Core) error {
	domainEventDispatcher := ddd.NewEventDispatcher()

	conn, err := grpc.Dial(ctx, core.Config().Grpc.Address())
	if err != nil {
		return err
	}

	baskets := repository.NewBasketRepository("baskets.baskets", core.DB())
	products := grpc.NewProductRepository(conn)
	stores := grpc.NewStoreRepository(conn)
	orders := grpc.NewOrderRepository(conn)

	var app application.App
	app = application.NewApplication(baskets, products, stores)
	app = logger.NewApplication(core.Logger(), app)

	grpc.RegisterServer(core.RPC(), app)
	if err = rest.RegisterGateway(ctx, core.Mux(), core.Config().Grpc.Address()); err != nil {
		return err
	}
	if err = rest.RegisterSwagger(core.Mux()); err != nil {
		return err
	}

	ordersEventHandlers := logger.NewDomainEventHandlers(application.NewOrderHandlers(orders), core.Logger())
	handlers.RegisterOrderHandlers(ordersEventHandlers, domainEventDispatcher)

	return nil
}
