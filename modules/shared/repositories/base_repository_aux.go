package repositories

// SetLastInsertId sets the last inserted ID safely with a mutex.
func (repo *BaseRepository[T]) SetLastInsertId(id int64) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.lastInsertId = id
}

// GetLastInsertId retrieves the last inserted ID safely with a mutex.
func (repo *BaseRepository[T]) GetLastInsertId() int64 {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	return repo.lastInsertId
}

// setAffectedRows sets the number of affected rows safely with a mutex.
func (repo *BaseRepository[T]) setAffectedRows(rows int64) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.affectedRows = rows
}

// GetAffectedRows retrieves the number of affected rows safely with a mutex.
func (repo *BaseRepository[T]) GetAffectedRows() int64 {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	return repo.affectedRows
}
