package stores

import (
	"context"

	"github.com/hnamzian/go-mallbots/internal/module"
	"github.com/hnamzian/go-mallbots/stores/internal/application"
	"github.com/hnamzian/go-mallbots/stores/internal/grpc"
	"github.com/hnamzian/go-mallbots/stores/internal/rest"
	"github.com/hnamzian/go-mallbots/stores/internal/repository"
	"github.com/hnamzian/go-mallbots/stores/internal/logger"
)

type Module struct{}

func (m Module) Startup(ctx context.Context, core module.Core) error {
	var app application.App
	storeRepository := repository.NewStoreRepository("stores.stores", core.DB())
	participatingStoreRepository := repository.NewParticipatingStoreRepository("stores.stores", core.DB())
	productRepository := repository.NewProductRepository("stores.products", core.DB())
	app = application.NewApplication(storeRepository, productRepository, participatingStoreRepository)
	
	app = logger.NewApplication(app, core.Logger())

	grpc.RegisterServer(core.RPC(), app)
	if err := rest.RegisterGateway(ctx, core.Mux(), core.Config().Grpc.Address()); err!= nil {
		return err
	}
	if err := rest.RegisterSwagger(core.Mux()); err!= nil {
		return err
	}

	return nil
}
