package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/hnamzian/go-mallbots/depot/internal/domain"
)

type ShoppingListRepository struct {
	tableName string
	db        *sql.DB
}

func NewShoppingListRepository(db *sql.DB) *ShoppingListRepository {
	return &ShoppingListRepository{
		tableName: "shopping_lists",
		db:        db,
	}
}

func (r *ShoppingListRepository) Find(ctx context.Context, id string) (*domain.ShoppingList, error) {
	query := fmt.Sprintf("SELECT order_id, stops, assigned_bot_id, status FROM %s WHERE id = $1 LIMIT 1", r.tableName)

	shopping := &domain.ShoppingList{
		ID: id,
	}
	var stops []byte
	var status string
	if err := r.db.QueryRowContext(ctx, query, id).Scan(&shopping.OrderID, &stops, &shopping.AssignedBotID, &status); err != nil {
		return nil, err
	}

	if err := json.Unmarshal(stops, &shopping.Stops); err != nil {
		return nil, err
	}

	shopping.Status = domain.ShoppingListStatus(status)

	return shopping, nil
}
func (r *ShoppingListRepository) FindByOrderID(ctx context.Context, orderID string) (*domain.ShoppingList, error) {
	query := fmt.Sprintf("SELECT id, stops, assigned_bot_id, status FROM %s WHERE order_id = $1 LIMIT 1", r.tableName)

	shopping := &domain.ShoppingList{
		OrderID: orderID,
	}
	var stops []byte
	var status string
	if err := r.db.QueryRowContext(ctx, query, orderID).Scan(&shopping.ID, &stops, &shopping.AssignedBotID, &status); err != nil {
		return nil, err
	}

	if err := json.Unmarshal(stops, &shopping.Stops); err != nil {
		return nil, err
	}

	shopping.Status = domain.ShoppingListStatus(status)

	return shopping, nil
}

func (r *ShoppingListRepository) Save(ctx context.Context, list *domain.ShoppingList) error {
	query := fmt.Sprintf("INSERT INTO %s (id, order_id, stops, assigned_bot_id, status) VALUES ($1, $2, $3, $4, $5)", r.tableName)

	stops, err := json.Marshal(list.Stops)
	if err != nil {
		return err
	}
	_, err = r.db.ExecContext(ctx, query, list.ID, list.OrderID, stops, list.AssignedBotID, list.Status.String())
	if err != nil {
		return err
	}

	return nil
}

func (r *ShoppingListRepository) Update(ctx context.Context, list *domain.ShoppingList) error {
	query := fmt.Sprintf("UPDATE %s SET stops = $2, assigned_bot_id = $3, status = $4 WHERE id = $1", r.tableName)

	stops, err := json.Marshal(list.Stops)
	if err != nil {
		return err
	}

	_, err = r.db.ExecContext(ctx, query, list.ID, stops, list.AssignedBotID, list.Status.String())
	if err != nil {
		return err
	}

	return nil
}
