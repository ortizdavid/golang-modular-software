package repositories

import (
	"context"
	"sync"
	"github.com/ortizdavid/golang-modular-software/database"
)

// BaseRepository provides generic CRUD operations for entities.
type BaseRepository[T any] struct {
	db *database.Database
	LastInsertId int64
	mu sync.Mutex
}

func NewBaseRepository[T any](db *database.Database) *BaseRepository[T] {
	return &BaseRepository[T]{
		db: db,
	}
}

// Create inserts a new entity into the database.
func (repo *BaseRepository[T]) Create(ctx context.Context, entity T) error {
	result := repo.db.WithContext(ctx).Create(&entity)
	return result.Error
}

// Update saves changes to an existing entity in the database.
func (repo *BaseRepository[T]) Update(ctx context.Context, entity T) error {
	result := repo.db.WithContext(ctx).Save(&entity)
	return result.Error
}

// Delete removes an entity from the database.
func (repo *BaseRepository[T]) Delete(ctx context.Context, entity T) error {
	result := repo.db.WithContext(ctx).Delete(&entity)
	return result.Error
}

// FindAll retrieves all entities from the database.
func (repo *BaseRepository[T]) FindAll(ctx context.Context) ([]T, error) {
	var entities []T
	result := repo.db.WithContext(ctx).Find(&entities)
	return entities, result.Error
}

// FindById retrieves an entity by its ID.
func (repo *BaseRepository[T]) FindById(ctx context.Context, id interface{}) (T, error) {
	var entity T
	result := repo.db.WithContext(ctx).First(&entity, id)
	return entity, result.Error
}
