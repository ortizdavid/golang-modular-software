package repositories

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/database"
)

// This method can be used in services and other packages
func (repo *BaseRepository[T]) BeginTransaction(ctx context.Context) (*database.Database, error) {
	return repo.db.BeginTx(ctx)
}

// WithTransaction executes the provided function within a transaction.
// This method can be used in services and other packages that require transactional integrity
func (repo *BaseRepository[T]) WithTransaction(ctx context.Context, fn func(tx *database.Database) error) error {
	return repo.db.WithTx(ctx, fn)
}