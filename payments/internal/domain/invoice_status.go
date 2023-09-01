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

func ToInvoiceStatus(status string) InvoiceStatus {
	switch status {
	case string(InvoicePending):
		return InvoicePending
	case string(InvoicePaid):
		return InvoicePaid
	case string(InvoiceCancelled):
		return InvoiceCancelled
	default:
		return InvoiceUnknown
	}
}