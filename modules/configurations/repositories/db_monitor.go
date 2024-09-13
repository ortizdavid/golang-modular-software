package repositories

import (
	"context"
	"log"

	"gorm.io/gorm"
)

type TableInfo struct {
    TableName string `gorm:"column:TABLE_NAME"`
}

type DBMonitor struct {
	db *gorm.DB
}

func NewDBMonitor(db *gorm.DB) *DBMonitor {
	return &DBMonitor{
		db: db,
	}
}

func (monitor *DBMonitor) GetAllDBTables(ctx context.Context) ([]TableInfo, error) {
	var tables []TableInfo
	result := monitor.db.Raw("SELECT table_name FROM information_schema.tables WHERE table_schema = 'iguana_balancete_bna' AND table_type = 'BASE TABLE'").Scan(&tables)
	if result.Error != nil {
		return nil, result.Error
	}
	return tables, nil
}

func (monitor *DBMonitor) UpdateStatistics(ctx context.Context) {
	tables, _ := monitor.GetAllDBTables(ctx)
	for _, table := range tables {
		result := monitor.db.Exec(" ANALYZE TABLE "+table.TableName+";")
		if result.Error != nil {
			log.Fatal(result.Error)
			break
		}
	}
}