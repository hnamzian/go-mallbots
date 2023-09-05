package domain

import "github.com/stackus/errors"

var (
	ErrShoppingListCannotBeCancelled = errors.Wrap(errors.ErrBadRequest, "shopping cannot be cancelled")
)

type ShoppingListStatus string

const (
	ShoppingListUnknown   ShoppingListStatus = "unknown"
	ShoppingListAvailable ShoppingListStatus = "available"
	ShoppingListAssigned  ShoppingListStatus = "assigned"
	ShoppingListActive    ShoppingListStatus = "active"
	ShoppingListCompleted ShoppingListStatus = "completed"
	ShoppingListCancelled ShoppingListStatus = "cancelled"
)

type ShoppingList struct {
	ID            string
	OrderID       string
	AssignedBotID string
	Stops         Stops
	Status        ShoppingListStatus
}

func (s ShoppingListStatus) String() string {
	switch s {
	case ShoppingListAvailable, ShoppingListAssigned, ShoppingListActive, ShoppingListCompleted, ShoppingListCancelled:
		return string(s)
	default:
		return string(ShoppingListUnknown)
	}
}

func ToShoppingLiistStatus(s string) ShoppingListStatus {
	switch s {
	case ShoppingListAvailable.String():
		return ShoppingListAvailable
	case ShoppingListAssigned.String():
		return ShoppingListAssigned
	case ShoppingListActive.String():
		return ShoppingListActive
	case ShoppingListCompleted.String():
		return ShoppingListCompleted
	case ShoppingListCancelled.String():
		return ShoppingListCancelled
	default:
		return ShoppingListUnknown
	}
}

func CreateShoppingList(ID, OrderID string) *ShoppingList {
	return &ShoppingList{
		ID:      ID,
		OrderID: OrderID,
		Status:  ShoppingListAvailable,
		Stops:   make(Stops),
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
	if sl.Status == ShoppingListCancelled {
		return ErrShoppingListCannotBeCancelled
	}

	sl.Status = ShoppingListCancelled

	return nil
}

func (sl *ShoppingList) AssignBot(id string) error {
	sl.AssignedBotID = id
	sl.Status = ShoppingListAssigned
	return nil
}

func (sl *ShoppingList) Complete() error {
	sl.Status = ShoppingListCompleted
	return nil
}
