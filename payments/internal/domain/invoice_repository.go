package domain

import "context"

type InvoiceRepository interface {
	Save(ctx context.Context, invoice *Invoice) error
	Update(ctx context.Context, invoice *Invoice) error
	Find(ctx context.Context, id string) (*Invoice, error)
}
