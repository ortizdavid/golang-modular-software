package database

import (
	"context"
	"fmt"
	"log"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/entities"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/services"
)

func CreateAdminUsers(db *database.Database) error {
	userService := services.NewUserService(db)
	ctx := context.Background()
	// Entities ----------------------------------
	user1 := entities.CreateUserRequest{
		UserName: "sup-admin01",
		RoleId:   entities.RoleSuperAdmin.Id,
		Email:    "sup-admin01@gmail.com",
		Password: "12345678",
	}
	user2 := entities.CreateUserRequest{
		UserName: "admin01",
		RoleId:   entities.RoleAdmin.Id,
		Email:    "admin01@gmail.com",
		Password: "12345678",
	}
	// create users -----------------------------------------------
	err1 := userService.CreateUser(ctx, user1)
	err2 := userService.CreateUser(ctx, user2)
	if err1 != nil || err2 != nil{
		return fmt.Errorf("error while creating admin users: \n%v\n%v", err1.Error(), err2.Error())
	}
	log.Println("Users Created with roles")
	return nil
}