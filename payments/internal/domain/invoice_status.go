package domain

type InvoiceStatus string

const (
	InvoiceUnknown   InvoiceStatus = "unknown"
	InvoicePending   InvoiceStatus = "pending"
	InvoicePaid      InvoiceStatus = "paid"
	InvoiceCancelled InvoiceStatus = "cancelled"
)

func (s InvoiceStatus) String() string {
	switch s {
	case InvoicePending, InvoicePaid, InvoiceCancelled:
		return string(s)
	default:
		return string(InvoiceUnknown)
	}
}
