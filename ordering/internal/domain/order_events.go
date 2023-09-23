package domain

type OrderCreated struct {
	Order *Order
}

func (OrderCreated) EventName() string {
	return "ordering.OrderCreated"
}

type OrderCancelled struct {
	Order *Order
}

func (OrderCancelled) EventName() string {
	return "ordering.OrderCancelled"
}

type OrderReadied struct {
	Order *Order
}

func (OrderReadied) EventName() string {
	return "ordering.OrderReadied"
}

type OrderCompleted struct {
	Order *Order
}

func (OrderCompleted) EventName() string {
	return "ordering.OrderCompleted"
}