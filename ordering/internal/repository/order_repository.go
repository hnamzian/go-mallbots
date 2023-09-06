package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/hnamzian/go-mallbots/ordering/internal/domain"
)

type OrderRepository struct {
	tableName string
	db        *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{
		tableName: "orders",
		db:        db,
	}
}

func (r OrderRepository) Save(ctx context.Context, order *domain.Order) error {
	query := fmt.Sprintf("INSERT INTO %s (id, customer_id, payment_id, invoice_id, shopping_id, items, status) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)", r.tableName)

	items, err := json.Marshal(order.Items)
	if err != nil {
		return err
	}

	_, err = r.db.ExecContext(ctx, query, order.ID, order.CustomerID, order.PaymentID, order.InvoiceID, order.ShoppingID, items, order.Status.String())
	if err != nil {
		return err
	}

	return nil
}

func (r OrderRepository) Find(ctx context.Context, id string) (*domain.Order, error) {
	query := fmt.Sprintf("SELECT id, customer_id, payment_id, invoice_id, shopping_id, items, status FROM %s WHERE id = $1", r.tableName)

	var order domain.Order
	var items []byte
	var status string
	err := r.db.QueryRowContext(ctx, query, id).Scan(&order.ID, &order.CustomerID, &order.PaymentID, &order.InvoiceID, &order.ShoppingID, &items, status)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(items, &order.Items); err != nil {
		return nil, err
	}

	order.Status = domain.OrderStatus(status)

	return &order, nil
}

func (r OrderRepository) Update(ctx context.Context, order *domain.Order) error {
	query := fmt.Sprintf("UPDATE %s SET customer_id = $1, payment_id = $2, invoice_id = $3, shopping_id = $4, items = $5, status = $6 WHERE id = $7", r.tableName)

	items, err := json.Marshal(order.Items)
	if err != nil {
		return err
	}

	_, err = r.db.ExecContext(ctx, query, order.CustomerID, order.PaymentID, order.InvoiceID, order.ShoppingID, items, order.Status.String(), order.ID)
	if err != nil {
		return err
	}

	return nil
}
