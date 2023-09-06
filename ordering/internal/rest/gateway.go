package rest

import (
	"context"

	"github.com/go-chi/chi/v5"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/hnamzian/go-mallbots/ordering/orderingpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Registergateway(ctx context.Context, mux *chi.Mux, addr string) error {
	apiRoot := "/api/ordering"

	gateway := runtime.NewServeMux()
	err := orderingpb.RegisterOrderingServiceHandlerFromEndpoint(ctx, gateway, addr, []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	})
	if err != nil {
		return err
	}

	mux.Mount(apiRoot, gateway)

	return nil
}