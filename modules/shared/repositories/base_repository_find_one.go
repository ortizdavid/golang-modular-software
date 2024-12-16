package repositories

import "context"

// FindByUniqueId retrieves an entity by its unique ID.
func (repo *BaseRepository[T]) FindByUniqueId(ctx context.Context, uniqueId string) (T, error) {
	var entity T
	result := repo.db.WithContext(ctx).Where("unique_id = ?", uniqueId).First(&entity)
	return entity, result.Error
}

// FindByField retrieves an entity by a specific field and value.
func (repo *BaseRepository[T]) FindByField(ctx context.Context, field string, value interface{}) (T, error) {
	var entity T
	result := repo.db.WithContext(ctx).Where(field+" = ?", value).First(&entity)
	repo.setAffectedRows(result.RowsAffected)
	return entity, result.Error
}

// FindFirst retrieves the first entity found in the database.
func (repo *BaseRepository[T]) FindFirst(ctx context.Context) (T, error) {
	var entity T
	result := repo.db.WithContext(ctx).First(&entity)
	return entity, result.Error
}

// FindLast retrieves the last entity found in the database.
func (repo *BaseRepository[T]) FindLast(ctx context.Context) (T, error) {
	var entity T
	result := repo.db.WithContext(ctx).Last(&entity)
	return entity, result.Error
}
