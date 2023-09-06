package domain

type OrderStatus string

const (
	OrderUnknown   OrderStatus = "unknown"
	OrderPending   OrderStatus = "pending"
	OrderInProcess OrderStatus = "in_process"
	OrderReady     OrderStatus = "ready"
	OrderCompleted OrderStatus = "completed"
	OrderCancelled OrderStatus = "cancelled"
)

func (s OrderStatus) String() string {
	switch s {
	case OrderPending, OrderInProcess, OrderReady, OrderCompleted, OrderCancelled:
		return string(s)
	default:
		return string(OrderUnknown)
	}
}

func ToOrderStatus(s string) OrderStatus {
	switch s {
	case OrderPending.String():
		return OrderPending
	case OrderInProcess.String():
		return OrderInProcess
	case OrderReady.String():
		return OrderReady
	case OrderCompleted.String():
		return OrderCompleted
	case OrderCancelled.String():
		return OrderCancelled
	default:
		return OrderUnknown
	}
}