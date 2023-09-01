package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/hnamzian/go-mallbots/payments/internal/domain"
)

type PaymentRepository struct {
	tableName string
	db        *sql.DB
}

func NewPaymentRepository(db *sql.DB) *PaymentRepository {
	return &PaymentRepository{
		tableName: "payments",
		db:        db,
	}
}

func (r PaymentRepository) Save(ctx context.Context, payment *domain.Payment) error {
	query := fmt.Sprintf(
		"INSERT INTO %s (id, customer_id, amount) VALUES ($1, $2, $3)",
		r.tableName,
	)

	_, err := r.db.ExecContext(ctx, query, payment.ID, payment.CustomerID, payment.Amount)
	if err != nil {
		return err
	}

	return nil
}

func (r PaymentRepository) Update(ctx context.Context, payment *domain.Payment) error {
	query := fmt.Sprintf(
		"UPDATE %s SET customer_id=$1, amount=$2 WHERE id=$3",
		r.tableName,
	)

	_, err := r.db.ExecContext(ctx, query, payment.CustomerID, payment.Amount, payment.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r PaymentRepository) Find(ctx context.Context, id string) (*domain.Payment, error) {
	query := fmt.Sprintf(
		"SELECT id, customer_id, amount FROM %s WHERE id=$1",
		r.tableName,
	)

	row := r.db.QueryRowContext(ctx, query, id)

	var payment domain.Payment
	err := row.Scan(&payment.ID, &payment.CustomerID, &payment.Amount)
	if err != nil {
		return nil, err
	}

	return &payment, nil
}