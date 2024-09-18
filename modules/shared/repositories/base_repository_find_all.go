package repositories

import "context"

// FindAllLimit retrieves entities with pagination (limit and offset).
func (repo *BaseRepository[T]) FindAllLimit(ctx context.Context, limit int, offset int) ([]T, error) {
	var entities []T
	result := repo.db.WithContext(ctx).Model(&entities).Limit(limit).Offset(offset).Find(&entities)
	return entities, result.Error
}

// FindAllOrdered retrieves entities ordered by a specific field.
func (repo *BaseRepository[T]) FindAllOrdered(ctx context.Context, fieldAnOrder string) ([]T, error) {
	var entities []T
	result := repo.db.WithContext(ctx).Order(fieldAnOrder).Find(&entities)
	return entities, result.Error
}

