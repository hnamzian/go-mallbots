package grpc

import (
	"context"

	"github.com/hnamzian/go-mallbots/notifications/notificationspb"
	"google.golang.org/grpc"
)

type NotificationeRepository struct {
	client notificationspb.NotificationsServiceClient
}

func NewNotificationeRepository(conn *grpc.ClientConn) *NotificationeRepository {
	return &NotificationeRepository{
		client: notificationspb.NewNotificationsServiceClient(conn),
	}
}

func (r NotificationeRepository) NotifyOrderCreated(ctx context.Context, orderID, customerID string) error {
	_, err := r.client.NotifyOrderCreated(ctx, &notificationspb.NotifyOrderCreatedRequest{
		OrderId: orderID,
		CustomerId: customerID,
	})
	return err
}

func (r NotificationeRepository) NotifyOrderCanceled(ctx context.Context, orderID, customerID string) error {
	_, err := r.client.NotifyOrderCanceled(ctx, &notificationspb.NotifyOrderCanceledRequest{
		OrderId: orderID,
		CustomerId: customerID,
	})
	return err
}

func (r NotificationeRepository) NotifyOrderReady(ctx context.Context, orderID, customerID string) error {
	_, err := r.client.NotifyOrderReady(ctx, &notificationspb.NotifyOrderReadyRequest{
		OrderId: orderID,
		CustomerId: customerID,
	})
	return err
}

