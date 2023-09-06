package domain

type InvoiceID string

func (id InvoiceID) String() string {
	return string(id)
}