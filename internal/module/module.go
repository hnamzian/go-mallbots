package module

import (
	"context"
	"database/sql"

	"github.com/hnamzian/go-mallbots/internal/config"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
)

type Core interface {
	Config() *config.AppConfig
	Logger() zerolog.Logger
	DB() *sql.DB
	RPC() *grpc.Server
}

type Module interface {
	Startup(ctx context.Context, c Core) error
}
