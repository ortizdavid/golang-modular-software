package database

import (
	"context"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

func NewDatabase(db *gorm.DB) *Database {
	return &Database{
		DB: db,
	}
}

// BeginTx starts a transaction
func (d *Database) beginTx(ctx context.Context) (*Database, error) {
	tx := d.DB.WithContext(ctx).Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &Database{DB: tx}, nil
}

// CommitTx commits the transaction
func (d *Database) commitTx() error {
	return d.DB.Commit().Error
}

// RollbackTx rolls back the transaction
func (d *Database) rollbackTx() error {
	return d.DB.Rollback().Error
}

// WithTx runs the provided function within a transaction
// Rollback if an error occurred
func (d *Database) WithTx(ctx context.Context, fn func(tx *Database) error) error{
	tx, err := d.beginTx(ctx)
	if err != nil {
		return err
	}
	err = fn(tx)
	if err != nil {
		if rbErr := tx.rollbackTx(); rbErr != nil {
			return rbErr // rollback Error
		}
		return err // function error
	}
	return tx.commitTx()
}

