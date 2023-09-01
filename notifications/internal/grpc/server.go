package grpc

import (
	"context"

	"github.com/hnamzian/go-mallbots/notifications/internal/application"
	"github.com/hnamzian/go-mallbots/notifications/notificationspb"
	"google.golang.org/grpc"
)

type Server struct {
	application.App
	notificationspb.UnimplementedNotificationsServiceServer
}

func RegisterServer(registrar grpc.ServiceRegistrar, app application.App) {
	notificationspb.RegisterNotificationsServiceServer(registrar, Server{App: app})
}

func (s Server) NotifyOrderCreated(ctx context.Context, request *notificationspb.NotifyOrderCreatedRequest) (*notificationspb.NotifyOrderCreatedResponse, error) {
	s.App.NotifyOrderCreated(ctx, application.OrderCreated{
		OrderID:    request.OrderId,
		CustomerID: request.CustomerId,
	})
	return &notificationspb.NotifyOrderCreatedResponse{}, nil
}

func (s Server) NotifyOrderCanceled(ctx context.Context, request *notificationspb.NotifyOrderCanceledRequest) (*notificationspb.NotifyOrderCanceledResponse, error) {
	s.App.NotifyOrderCanceled(ctx, application.OrderCanceled{
		OrderID:    request.OrderId,
		CustomerID: request.CustomerId,
	})
	return &notificationspb.NotifyOrderCanceledResponse{}, nil
}

func (s Server) NotifyOrderReady(ctx context.Context, request *notificationspb.NotifyOrderReadyRequest) (*notificationspb.NotifyOrderReadyResponse, error) {
	s.App.NotifyOrderReady(ctx, application.OrderReady{
		OrderID:    request.OrderId,
		CustomerID: request.CustomerId,
	})
	return &notificationspb.NotifyOrderReadyResponse{}, nil
}
