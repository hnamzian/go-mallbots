package repository

import (
	"context"
	"database/sql"
	"github.com/stackus/errors"
	"fmt"

	"github.com/hnamzian/go-mallbots/stores/internal/domain"
)

type StoreRepository struct {
	tableName string
	db        *sql.DB
}

func NewStore(tableName string, db *sql.DB) *StoreRepository {
	return &StoreRepository{
		tableName: tableName,
		db:        db,
	}
}

func (r StoreRepository) Save(ctx context.Context, store *domain.Store) error {
	query := fmt.Sprintf("INSERT INTO %s (id, name, location, participating) VALUES ($1, $2, $3, $4)", r.tableName)

	_, err := r.db.ExecContext(ctx, query, store.ID, store.Name, store.Location, store.Participating)
	if err != nil {
		return errors.Wrap(err, "failed to save store")
	}

	return nil
}

func (r StoreRepository) Get(ctx context.Context, id string) (*domain.Store, error) {
	query := fmt.Sprintf("SELECT id, name, location, participating FROM %s WHERE id = $1", r.tableName)

	store := &domain.Store{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(&store.ID, &store.Name, &store.Location, &store.Participating)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find store")
	}
	return store, nil
}

func (r StoreRepository) GetAll(ctx context.Context) ([]*domain.Store, error) {
	query := fmt.Sprintf("SELECT id, name, location, participating FROM %s ORDER BY id ASC", r.tableName)

	rows, err := r.db.QueryContext(ctx, query)
	if err!= nil {
        return nil, errors.Wrap(err, "failed to find stores")
    }

	defer func(rows *sql.Rows) {
		if err := rows.Close(); err != nil {
			err = errors.Wrap(err, "failed to close rows")
        }
	}(rows)

	stores := []*domain.Store{}
	for rows.Next() {
		var store *domain.Store

		if err = rows.Scan(&store.ID, &store.Name, &store.Location, &store.Participating); err!= nil {
			return nil, errors.Wrap(err, "failed to scan store")
        }

		stores = append(stores, store)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "finishing store rows")
	}

	return stores, nil
}

func (r StoreRepository) GetParticipatingStore(ctx context.Context, store *domain.Store) ([]*domain.Store ,error) {
	query := fmt.Sprintf("SELECT id, name, location, participating FROM %s WHERE participating is true", r.tableName)

	rows, err := r.db.QueryContext(ctx, query)
	if err!= nil {
        return nil, errors.Wrap(err, "failed to find participating stores")
    }

	defer func(rows *sql.Rows) {
        if err := rows.Close(); err!= nil {
            err = errors.Wrap(err, "failed to close rows")
        }
    }(rows)

	stores := []*domain.Store{}
	for rows.Next() {
		var store *domain.Store

        if err = rows.Scan(&store.ID, &store.Name, &store.Location, &store.Participating); err!= nil {
            return nil, errors.Wrap(err, "failed to scan store")
        }

        stores = append(stores, store)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "finishing participating store rows")
	}

	return stores, nil
} 

func (r StoreRepository) Update(ctx context.Context, store *domain.Store) error {
	query := fmt.Sprintf("UPDATE %s SET name = $2, location = $3, participating = $3 WHERE id = $1", r.tableName)

	_, err := r.db.ExecContext(ctx, query, store.ID, store.Name, store.Location, store.Participating)
	if err!= nil {
        return errors.Wrap(err, "failed to update store")
    }

	return nil
}

func (r StoreRepository) Delete(ctx context.Context, store *domain.Store) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", r.tableName)

    _, err := r.db.ExecContext(ctx, query, store.ID)
    if err!= nil {
        return errors.Wrap(err, "failed to delete store")
    }

    return nil
}