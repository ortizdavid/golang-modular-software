package repositories

import "context"

// Count retrieves the total number of entities in the database.
func (repo *BaseRepository[T]) Count(ctx context.Context) (int64, error) {
	var count int64
	var entity T
	tableName := entity.TableName()  // Get table name for the entity
	result := repo.db.WithContext(ctx).Table(tableName).Count(&count)
	return count, result.Error
}

// CountWhere retrieves the count of entities that match a specific field and value.
func (repo *BaseRepository[T]) CountWhere(ctx context.Context, field string, value interface{}) (int64, error) {
	var count int64
	var entity T
	tableName := entity.TableName()  // Get table name for the entity
	result := repo.db.WithContext(ctx).Table(tableName).Where(field+" = ?", value).Count(&count)
	return count, result.Error
}