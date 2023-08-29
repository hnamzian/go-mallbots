package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/hnamzian/go-mallbots/basket/internal/domain"
	"github.com/stackus/errors"
)

type BasketRepository struct {
	db        *sql.DB
	tableName string
}

func NewBasketRepository(tableName string, db *sql.DB) *BasketRepository {
	return &BasketRepository{
		db:        db,
		tableName: tableName,
	}
}

func (r BasketRepository) Save(ctx context.Context, basket *domain.Basket) error {
	query := fmt.Sprintf(
		"INSERT INTO %s (id, customer_id, items, payment_id, status) VALUES ($1, $2, $3, $4, $5)", r.tableName,
	)

	items, err := json.Marshal(basket.Items)
	if err != nil {
		return errors.ErrInternal.Err(err)
	}

	_, err = r.db.ExecContext(ctx, query, basket.ID, basket.CustomerID, items, basket.PaymentID, basket.Status)
	if err != nil {
		return errors.ErrInternal.Err(err)
	}

	return nil
}

func (r BasketRepository) Get(ctx context.Context, id string) (*domain.Basket, error) {
	query := fmt.Sprintf(
		"SELECT id, customer_id, items, payment_id, status FROM %s WHERE id = $1", r.tableName,
	)

	basket := &domain.Basket{}
	var items []byte
	var status string
	err := r.db.QueryRowContext(ctx, query, id).Scan(&basket.ID, &basket.CustomerID, &items, &basket.PaymentID, &status)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(items, &basket.Items); err != nil {
		return nil, errors.ErrInternal.Err(err)
	}

	basket.Status, err = statusToDomain(status)
	if err != nil {
		return nil, errors.ErrInternal.Err(err)
	}

	return basket, nil
}

func (r BasketRepository) Update(ctx context.Context, basket *domain.Basket) error {
	query := fmt.Sprintf(
		"UPDATE %s SET customer_id = $1, items = $2, payment_id = $3, status = $4 WHERE id = $5", r.tableName,
	)

	items, err := json.Marshal(basket.Items)
	if err != nil {
		return errors.ErrInternal.Err(err)
	}

	_, err = r.db.ExecContext(ctx, query, basket.CustomerID, items, basket.PaymentID, basket.Status, basket.ID)
	if err != nil {
		return errors.ErrInternal.Err(err)
	}

	return nil
}

func (r BasketRepository) Delete(ctx context.Context, id string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", r.tableName)

	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return errors.ErrInternal.Err(err)
	}

	return nil
}

func statusToDomain(status string) (domain.BasketStatus, error) {
	switch status {
	case domain.BasketOpen.String():
		return domain.BasketOpen, nil
	case domain.BasketCancelled.String():
		return domain.BasketCancelled, nil
	case domain.BasketCheckedOut.String():
		return domain.BasketCheckedOut, nil
	default:
		return domain.BasketUnknown, fmt.Errorf("unknown basket status: %s", status)
	}
}
