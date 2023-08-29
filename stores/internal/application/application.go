package application

import (
	"context"

	"github.com/hnamzian/go-mallbots/stores/internal/domain"
)

type App interface {
	CreateStore(ctx context.Context, id, name, location string) error
	GetStore(ctx context.Context, id string) (*domain.Store, error)
	GetStores(ctx context.Context) ([]*domain.Store, error)
	EnableParticipation(ctx context.Context, id string) error
	DisableParticipation(ctx context.Context, id string) error
	GetParticipatingStores(ctx context.Context) ([]*domain.Store, error)
	CreateProduct(ctx context.Context, id string, storeID string, name string, description string, sku string, price float64) (*domain.Product, error)
	GetProduct(ctx context.Context, id string) (*domain.Product, error)
	GetCatalog(ctx context.Context, storeID string) ([]*domain.Product, error)
	DeleteProduct(ctx context.Context, id string) error
}

type Application struct {
	domain.StoreRepository
	domain.ParticipatingStoreRepository
	domain.ProductRepository
}

func NewApplication(storeRepository domain.StoreRepository, productRepository domain.ProductRepository, participatingStoreRepository domain.ParticipatingStoreRepository) Application {
	return Application{
		StoreRepository:              storeRepository,
		ProductRepository:            productRepository,
		ParticipatingStoreRepository: participatingStoreRepository,
	}
}

func (a Application) CreateStore(ctx context.Context, id, name, location string) error {
	store, err := domain.CreateStore(id, name, location)
	if err != nil {
		return err
	}
	return a.StoreRepository.Save(ctx, store)
}

func (a Application) GetStore(ctx context.Context, id string) (*domain.Store, error) {
	return a.StoreRepository.Get(ctx, id)
}

func (a Application) GetStores(ctx context.Context) ([]*domain.Store, error) {
	return a.StoreRepository.GetAll(ctx)
}

func (a Application) EnableParticipation(ctx context.Context, id string) error {
	store, err := a.StoreRepository.Get(ctx, id)
	if err != nil {
		return err
	}

	if err = store.EnableParticipation(); err != nil {
		return err
	}

	return a.StoreRepository.Update(ctx, store)
}

func (a Application) DisableParticipation(ctx context.Context, id string) error {
	store, err := a.StoreRepository.Get(ctx, id)
	if err != nil {
		return err
	}

	if err = store.DisableParticipation(); err != nil {
		return err
	}

	return a.StoreRepository.Update(ctx, store)
}

func (a Application) GetParticipatingStores(ctx context.Context) ([]*domain.Store, error) {
	return a.ParticipatingStoreRepository.GetAll(ctx)
}

func (a Application) CreateProduct(ctx context.Context, id string, storeID string, name string, description string, sku string, price float64) (*domain.Product, error) {
	product, err := domain.CreateProduct(id, storeID, name, description, sku, price)
	if err != nil {
		return nil, err
	}

	if err := a.ProductRepository.Save(ctx, product); err != nil {
		return nil, err
	}

	return product, nil
}

func (a Application) GetProduct(ctx context.Context, id string) (*domain.Product, error) {
	return a.ProductRepository.Get(ctx, id)
}

func (a Application) GetCatalog(ctx context.Context, storeID string) ([]*domain.Product, error) {
	return a.ProductRepository.GetCatalog(ctx, storeID)
}

func (a Application) DeleteProduct(ctx context.Context, id string) error {
	return a.ProductRepository.Delete(ctx, id)
}
