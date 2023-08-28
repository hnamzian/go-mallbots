package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/hnamzian/go-mallbots/stores/internal/domain"
	"github.com/stackus/errors"
)

type ProductRepository struct {
	tableName string
	db        *sql.DB
}

func NewProductRepository(tableName string, db *sql.DB) *ProductRepository {
	return &ProductRepository{
		tableName: tableName,
		db:        db,
	}
}

func (r *ProductRepository) Save(ctx context.Context, product *domain.Product) error {
	query := fmt.Sprintf("INSERT INTO %s (id, store_id, name, description, price, sku) VALUES ($1, $2, $3, $4, $5, $6)", r.tableName)

	_, err := r.db.ExecContext(ctx, query, product.ID, product.StoreID, product.Name, product.Description, product.Price, product.SKU)
	if err!= nil {
		return errors.Wrap(err, "failed to save product")
	}
	return nil
}

func (r *ProductRepository) Get(ctx context.Context, id string) (*domain.Product, error) {
	query := fmt.Sprintf("SELECT store_id, name, description, price, sku FROM %s WHERE id = $1", r.tableName)

	product := &domain.Product{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(&product.StoreID, &product.Name, &product.Description, &product.Price, &product.SKU)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get product")
    }

	return product, nil
}
func (r *ProductRepository) GetCatalog(ctx context.Context, storeID string) ([]*domain.Product, error) {
	query := fmt.Sprintf("SELECT id, name, description, price, sku FROM %s WHERE store_id = $1", r.tableName)

	rows, err := r.db.QueryContext(ctx, query, storeID)
	if err!= nil {
        return nil, errors.Wrap(err, "failed to get catalog")
    }

	defer func(rows *sql.Rows) {
		if err = rows.Close(); err!= nil {
			err = errors.Wrap(err, "failed to close catalog rows")
		}
	}(rows)

	products := []*domain.Product{}
	for rows.Next() {
		product := &domain.Product{}
        err = rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.SKU)
        if err!= nil {
            return nil, errors.Wrap(err, "failed to scan catalog row")
        }
        products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "finishing products rows")
	}

	return products, nil
}
func (r *ProductRepository) Delete(ctx context.Context, id string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", r.tableName)

    _, err := r.db.ExecContext(ctx, query, id)
    if err!= nil {
        return errors.Wrap(err, "failed to delete product")
    }

    return nil
}
