package main

import (
	"log"
	"os"
	"path/filepath"
	"github.com/ortizdavid/golang-modular-software/config"
	"gorm.io/gorm"
)


const (
	structureDir = "_structure"
	authDir = "authentication"
	configDir = "configuration"
	hrDir = "human-resources"
	customerDir = "customers"
)

// Execute a sql script located in a directory
func execDatabaseScript(db *gorm.DB, directory string, scriptFile string) {
	parentDir := "database"
	scriptDir := filepath.Join(parentDir, directory)
	scriptPath := filepath.Join(scriptDir, scriptFile)

	scriptContent, err := os.ReadFile(scriptPath)
	if err != nil {
		log.Fatalf("Failed to read script '%s': %v", scriptFile, err)
	}
	result := db.Exec(string(scriptContent))
	if result.Error != nil {
		log.Fatalf("Error executing script '%s': %v", scriptFile, result.Error)
	}
	log.Printf("Script '%s' executed successfully!\n", scriptPath)
}

// Create database schemas
func execCreateScemas(db *gorm.DB) {
	log.Printf("Executing schema creation scripts...")
	execDatabaseScript(db, structureDir, "schemas.sql")
}

// execute all authentication scripts
func execAuthenticationScripts(db *gorm.DB) {
	log.Printf("Executing authentication schema scripts...")
	execDatabaseScript(db, authDir, "tables.sql")
	execDatabaseScript(db, authDir, "views.sql")
}

// execute all configuration scripts
func execConfigurationScripts(db *gorm.DB) {
	log.Printf("Executing configurationn schema scripts...")
	execDatabaseScript(db, configDir, "tables.sql")
	execDatabaseScript(db, configDir, "views.sql")
}

// execute all human_resources scripts
func execHumanResourcesScripts(db *gorm.DB) {
	log.Printf("Executing human_resources schema scripts...")
	execDatabaseScript(db, hrDir, "tables.sql")
	execDatabaseScript(db, hrDir, "views.sql")
}

// execute all customers scripts
func execCustomerScripts(db *gorm.DB) {
	log.Printf("Executing customers schema scripts...")
	execDatabaseScript(db, customerDir, "tables.sql")
	execDatabaseScript(db, customerDir, "views.sql")
}


func main() {

	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to database")
		panic(err)
	}
	defer config.DisconnectDB(db)
	log.Printf("Connected to Database ...\n\n")

	log.Println("Executing database scripts...")
	execCreateScemas(db)
	execAuthenticationScripts(db)
	execConfigurationScripts(db)
	execHumanResourcesScripts(db)
	execCustomerScripts(db)
}