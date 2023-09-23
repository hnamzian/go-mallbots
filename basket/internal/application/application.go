package application

import (
	"context"

	"github.com/hnamzian/go-mallbots/basket/internal/domain"
	"github.com/hnamzian/go-mallbots/internal/ddd"
)

type App interface {
	StartBasket(ctx context.Context, start *StartBasket) error
	CancelBasket(ctx context.Context, cancel *CancelBasket) error
	CheckoutBasket(ctx context.Context, checkout CheckoutBasket) error
	AddItem(ctx context.Context, add *AddItem) error
	RemoveItem(ctx context.Context, remove *RemoveItem) error
	GetBasket(ctx context.Context, get *GetBasket) (*domain.Basket, error)
}

type Application struct {
	baskets         domain.BasketRepository
	products        domain.ProductRepository
	stores          domain.StoreRepository
	domainPublisher ddd.EventPublisher
}

type (
	StartBasket struct {
		ID         string
		CustomerID string
	}

	CancelBasket struct {
		ID string
	}

	CheckoutBasket struct {
		ID        string
		PaymentID string
	}

	AddItem struct {
		ID        string
		ProductID string
		StoreID   string
		Quantity  int
	}

	RemoveItem struct {
		ID        string
		ProductID string
		StoreID   string
		Quantity  int
	}

	GetBasket struct {
		ID string
	}
)

func NewApplication(baskets domain.BasketRepository, products domain.ProductRepository, stores domain.StoreRepository, domainPublisher ddd.EventPublisher) Application {
	return Application{
		baskets,
		products,
		stores,
		domainPublisher,
	}
}

func (a Application) StartBasket(ctx context.Context, start *StartBasket) error {
	basket, err := domain.StartBasket(start.ID, start.CustomerID)
	if err != nil {
		return err
	}

	if err := a.baskets.Save(ctx, basket); err != nil {
		return err
	}

	a.domainPublisher.Publish(ctx, basket.GetEvents()...)

	return nil
}

func (a Application) CancelBasket(ctx context.Context, cancel *CancelBasket) error {
	basket, err := a.baskets.Get(ctx, cancel.ID)
	if err != nil {
		return err
	}

	if err = basket.Cancel(); err != nil {
		return err
	}

	if err = a.baskets.Update(ctx, basket); err != nil {
		return nil
	}

	a.domainPublisher.Publish(ctx, basket.GetEvents()...)

	return nil
}

func (a Application) CheckoutBasket(ctx context.Context, checkout CheckoutBasket) error {
	basket, err := a.baskets.Get(ctx, checkout.ID)
	if err != nil {
		return err
	}

	if err = basket.Checkout(checkout.PaymentID); err != nil {
		return err
	}

	if err = a.baskets.Update(ctx, basket); err != nil {
		return nil
	}

	a.domainPublisher.Publish(ctx, basket.GetEvents()...)

	return nil
}

func (a Application) AddItem(ctx context.Context, add *AddItem) error {
	basket, err := a.baskets.Get(ctx, add.ID)
	if err != nil {
		return err
	}

	product, err := a.products.Find(ctx, add.ProductID)
	if err != nil {
		return err
	}

	store, err := a.stores.Find(ctx, add.StoreID)
	if err != nil {
		return err
	}

	if err = basket.AddItem(store, product, add.Quantity); err != nil {
		return err
	}

	a.domainPublisher.Publish(ctx, basket.GetEvents()...)

	return nil
}

func (a Application) RemoveItem(ctx context.Context, remove *RemoveItem) error {
	basket, err := a.baskets.Get(ctx, remove.ID)
	if err != nil {
		return err
	}

	product, err := a.products.Find(ctx, remove.ProductID)
	if err != nil {
		return err
	}

	if err = basket.RemoveItem(product, remove.Quantity); err != nil {
		return err
	}

	a.domainPublisher.Publish(ctx, basket.GetEvents()...)

	return nil
}

func (a Application) GetBasket(ctx context.Context, get *GetBasket) (*domain.Basket, error) {
	return a.baskets.Get(ctx, get.ID)
}
