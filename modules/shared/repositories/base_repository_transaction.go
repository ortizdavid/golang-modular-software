package repositories

import (
	"context"
	"gorm.io/gorm"
)

func (repo *BaseRepository[T]) BeginTransaction(ctx context.Context) (*gorm.DB, error) {
	return repo.db.BeginTx(ctx)
}
