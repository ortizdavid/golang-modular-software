package database

import (
	"log"
	"time"

	"github.com/ortizdavid/golang-modular-software/common/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConnection struct {
	DB 			*gorm.DB
	url 		string
	poolConfig 	ConnPoolConfig
}

type ConnPoolConfig struct {
	MaxIdleConns	int
	MaxOpenConns	int
	MaxLifeTime		time.Duration
	MaxIdleTime		time.Duration 	
}

// NewDBConnection creates a new database connection using the database URL
func NewDBConnection(dbURL string) (*DBConnection, error) {
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	// Create DBConnection struct with default pool settings
	dbConn := &DBConnection{
		DB: db,
		url: dbURL,
		poolConfig: ConnPoolConfig{
			MaxIdleConns: 	10,
			MaxOpenConns: 	100,
			MaxLifeTime:	time.Hour,
			MaxIdleTime: 	30 * time.Minute,
		}, 
	}
	dbConn.configurePool()
	return dbConn, nil
}

// NewDBConnectionFromEnvVar creates a new database connection using environment variable
func NewDBConnectionFromEnv(dbEnvVarName string) (*DBConnection, error) {
	dbConn, err := NewDBConnection(config.GetEnv(dbEnvVarName))
	if err != nil {
		return nil, err
	}
	return dbConn, nil
}

// Close database connection
func (dbConn *DBConnection) Close() {
	if dbConn == nil {
		return
	}
	sqlDB, err := dbConn.DB.DB()
	if err != nil {
		log.Fatalf("Failed to disconnect DB: %v", err)
	}
	sqlDB.Close()
}

// SetConnectionPool allows modifying the connection pool settings
func (dbConn *DBConnection) SetConnectionPool(poolConfig ConnPoolConfig) {
	dbConn.poolConfig = poolConfig
	dbConn.configurePool()
}

// configurePool applies the connection pool settings to the database connection
func (dbConn *DBConnection) configurePool() {
	sqlDB, err := dbConn.DB.DB()
	if err != nil {
		log.Fatalf("Failed to apply connection pool settings: %v", err)
	}
	sqlDB.SetMaxIdleConns(dbConn.poolConfig.MaxIdleConns)
	sqlDB.SetMaxOpenConns(dbConn.poolConfig.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(dbConn.poolConfig.MaxLifeTime)
	sqlDB.SetConnMaxIdleTime(dbConn.poolConfig.MaxIdleTime)
}
