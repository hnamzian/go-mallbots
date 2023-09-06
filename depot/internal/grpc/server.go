package grpc

import (
	"context"

	"github.com/google/uuid"
	"github.com/hnamzian/go-mallbots/depot/depotpb"
	"github.com/hnamzian/go-mallbots/depot/internal/application"
	"github.com/hnamzian/go-mallbots/depot/internal/domain"
	"google.golang.org/grpc"
)

type DepotServer struct {
	application.App
	depotpb.UnimplementedDepotServiceServer
}

func RegisterServer(registrar grpc.ServiceRegistrar, app application.App) error {
	depotpb.RegisterDepotServiceServer(registrar, &DepotServer{App: app})
	return nil
}

func (s DepotServer) CreateShoppingList(ctx context.Context, request *depotpb.CreateShoppingListRequest) (*depotpb.CreateShoppingListResponse, error) {
	id := uuid.New().String()
	orderItems := []*application.OrderItem{}
	for _, item := range request.Items {
		orderItems = append(orderItems, &application.OrderItem{
			StoreID:   item.StoreId,
			ProductID: item.ProductId,
			Quantity:  item.Quantity,
		})
	}
	err := s.App.CreateShoppingList(ctx, &application.CreateShoppingList{
		ID:      id,
		OrderID: request.OrderId,
		Items:   orderItems,
	})
	if err != nil {
		return nil, err
	}
	return &depotpb.CreateShoppingListResponse{Id: id}, nil
}
func (s DepotServer) CancelShoppingList(ctx context.Context, request *depotpb.CancelShoppingListRequest) (*depotpb.CancelShoppingListResponse, error) {
	err := s.App.CancelShoppingList(ctx, &application.CancelShoppingList{ID: request.Id})
	return &depotpb.CancelShoppingListResponse{}, err
}
func (s DepotServer) CompleteShoppingList(ctx context.Context, request *depotpb.CompleteShoppingListRequest) (*depotpb.CompleteShoppingListResponse, error) {
	err := s.App.CompleteShoppingList(ctx, &application.CompleteShoppingList{ID: request.Id})
	return &depotpb.CompleteShoppingListResponse{}, err
}
func (s DepotServer) AssignBotToShoppingList(ctx context.Context, request *depotpb.AssignBotToShoppingListRequest) (*depotpb.AssignBotToShoppingListResponse, error) {
	err := s.App.AssignBotToShoppingList(ctx, &application.AssignBotToShoppingList{ID: request.Id, BotID: request.BotId})
	return &depotpb.AssignBotToShoppingListResponse{}, err
}
func (s DepotServer) GetShoppingList(ctx context.Context, request *depotpb.GetShoppingListRequest) (*depotpb.GetShoppingListResponse, error) {
	list, err := s.App.GetShoppingList(ctx, &application.GetShoppingList{ID: request.Id})
	if err != nil {
		return nil, err
	}
	return &depotpb.GetShoppingListResponse{
		ShoppingList: shoppingListFromDomain(list),
	}, nil
}

func shoppingListFromDomain(list *domain.ShoppingList) *depotpb.ShoppingList {
	stops := make(map[string]*depotpb.Stop, 0)
	for id, stop := range list.Stops {
		stops[id] = &depotpb.Stop{
			StoreName:     stop.StoreName,
			StoreLocation: stop.StoreLocation,
		}
		items := make(map[string]*depotpb.Item, 0)
		for itemID, item := range stop.Items {
			items[itemID] = &depotpb.Item{
				ProductName: item.ProductName,
				Quantity:    int32(item.Quantity),
			}
		}
		stops[id].Items = items
	}
	return &depotpb.ShoppingList{
		Id:            list.ID,
		OrderId:       list.OrderID,
		AssignedBotId: list.AssignedBotID,
		Stops:         stops,
		Status:        string(list.Status),
	}
}
