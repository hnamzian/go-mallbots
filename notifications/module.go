package notifications

import (
	"context"

	"github.com/hnamzian/go-mallbots/internal/module"
	"github.com/hnamzian/go-mallbots/notifications/internal/application"
	"github.com/hnamzian/go-mallbots/notifications/internal/logger"
	"github.com/hnamzian/go-mallbots/notifications/internal/grpc"
)

type Module struct{}

func (m Module) Startup(ctx context.Context, core module.Core) error {
	conn, err := grpc.Dial(ctx, core.Config().Grpc.Address())
	if err != nil {
		return err
	}
	customers := grpc.NewCustomerRepository(conn)
	
	var app application.App
	app = application.NewApplication(customers)
	app = logger.NewApplication(app, core.Logger())

	grpc.RegisterServer(core.RPC(), app)
	
	return nil
}
