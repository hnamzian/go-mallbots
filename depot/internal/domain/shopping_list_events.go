package domain

type ShoppingListCreated struct {
	ShoppingList *ShoppingList
}

func (e *ShoppingListCreated) EventName() string {
	return "depot.ShoppingListCreated"
}

type ShoppingListCompeted struct {
	ShoppingList *ShoppingList
}

func (e *ShoppingListCompeted) EventName() string {
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