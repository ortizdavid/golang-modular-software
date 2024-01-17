package main

import (
	"fmt"
	"log"
	"time"
	"github.com/ortizdavid/go-nopain/encryption"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/entities"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/models"
)

func main() {
	var userName, password, roleName string
	var roleId int
	var roleModel models.RoleModel
	var userModel models.UserModel

	fmt.Println("Create Admin Users")
	fmt.Print("Select Roles: \n[1]-Super Admin \n[2]-Admin\n")
	fmt.Print("Select: ")
	fmt.Scanln(&roleId)
	fmt.Print("\nUsername: ")
	fmt.Scanln(&userName)
	fmt.Print("Password: ")
	fmt.Scanln(&password)

	if roleId < 1 || roleId > 2 {
		roleId = 2
	}
	role, _ := roleModel.FindById(roleId)
	roleName = role.RoleName

	user := entities.User{
		UserId:    0,
		RoleId:    roleId,
		UserName:  userName,
		Password:  encryption.HashPassword(password),
		Active:    "Yes",
		Image:     "",
		Token:     "",
		UniqueId:  encryption.GenerateUUID(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	_, err := userModel.Create(user)
	if err != nil {
		log.Printf("Error while creating User '%s' -> %v", userName, err.Error())
	} else {
		log.Printf("User '%s' Created with role '%s'", userName, roleName)
	}
}