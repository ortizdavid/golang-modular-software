package repositories

import "context"

// UpdateWhere Updates an entity from the database based on a specified field and its value.
func (repo *BaseRepository[T]) UpdateWhere(ctx context.Context, field string, value interface{}, updates map[string]interface{}) error {
	var entity T
	result := repo.db.WithContext(ctx).Model(&entity).Where(field+" = ?", value).Updates(updates)
	repo.setAffectedRows(result.RowsAffected)
	return result.Error
}
