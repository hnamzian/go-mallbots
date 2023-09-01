package domain

type Invoice struct {
	ID      string
	OrderID string
	Amount  float64
	Status  InvoiceStatus
}

func CreateInvoice(id, orderID string, amount float64) Invoice {
	return Invoice{
		ID:      id,
		OrderID: orderID,
		Amount:  amount,
		Status:  InvoicePending,
	}
}