package main

import (
	"fmt"
	"os"

	"github.com/ortizdavid/golang-modular-software/config"
	"gorm.io/gorm"
)


const (
	parentDir = "database"
	structureDir = "_structure"
	authDir = "authentication"
	configDir = "configuration"
	hrDir = "human-resources"
	customerDir = "customers"
)

func execDatabaseScript(db *gorm.DB, directory string, scriptFile string) {
	scriptDir := parentDir +"/"+ directory + "/"
	scriptContent, err := os.ReadFile(scriptDir + scriptFile)
	if err != nil {
		panic(err)
	}
	result := db.Exec(string(scriptContent))
	if result.Error != nil {
		panic(result.Error)
	}
	fmt.Printf("Script '%s%s' executed successfully!\n", scriptDir, scriptFile)
}

func CreateScemas(db *gorm.DB) {
	execDatabaseScript(db, structureDir, "schemas.sql")
}

func execAuthenticationScripts(db *gorm.DB) {
	execDatabaseScript(db, authDir, "tables.sql")
	execDatabaseScript(db, authDir, "views.sql")
}

func execConfigurationScripts(db *gorm.DB) {
	execDatabaseScript(db, configDir, "tables.sql")
	execDatabaseScript(db, configDir, "views.sql")
}

func execHumanResourcesScripts(db *gorm.DB) {
	execDatabaseScript(db, hrDir, "tables.sql")
	execDatabaseScript(db, hrDir, "views.sql")
}

func execCustomerScripts(db *gorm.DB) {
	execDatabaseScript(db, customerDir, "tables.sql")
	execDatabaseScript(db, customerDir, "views.sql")
}


func main() {

	db, err := config.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer config.DisconnectDB(db)
	fmt.Println("Connected to Database ...")

	//Execute Scripts
	CreateScemas(db)
	execAuthenticationScripts(db)
	execConfigurationScripts(db)
	execHumanResourcesScripts(db)
	execCustomerScripts(db)
}