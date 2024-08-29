package database

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"github.com/ortizdavid/golang-modular-software/database"
)

const (
	structureDir   = "_structure"
	authDir        = "authentication"
	configDir      = "configurations"
	companyDir     = "company"
	referenceDir   = "references"
	employeeDir    = "employees"
	sqlScriptsBase = "./database/sql"
)

// execDatabaseScript executes a SQL script located in the specified directory.
func execDatabaseScript(db *database.Database, directory, scriptFile string) error {
	scriptPath := filepath.Join(sqlScriptsBase, directory, scriptFile)
	
	scriptContent, err := os.ReadFile(scriptPath)
	if err != nil {
		return fmt.Errorf("failed to read script '%s': %w", scriptFile, err)
	}

	tx := db.Begin()
	if err := tx.Error; err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	if err := tx.Exec(string(scriptContent)).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("error executing script '%s': %w", scriptFile, err)
	}

	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("error committing transaction for script '%s': %w", scriptFile, err)
	}

	log.Printf("Script '%s' executed successfully!\n", scriptPath)
	return nil
}

// execCreateSchemas creates database schemas.
func execCreateSchemas(db *database.Database) error {
	log.Println("Executing schema creation scripts...")
	return execDatabaseScript(db, structureDir, "schemas.sql")
}

// execCreateExtensions creates database extensions.
func execCreateExtensions(db *database.Database) error {
	log.Println("Executing extensions creation scripts...")
	return execDatabaseScript(db, structureDir, "extensions.sql")
}


// execAuthenticationScripts executes all authentication scripts.
func execAuthenticationScripts(db *database.Database) error {
	log.Println("Executing authentication schema scripts...")
	if err := execDatabaseScript(db, authDir, "tables.sql"); err != nil {
		return err
	}
	if err := execDatabaseScript(db, authDir, "triggers.sql"); err != nil {
		return err
	}
	if err := execDatabaseScript(db, authDir, "procedures.sql"); err != nil {
		return err
	}
	return execDatabaseScript(db, authDir, "views.sql")
}

// execConfigurationScripts executes all configuration scripts.
func execConfigurationScripts(db *database.Database) error {
	log.Println("Executing configuration schema scripts...")
	if err := execDatabaseScript(db, configDir, "tables.sql"); err != nil {
		return err
	}
	if err := execDatabaseScript(db, configDir, "triggers.sql"); err != nil {
		return err
	}
	if err := execDatabaseScript(db, configDir, "procedures.sql"); err != nil {
		return err
	}
	return execDatabaseScript(db, configDir, "views.sql")
}

// execCompanyScripts executes all company scripts.
func execCompanyScripts(db *database.Database) error {
	log.Println("Executing company schema scripts...")
	if err := execDatabaseScript(db, companyDir, "tables.sql"); err != nil {
		return err
	}
	if err := execDatabaseScript(db, companyDir, "triggers.sql"); err != nil {
		return err
	}
	if err := execDatabaseScript(db, companyDir, "procedures.sql"); err != nil {
		return err
	}
	return execDatabaseScript(db, companyDir, "views.sql")
}

// execEmployeeScripts executes all company scripts.
func execEmployeeScripts(db *database.Database) error {
	log.Println("Executing employee schema scripts...")
	if err := execDatabaseScript(db, employeeDir, "tables.sql"); err != nil {
		return err
	}
	if err := execDatabaseScript(db, employeeDir, "triggers.sql"); err != nil {
		return err
	}
	if err := execDatabaseScript(db, employeeDir, "procedures.sql"); err != nil {
		return err
	}
	return execDatabaseScript(db, employeeDir, "views.sql")
}


// execReferenceScripts executes all reference scripts.
func execReferenceScripts(db *database.Database) error {
	log.Println("Executing reference schema scripts...")
	if err := execDatabaseScript(db, referenceDir, "tables.sql"); err != nil {
		return err
	}
	if err := execDatabaseScript(db, referenceDir, "triggers.sql"); err != nil {
		return err
	}
	if err := execDatabaseScript(db, referenceDir, "procedures.sql"); err != nil {
		return err
	}
	return execDatabaseScript(db, referenceDir, "views.sql")
}

// InitDatabaseScripts initializes and executes all database scripts.
func InitDatabaseScripts(db *database.Database) error {
	log.Println("Connected to Database...")
	log.Println("Executing database scripts...")

	if err := execCreateSchemas(db); err != nil {
		return err
	}
	if err := execCreateExtensions(db); err != nil {
		return err
	}
	if err := execAuthenticationScripts(db); err != nil {
		return err
	}
	if err := execConfigurationScripts(db); err != nil {
		return err
	}
	if err := execCompanyScripts(db); err != nil {
		return err
	}
	if err := execEmployeeScripts(db); err != nil {
		return err
	}
	return execReferenceScripts(db)
}
