package domain

type OrderStatus string

const (
	OrderUnknown   OrderStatus = "unknown"
	OrderIsPending   OrderStatus = "pending"
	OrderIsInProcess OrderStatus = "in_process"
	OrderIsReady     OrderStatus = "ready"
	OrderIsCompleted OrderStatus = "completed"
	OrderIsCancelled OrderStatus = "cancelled"
)

func (s OrderStatus) String() string {
	switch s {
	case OrderIsPending, OrderIsInProcess, OrderIsReady, OrderIsCompleted, OrderIsCancelled:
		return string(s)
	default:
		return string(OrderUnknown)
	}
}

func ToOrderStatus(s string) OrderStatus {
	switch s {
	case OrderIsPending.String():
		return OrderIsPending
	case OrderIsInProcess.String():
		return OrderIsInProcess
	case OrderIsReady.String():
		return OrderIsReady
	case OrderIsCompleted.String():
		return OrderIsCompleted
	case OrderIsCancelled.String():
		return OrderIsCancelled
	default:
		return OrderUnknown
	}
}