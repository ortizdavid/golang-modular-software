package controllers

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/shared/entities"
)

// BaseRepository provides generic CRUD operations for entities.
type BaseRepository[T entities.Entity] struct {
	db *database.Database
}

// Create inserts a new entity into the database.
func (repo *BaseRepository[T]) Create(ctx context.Context, entity T) error {
	result := repo.db.WithContext(ctx).Create(&entity)
	return result.Error
}

// CreateBatch inserts multiple entities into the database within a transaction.
func (repo *BaseRepository[T]) CreateBatch(ctx context.Context, entities []T) error {
	tx := repo.db.Begin()  // Start a new transaction
	if tx.Error != nil {
		return tx.Error
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()  // Rollback on panic
		}
	}()
	result := tx.WithContext(ctx).Create(&entities)
	if result.Error != nil {
		tx.Rollback()  // Rollback on error
		return result.Error
	}
	if err := tx.Commit().Error; err != nil {
		return err
	}
	return nil
}

// Update saves changes to an existing entity in the database.
func (repo *BaseRepository[T]) Update(ctx context.Context, entity T) error {
	result := repo.db.WithContext(ctx).Save(&entity)
	return result.Error
}

// UpdateBatch updates multiple entities in the database within a transaction.
func (repo *BaseRepository[T]) UpdateBatch(ctx context.Context, entities []T) error {
	tx := repo.db.Begin()  // Start a new transaction
	if tx.Error != nil {
		return tx.Error
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()  // Rollback on panic
		}
	}()
	for _, entity := range entities {
		result := tx.WithContext(ctx).Save(&entity)
		if result.Error != nil {
			tx.Rollback()  // Rollback on error
			return result.Error
		}
	}
	if err := tx.Commit().Error; err != nil {
		return err
	}
	return nil
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

// FindAllLimit retrieves entities with pagination (limit and offset).
func (repo *BaseRepository[T]) FindAllLimit(ctx context.Context, limit int, offset int) ([]T, error) {
	var entities []T
	var entity T
	tableName := entity.TableName()  // Get table name for the entity
	result := repo.db.WithContext(ctx).Table(tableName).Limit(limit).Offset(offset).Find(&entities)
	return entities, result.Error
}

// FindAllOrdered retrieves entities ordered by a specific field.
func (repo *BaseRepository[T]) FindAllOrdered(ctx context.Context, fieldAnOrder string) ([]T, error) {
	var entities []T
	result := repo.db.WithContext(ctx).Order(fieldAnOrder).Find(&entities)
	return entities, result.Error
}

// FindById retrieves an entity by its ID.
func (repo *BaseRepository[T]) FindById(ctx context.Context, id int) (T, error) {
	var entity T
	result := repo.db.WithContext(ctx).First(&entity, id)
	return entity, result.Error
}

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

// ExistsField checks if an entity with a specific field value exists.
func (repo *BaseRepository[T]) ExistsField(ctx context.Context, field string, value interface{}) (bool, error) {
	var count int64
	var entity T
	tableName := entity.TableName()  // Get table name for the entity
	result := repo.db.WithContext(ctx).Table(tableName).Where(field+" = ?", value).Count(&count)
	if result.Error != nil {
		return false, result.Error
	}
	return count > 0, nil
}
