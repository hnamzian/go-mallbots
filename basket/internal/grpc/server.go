package grpc

import (
	"context"

	"github.com/google/uuid"
	"github.com/hnamzian/go-mallbots/basket/basketspb"
	"github.com/hnamzian/go-mallbots/basket/internal/application"
	"github.com/hnamzian/go-mallbots/basket/internal/domain"
	"google.golang.org/grpc"
)

type Server struct {
	application.App
	basketspb.UnimplementedBasketServiceServer
}

func RegisterServer(registrar grpc.ServiceRegistrar, app application.App) {
	basketspb.RegisterBasketServiceServer(registrar, &Server{App: app})
}

func (s Server) StartBasket(ctx context.Context, request *basketspb.StartBasketRequest) (*basketspb.StartBasketResponse, error) {
	id := uuid.New().String()
	err := s.App.StartBasket(ctx, &application.StartBasket{
		ID:         id,
		CustomerID: request.CustomerId,
	})
	return &basketspb.StartBasketResponse{Id: id}, err
}

func (s Server) CancelBasket(ctx context.Context, request *basketspb.CancelBasketRequest) (*basketspb.CancelBasketResponse, error) {
	err := s.App.CancelBasket(ctx, &application.CancelBasket{ID: request.Id})
	return &basketspb.CancelBasketResponse{}, err
}

func (s Server) CheckoutBasket(ctx context.Context, request *basketspb.CheckoutBasketRequest) (*basketspb.CheckoutBasketResponse, error) {
	err := s.App.CheckoutBasket(ctx, application.CheckoutBasket{
		ID:        request.Id,
		PaymentID: request.PaymentId,
	})
	return &basketspb.CheckoutBasketResponse{}, err
}

func (s Server) AddItem(ctx context.Context, request *basketspb.AddItemRequest) (*basketspb.AddItemResponse, error) {
	err := s.App.AddItem(ctx, &application.AddItem{
		ID:        request.Id,
		ProductID: request.ProductId,
		StoreID:   request.StoreId,
		Quantity:  int(request.Quantity),
	})
	return &basketspb.AddItemResponse{}, err
}

func (s Server) RemoveItem(ctx context.Context, request *basketspb.RemoveItemRequest) (*basketspb.RemoveItemResponse, error) {
	err := s.App.RemoveItem(ctx, &application.RemoveItem{
		ID:        request.Id,
		ProductID: request.ProductId,
		StoreID:   request.StoreId,
		Quantity:  int(request.Quantity),
	})
	return &basketspb.RemoveItemResponse{}, err
}

func (s Server) GetBasket(ctx context.Context, request *basketspb.GetBasketRequest) (*basketspb.GetBasketResponse, error) {
	basket, err := s.App.GetBasket(ctx, &application.GetBasket{ID: request.Id})
	if err != nil {
		return nil, err
	}

	return &basketspb.GetBasketResponse{Basket: basketFromDomain(basket)}, nil
}

func basketFromDomain(basket *domain.Basket) *basketspb.Basket {
	protoBasket := &basketspb.Basket{
		Id: basket.ID,
	}

	items := make([]*basketspb.Item, 0, len(basket.Items))

	for i, item := range basket.Items {
		items[i] = &basketspb.Item{
			ProductId:   item.ProductID,
			StoreId:     item.StoreID,
			ProductName: item.ProductName,
			StoreName:   item.StoreName,
			Price:       item.ProductPrice,
			Quantity:    int32(item.Quantity),
		}
	}

	protoBasket.Items = items

	return protoBasket
}
