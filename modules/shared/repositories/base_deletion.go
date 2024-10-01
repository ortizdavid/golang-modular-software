package repositories

import "context"

// DeleteById  deletes an entity from the database using its primary key or unique identifier.
func (repo *BaseRepository[T]) DeleteByUniqueId(ctx context.Context, uniqueId string) error {
	var entity T
	result := repo.db.WithContext(ctx).Where("unique_id = ?", uniqueId).Delete(&entity)
	return result.Error
}

// DeleteByUniqueId  deletes an entity from the database using its unique_id field.
func (repo *BaseRepository[T]) DeleteById(ctx context.Context, id interface{}) error {
	var entity T
	result := repo.db.WithContext(ctx).First(&entity, id).Delete(&entity)
	return result.Error
}

// DeleteWhere deletes an entity from the database based on a specified field and its value.
func (repo *BaseRepository[T]) DeleteWhere(ctx context.Context, field string, value interface{}) error {
	var entity T
	result := repo.db.WithContext(ctx).Where(field+" = ?", value).Delete(&entity)
	return result.Error
}
