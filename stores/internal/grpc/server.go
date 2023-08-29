package grpc

import (
	"context"

	"github.com/google/uuid"
	"github.com/hnamzian/go-mallbots/stores/internal/application"
	"github.com/hnamzian/go-mallbots/stores/internal/domain"
	"github.com/hnamzian/go-mallbots/stores/storespb"
	"google.golang.org/grpc"
)

type Server struct {
	app application.App
	storespb.UnimplementedStoresServer
}

func RegisterServer(registrar grpc.ServiceRegistrar, app application.App) {
	storespb.RegisterStoresServer(registrar, &Server{app: app})
}

func (s *Server) CreateStore(ctx context.Context, store *storespb.CreateStoreRequest) (*storespb.CreateStoreResponse, error) {
	id := uuid.New().String()
	if err := s.app.CreateStore(ctx, id, store.Name, store.Location); err != nil {
		return nil, err
	}
	return &storespb.CreateStoreResponse{Id: id}, nil
}

func (s *Server) GetStore(ctx context.Context, get *storespb.GetStoreRequest) (*storespb.GetStoreResponse, error) {
	store, err := s.app.GetStore(ctx, get.Id)
	if err != nil {
		return nil, err
	}
	return &storespb.GetStoreResponse{Store: storeFromDomain(store)}, nil
}

func (s *Server) GetStores(ctx context.Context, get *storespb.GetStoresRequest) (*storespb.GetStoresResponse, error) {
	stores, err := s.app.GetStores(ctx)
	if err != nil {
		return nil, err
	}

	protoStores := []*storespb.Store{}
	for _, store := range stores {
		protoStore := storeFromDomain(store)
		protoStores = append(protoStores, protoStore)
	}

	return &storespb.GetStoresResponse{Stores: protoStores}, nil
}

func (s *Server) EnableParticipation(ctx context.Context, enable *storespb.EnableParticipationRequest) (*storespb.EnableParticipationResponse, error) {
	if err := s.app.EnableParticipation(ctx, enable.Id); err != nil {
		return nil, err
	}
	return &storespb.EnableParticipationResponse{}, nil
}

func (s *Server) DisableParticipation(ctx context.Context, disable *storespb.DisableParticipationRequest) (*storespb.DisableParticipationResponse, error) {
	if err := s.app.DisableParticipation(ctx, disable.Id); err != nil {
		return nil, err
	}
	return &storespb.DisableParticipationResponse{}, nil
}

func (s *Server) GetParticipatingStores(ctx context.Context, get *storespb.GetParticipatingStoresRequest) (*storespb.GetParticipatingStoresResponse, error) {
	stores, err := s.app.GetParticipatingStores(ctx)
	if err != nil {
		return nil, err
	}

	protoStores := []*storespb.Store{}
	for _, store := range stores {
		protoStores = append(protoStores, storeFromDomain(store))
	}
	return &storespb.GetParticipatingStoresResponse{Stores: protoStores}, nil
}

func (s *Server) AddProduct(ctx context.Context, create *storespb.AddProductRequest) (*storespb.AddProductResponse, error) {
	id := uuid.New().String()
	_, err := s.app.CreateProduct(ctx, id, create.StoreId, create.Name, create.Description, create.Sku, create.Price)
	if err != nil {
		return nil, err
	}

	return &storespb.AddProductResponse{Id: id}, nil
}

func (s *Server) GetProduct(ctx context.Context, get *storespb.GetProductRequest) (*storespb.GetProductResponse, error) {
	product, err := s.app.GetProduct(ctx, get.Id)

	if err != nil {
		return nil, err
	}

	return &storespb.GetProductResponse{Product: productFromDomain(product)}, nil
}

func (s *Server) GetCatalog(ctx context.Context, get *storespb.GetCatalogRequest) (*storespb.GetCatalogResponse, error) {
	product, err := s.app.GetCatalog(ctx, get.StoreId)
	if err != nil {
		return nil, err
	}

	protoProducts := []*storespb.Product{}
	for _, product := range product {
		protoProducts = append(protoProducts, productFromDomain(product))
	}

	return &storespb.GetCatalogResponse{Products: protoProducts}, nil
}

func (s *Server) RemoveProduct(ctx context.Context, remove *storespb.RemoveProductRequest) (*storespb.RemoveProductResponse, error) {
	if err := s.app.DeleteProduct(ctx, remove.Id); err != nil {
		return nil, err
	}
	return &storespb.RemoveProductResponse{}, nil
}

func storeFromDomain(store *domain.Store) *storespb.Store {
	return &storespb.Store{
		Id:            store.ID,
		Name:          store.Name,
		Location:      store.Location,
		Participating: store.Participating,
	}
}

func productFromDomain(product *domain.Product) *storespb.Product {
	return &storespb.Product{
		Id:          product.ID,
		StoreId:     product.StoreID,
		Name:        product.Name,
		Description: product.Description,
		Sku:         product.SKU,
		Price:       product.Price,
	}
}
