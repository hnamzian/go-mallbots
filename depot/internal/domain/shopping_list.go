package domain

import (
	"github.com/hnamzian/go-mallbots/internal/ddd"
	"github.com/stackus/errors"
)

var (
	ErrShoppingListCannotBeCancelled = errors.Wrap(errors.ErrBadRequest, "shopping cannot be cancelled")
)

type ShoppingListStatus string

const (
	ShoppingListUnknown     ShoppingListStatus = "unknown"
	ShoppingListIsAvailable ShoppingListStatus = "available"
	ShoppingListIsAssigned  ShoppingListStatus = "assigned"
	ShoppingListIsActive    ShoppingListStatus = "active"
	ShoppingListIsCompleted ShoppingListStatus = "completed"
	ShoppingListIsCancelled ShoppingListStatus = "cancelled"
)

type ShoppingList struct {
	ddd.AggregateBase
	OrderID       string
	AssignedBotID string
	Stops         Stops
	Status        ShoppingListStatus
}

func (s ShoppingListStatus) String() string {
	switch s {
	case ShoppingListIsAvailable, ShoppingListIsAssigned, ShoppingListIsActive, ShoppingListIsCompleted, ShoppingListIsCancelled:
		return string(s)
	default:
		return string(ShoppingListUnknown)
	}
}

func ToShoppingLiistStatus(s string) ShoppingListStatus {
	switch s {
	case ShoppingListIsAvailable.String():
		return ShoppingListIsAvailable
	case ShoppingListIsAssigned.String():
		return ShoppingListIsAssigned
	case ShoppingListIsActive.String():
		return ShoppingListIsActive
	case ShoppingListIsCompleted.String():
		return ShoppingListIsCompleted
	case ShoppingListIsCancelled.String():
		return ShoppingListIsCancelled
	default:
		return ShoppingListUnknown
	}
}

func CreateShoppingList(ID, OrderID string) *ShoppingList {
	return &ShoppingList{
		AggregateBase: ddd.AggregateBase{ID: ID},
		OrderID:       OrderID,
		Status:        ShoppingListIsAvailable,
		Stops:         make(Stops),
	}
}

func (sl *ShoppingList) AddItem(store *Store, product *Product, quantity int) error {
	if _, exists := sl.Stops[store.ID]; !exists {
		sl.Stops[store.ID] = &Stop{
			StoreName:     store.Name,
			StoreLocation: store.Location,
			Items:         make(Items),
		}
	}
	return sl.Stops[store.ID].AddItem(product, quantity)
}

func (sl *ShoppingList) Cancel() error {
	if sl.Status == ShoppingListIsCancelled {
		return ErrShoppingListCannotBeCancelled
	}

	sl.Status = ShoppingListIsCancelled

	return nil
}

func (sl *ShoppingList) AssignBot(id string) error {
	sl.AssignedBotID = id
	sl.Status = ShoppingListIsAssigned
	return nil
}

func (sl *ShoppingList) Complete() error {
	sl.Status = ShoppingListIsCompleted

	sl.AddEvent(
		&ShoppingListCompleted{
			ShoppingList: sl,
		},
	)

	return nil
}
