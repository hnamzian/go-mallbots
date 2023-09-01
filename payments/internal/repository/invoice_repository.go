package domain

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/hnamzian/go-mallbots/payments/internal/domain"
)

type InvoiceRepository struct {
	tableName string
	db        *sql.DB
}

func NewInvoiceRepository(db *sql.DB) *InvoiceRepository {
	return &InvoiceRepository{
		tableName: "invoices",
		db:        db,
	}
}

func (r InvoiceRepository) Save(ctx context.Context, invoice *domain.Invoice) error {
	query := fmt.Sprintf(
		"INSERT INTO %s (id, order_id, amount, status) VALUES ($1, $2, $3, $4)",
		r.tableName,
	)

	_, err := r.db.ExecContext(ctx, query, invoice.ID, invoice.OrderID, invoice.Amount, invoice.Status)
	if err != nil {
		return err
	}

	return nil
}

func (r InvoiceRepository) Update(ctx context.Context, invoice *domain.Invoice) error {
	query := fmt.Sprintf(
		"UPDATE %s SET order_id=$1, amount=$2, status=$3 WHERE id=$4",
		r.tableName,
	)

	_, err := r.db.ExecContext(ctx, query, invoice.OrderID, invoice.Amount, invoice.Status, invoice.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r InvoiceRepository) Find(ctx context.Context, id string) (*domain.Invoice, error) {
	query := fmt.Sprintf(
		"SELECT id, order_id, amount, status FROM %s WHERE id=$1",
		r.tableName,
	)

	row := r.db.QueryRowContext(ctx, query, id)

	var invoice domain.Invoice
	err := row.Scan(&invoice.ID, &invoice.OrderID, &invoice.Amount, &invoice.Status)
	if err != nil {
		return nil, err
	}

	return &invoice, nil
}
