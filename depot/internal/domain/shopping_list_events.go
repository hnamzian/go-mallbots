package domain

type ShoppingListCreated struct {
	ShoppingList *ShoppingList
}

func (e *ShoppingListCreated) EventName() string {
	return "depot.ShoppingListCreated"
}

type ShoppingListCompleted struct {
	ShoppingList *ShoppingList
}

func (e *ShoppingListCompleted) EventName() string {
	return "depot.ShoppingListCompleted"
}

type ShoppingListCancelled struct {
	ShoppingList *ShoppingList
}

func (e *ShoppingListCancelled) EventName() string {
	return "depot.ShoppingListCanceled"
}

type ShoppingListAssigned struct {
	ShoppingList *ShoppingList
}

func (e *ShoppingListAssigned) EventName() string {
	return "depot.ShoppingListAssigned"
}