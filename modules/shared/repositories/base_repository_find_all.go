package repositories

import (
	"context"
	"fmt"
)

// FindAllLimit retrieves entities with pagination (limit and offset).
func (repo *BaseRepository[T]) FindAllLimit(ctx context.Context, limit int, offset int) ([]T, error) {
	var entities []T
	result := repo.db.WithContext(ctx).Model(&entities).Limit(limit).Offset(offset).Find(&entities)
	repo.setAffectedRows(result.RowsAffected)
	return entities, result.Error
}

// FindAllOrdered retrieves entities ordered by a specific field.
func (repo *BaseRepository[T]) FindAllOrdered(ctx context.Context, fieldAnOrder string) ([]T, error) {
	var entities []T
	result := repo.db.WithContext(ctx).Order(fieldAnOrder).Find(&entities)
	repo.setAffectedRows(result.RowsAffected)
	return entities, result.Error
}

// FindAllNotIn retrieves entities where a specific field's value is not in the provided list.
func (repo *BaseRepository[T]) FindAllNotIn(ctx context.Context, fieldName string, values ...[]interface{}) ([]T, error) {
	var entities []T
	result := repo.db.WithContext(ctx).Model(&entities).Where(fmt.Sprintf("%s NOT IN (?)", fieldName), values).Find(&entities)
	repo.setAffectedRows(result.RowsAffected)
	return entities, result.Error
}

