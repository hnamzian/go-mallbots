package customers

import (
	"context"

	"github.com/hnamzian/go-mallbots/internal/module"
	"github.com/hnamzian/go-mallbots/customers/internal/grpc"
)

type Module struct {}

func (m Module) Startup(ctx context.Context, app module.Core) error{
	grpc.RegisterServer(app.RPC(), grpc.Server{})

	return nil
}
