package domain

type Order struct {
	ID string
	CustomerID string
	PaymentID string
	InvoiceID string
	ShoppingID string
	Items []Item
	Status OrderStatus
}