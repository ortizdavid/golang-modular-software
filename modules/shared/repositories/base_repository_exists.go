package repositories

import "context"

// ExistsField checks if an entity with a specific field value exists.
func (repo *BaseRepository[T]) ExistsField(ctx context.Context, field string, value interface{}) (bool, error) {
	var count int64
	var entity T
	result := repo.db.WithContext(ctx).Model(&entity).Where(field+" = ?", value).Count(&count)
	if result.Error != nil {
		return false, result.Error
	}
	return count > 0, nil
}
