package application

import (
	"context"

	"github.com/hnamzian/go-mallbots/depot/internal/domain"
	"github.com/stackus/errors"
)

type App interface {
	CreateShoppingList(context.Context, *CreateShoppingList) error
	CancelShoppingList(context.Context, *CancelShoppingList) error
	CompleteShoppingList(context.Context, *CompleteShoppingList) error
	AssignBotToShoppingList(context.Context, *AssignBotToShoppingList) error
	GetShoppingList(context.Context, *GetShoppingList) (*domain.ShoppingList, error)
}

type Application struct {
	shoppingLists domain.ShoppingListRepository
	orders        domain.OrderRepository
	products     domain.ProductRepository
	stores        domain.StoreRepository
}

type (
	CreateShoppingList struct {
		ID      string
		OrderID string
		Items   []*OrderItem
	}

	CancelShoppingList struct {
		ID string
	}

	CompleteShoppingList struct{
		ID string
	}

	AssignBotToShoppingList struct{
		ID string
		BotID string
	}

	GetShoppingList struct{
		ID string
	}

	OrderItem struct {
		StoreID   string
		ProductID string
		Quantity  int32
	}
)

func NewDepotApplication(shoppingLists domain.ShoppingListRepository,
	orders domain.OrderRepository,
	pproducts domain.ProductRepository,
	stores domain.StoreRepository) *Application {
	return &Application{
		shoppingLists: shoppingLists,
		orders:        orders,
		products:     pproducts,
		stores:        stores,
	}
}

func (a Application) CreateShoppingList(ctx context.Context, create *CreateShoppingList) error {
	shoppingList := domain.CreateShoppingList(create.ID, create.OrderID)

	for _, item := range create.Items {
		store, err := a.stores.Find(ctx, item.StoreID)
		if err != nil {
			return errors.Wrap(err, "create shopping list")
		}
		product, err := a.products.Find(ctx, item.ProductID)
		if err != nil {
			return errors.Wrap(err, "create shopping list")
		}
		if err = shoppingList.AddItem(store, product, int(item.Quantity)); err != nil {
			return errors.Wrap(err, "create shopping list")
		}
	}

	return nil
}

func (a Application) CancelShoppingList(ctx context.Context, cancel *CancelShoppingList) error {
	list, err := a.shoppingLists.Find(ctx, cancel.ID)
	if err != nil {
		return errors.Wrap(err, "cancel shopping list")
	}
	if err = list.Cancel(); err != nil {
		return errors.Wrap(err, "cancel shopping list")
	}
	return a.shoppingLists.Update(ctx, list)
}

func (a Application) CompleteShoppingList(ctx context.Context, complete *CompleteShoppingList) error {
	list, err := a.shoppingLists.Find(ctx, complete.ID)
	if err != nil {
		return errors.Wrap(err, "complete shopping list")
	}
	if err := list.Complete(); err != nil {
		return errors.Wrap(err, "complete shopping list")
	}
	return a.shoppingLists.Update(ctx, list)
}

func (a Application) AssignBotToShoppingList(ctx context.Context, assign *AssignBotToShoppingList) error {
	list, err := a.shoppingLists.Find(ctx, assign.ID)
	if err != nil {
		return errors.Wrap(err, "assign shopping list")
	}
	if err := list.AssignBot(assign.BotID); err != nil {
		return errors.Wrap(err, "assign shopping list")
	}
	return a.shoppingLists.Update(ctx, list)
}

func (a Application) GetShoppingList(ctx context.Context, get *GetShoppingList) (*domain.ShoppingList, error) {
	list, err := a.shoppingLists.Find(ctx, get.ID)
	if err != nil {
		return nil, errors.Wrap(err, "get shopping list")
	}
	return list, nil
}
