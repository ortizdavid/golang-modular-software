package repositories

import "context"

// UpdateWhere Updates an entity from the database based on a specified field and its value.
func (repo *BaseRepository[T]) UpdateWhere(ctx context.Context, field string, value interface{}, updates map[string]interface{}) error {
	result := repo.db.WithContext(ctx).Model(new(T)).Where(field+" = ?", value).Updates(updates)
	return result.Error
}
