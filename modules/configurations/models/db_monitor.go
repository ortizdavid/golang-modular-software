package models

import (
	"log"
	"github.com/ortizdavid/golang-modular-software/config"
)


type TableInfo struct {
    TableName string `gorm:"column:TABLE_NAME"`
}

type DBMonitor struct {
}


func (monitor DBMonitor) GetAllDBTables() ([]TableInfo, error) {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	var tables []TableInfo
	result := db.Raw("SELECT table_name FROM information_schema.tables WHERE table_schema = 'iguana_balancete_bna' AND table_type = 'BASE TABLE'").Scan(&tables)
	if result.Error != nil {
		return nil, result.Error
	}
	return tables, nil
}


func (monitor DBMonitor) UpdateStatistics() {
	db, _ := config.ConnectDB()
	defer config.DisconnectDB(db)
	tables, _ := monitor.GetAllDBTables()
	for _, table := range tables {
		result := db.Exec(" ANALYZE TABLE "+table.TableName+";")
		if result.Error != nil {
			log.Fatal(result.Error)
			break
		}
	}
}