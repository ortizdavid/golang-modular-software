package repositories

func (repo *BaseRepository[T]) SetLastInsertId(id int64) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.LastInsertId = id
}
