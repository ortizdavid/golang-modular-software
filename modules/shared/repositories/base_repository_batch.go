package repositories

import "context"

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
	
	// Track affected rows if needed
	repo.setAffectedRows(result.RowsAffected)

	if err := tx.Commit().Error; err != nil {
		return err
	}
	return nil
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

	var totalAffected int64
	for _, entity := range entities {
		result := tx.WithContext(ctx).Save(&entity)
		if result.Error != nil {
			tx.Rollback()  // Rollback on error
			return result.Error
		}
		totalAffected += result.RowsAffected
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}
	// Optionally set affected rows
	repo.setAffectedRows(totalAffected)
	return nil
}

