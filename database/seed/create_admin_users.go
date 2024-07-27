package database

import (
	"context"
	"fmt"
	"log"
	"github.com/ortizdavid/go-nopain/encryption"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/entities"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/services"
)

func CreateAdminUsers(db *database.Database) error {
	userService := services.NewUserService(db)
	ctx := context.Background()
	// Entities ----------------------------------
	user1 := entities.CreateUserRequest{
		UserName: "admin01",
		RoleId:   1,
		Email:    "admin01@gmail.com",
		Password: "admin123",
	}
	user2 := entities.CreateUserRequest{
		UserName: "admin02",
		RoleId:   2,
		Email:    "admin02@gmail.com",
		Password: encryption.HashPassword("admin123"),
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