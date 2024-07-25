package database

import (
	"log"
	"os"
	"path/filepath"
	"gorm.io/gorm"
)


const (
	structureDir = "_structure"
	authDir = "authentication"
	configDir = "configuration"
	companyDir = "company"
	referenceDir = "reference"
)

// Execute a sql script located in a directory
func execDatabaseScript(db *gorm.DB, directory string, scriptFile string) {
	parentDir := "../sql"
	scriptDir := filepath.Join(parentDir, directory)
	scriptPath := filepath.Join(scriptDir, scriptFile)
	// Read script
	scriptContent, err := os.ReadFile(scriptPath)
	if err != nil {
		log.Fatalf("Failed to read script '%s': %v", scriptFile, err)
	}
	// start transaction
	tx := db.Begin()
	// execute scrpt content
	result := db.Exec(string(scriptContent))
	if result.Error != nil {
		tx.Rollback() //Rollback transaction
		log.Fatalf("Error executing script '%s': %v", scriptFile, result.Error)
	}
	//commit transaction
	commit := tx.Commit()
	if commit.Error != nil {
		log.Fatalf("Error committing transaction for script '%s': %v", scriptFile, commit.Error)
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

// execute all company scripts
func execCompanyScripts(db *gorm.DB) {
	log.Printf("Executing company schema scripts...")
	execDatabaseScript(db, companyDir, "tables.sql")
	execDatabaseScript(db, companyDir, "views.sql")
}

// execute all reference scripts
func execReferenceScripts(db *gorm.DB) {
	log.Printf("Executing reference schema scripts...")
	execDatabaseScript(db, referenceDir, "tables.sql")
	execDatabaseScript(db, referenceDir, "views.sql")
}


func InitDatabaseScripts(db *gorm.DB) {
	log.Printf("Connected to Database ...\n\n")

	log.Println("Executing database scripts...")
	execCreateScemas(db)
	execAuthenticationScripts(db)
	execConfigurationScripts(db)
	execCompanyScripts(db)
	execReferenceScripts(db)
}