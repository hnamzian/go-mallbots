package rest

import (
	"context"

	"github.com/go-chi/chi/v5"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/hnamzian/go-mallbots/payments/paymentspb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func RegisterGateway(ctx context.Context, mux *chi.Mux, grpcAddr string) error {
	apiRoot := "/api/payments"

	gateway := runtime.NewServeMux()
	err := paymentspb.RegisterPaymentsServiceHandlerFromEndpoint(ctx, gateway, grpcAddr, []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	})
	if err != nil {
		return err
	}

	mux.Mount(apiRoot, gateway)

	return nil
}
