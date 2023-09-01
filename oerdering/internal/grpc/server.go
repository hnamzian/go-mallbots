package grpc

import (
	"context"

	"github.com/google/uuid"
	"github.com/hnamzian/go-mallbots/oerdering/internal/application"
	"github.com/hnamzian/go-mallbots/oerdering/internal/domain"
	"github.com/hnamzian/go-mallbots/oerdering/orderingpb"
)

type Server struct {
	app application.App
	orderingpb.UnimplementedOrderingServiceServer
}

func RegisterServer(app application.App) *Server {
	return &Server{app: app}
}

func (s *Server) CreateOrder(ctx context.Context, request *orderingpb.CreateOrderRequest) (*orderingpb.CreateOrderResponse, error) {
	id := uuid.New().String()

	err := s.app.CreateOrder(ctx, application.CreateOrder{
		ID:         id,
		CustomerID: request.CustomerId,
		PaymentID:  request.PaymentId,
		Items:      itemsFromProto(request.Items),
	})
	if err != nil {
		return nil, err
	}

	return &orderingpb.CreateOrderResponse{Id: id}, nil
}

func (s *Server) GetOrder(ctx context.Context, request *orderingpb.GetOrderRequest) (*orderingpb.GetOrderResponse, error) {
	order, err := s.app.GetOrder(ctx, application.GetOrder{ID: request.Id})
	if err != nil {
		return nil, err
	}
	return &orderingpb.GetOrderResponse{Order: orderFromDomain(order)}, nil
}

func (s *Server) CancelOrder(ctx context.Context, request *orderingpb.CancelOrderRequest) (*orderingpb.CancelOrderResponse, error) {
	err := s.app.CancelOrder(ctx, application.CancelOrder{ID: request.Id})
	return &orderingpb.CancelOrderResponse{}, err
}

func (s *Server) ReadyOrder(ctx context.Context, request *orderingpb.ReadyOrderRequest) (*orderingpb.ReadyOrderResponse, error) {
	err := s.app.ReadyOrder(ctx, application.ReadyOrder{ID: request.Id})
	return &orderingpb.ReadyOrderResponse{}, err
}

func (s *Server) CompletedOrder(ctx context.Context, request *orderingpb.CompletedOrderRequest) (*orderingpb.CompletedOrderResponse, error) {
	err := s.app.CompletedOrder(ctx, application.CompletedOrder{ID: request.Id, InvoiceID: request.InvoiceId})
	return &orderingpb.CompletedOrderResponse{}, err
}

func itemsFromProto(items []*orderingpb.Item) []domain.Item {
	var result []domain.Item
	for _, item := range items {
		result = append(result, domain.Item{
			ProductID:   item.ProductId,
			StoreID:     item.StoreId,
			ProductName: item.ProductName,
			StoreName:   item.StoreName,
			Price:       item.Price,
			Quantity:    int(item.Quantity),
		})
	}
	return result
}

func orderFromDomain(order *domain.Order) *orderingpb.Order {
	var items []*orderingpb.Item
	for _, item := range order.Items {
		items = append(items, &orderingpb.Item{
			ProductId:   item.ProductID,
			StoreId:     item.StoreID,
			ProductName: item.ProductName,
			StoreName:   item.StoreName,
			Price:       item.Price,
			Quantity:    int32(item.Quantity),
		})
	}

	return &orderingpb.Order{
		Id:         order.ID,
		CustomerId: order.CustomerID,
		PaymentId:  order.PaymentID,
		InvoiceId:  order.InvoiceID,
		ShoppingId: order.ShoppingID,
		Items:      items,
		Status:     order.Status.String(),
	}
}