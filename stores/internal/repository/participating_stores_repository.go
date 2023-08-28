package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/hnamzian/go-mallbots/stores/internal/domain"
	"github.com/stackus/errors"
)

type ParticipatingStoreRepository struct {
	tableName string
	db        *sql.DB
}

func NewParticipatingStoreRepository(tableName string, db *sql.DB) ParticipatingStoreRepository {
	return ParticipatingStoreRepository{tableName: tableName, db: db}
}

func (r ParticipatingStoreRepository) FindAll(ctx context.Context) (stores []*domain.Store, err error) {
	query := fmt.Sprintf("SELECT id, name, location, participating FROM %s WHERE participating is true", r.tableName)

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, errors.Wrap(err, "querying participating stores")
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			err = errors.Wrap(err, "closing participating store rows")
		}
	}(rows)

	for rows.Next() {
		store := &domain.Store{}
		err := rows.Scan(&store.ID, &store.Name, &store.Location, &store.Participating)
		if err != nil {
			return nil, errors.Wrap(err, "scanning participating store")
		}

		stores = append(stores, store)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "finishing participating store rows")
	}

	return stores, nil
}
