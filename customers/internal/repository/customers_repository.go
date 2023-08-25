package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/hnamzian/go-mallbots/customers/internal/domain"
)

type CustomersRepository struct {
	db        *sql.DB
	tableName string
}

func NewCustomersRepository(tableName string, db *sql.DB) *CustomersRepository {
	return &CustomersRepository{db, tableName}
}

func (r *CustomersRepository) Save(ctx context.Context, customer *domain.Customer) error {
	query := fmt.Sprintf("INSERT INTO %s (id, name, sms_number, enabled) VALUES ($1, $2, $3, $4)", r.tableName)

	_, err := r.db.ExecContext(ctx, query, customer.ID, customer.Name, customer.SmsNumber)

	return err
}

func (r *CustomersRepository) Get(ctx context.Context, customerID string) (*domain.Customer, error) {
	query := fmt.Sprintf("SELECT id, name, sms_number, enabled FROM %s WHERE id = $1", r.tableName)

	customer := &domain.Customer{
		ID: customerID,
	}
	err := r.db.QueryRowContext(ctx, query, customerID).Scan(&customer.Name, &customer.SmsNumber, &customer.Enabled)
	
	return customer, err
}

func (r *CustomersRepository) Update(ctx context.Context, customer *domain.Customer) error {
	query := fmt.Sprintf("UPDATE %s SET name = $1, sms_number = $2, enabled = $3 WHERE id = $4", r.tableName)

    _, err := r.db.ExecContext(ctx, query, customer.Name, customer.SmsNumber, customer.Enabled, customer.ID)

    return err
}
