package domain

type Payment struct {
	ID         string
	CustomerID string
	Amount     float64
}

func CreatePayment(id, customerID string, amount float64) Payment {
	return Payment{
		ID:         id,
		CustomerID: customerID,
		Amount:     amount,
	}
}