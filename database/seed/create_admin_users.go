package database

import (
	"context"
	"log"
	"time"
	"github.com/ortizdavid/go-nopain/encryption"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/entities"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/repositories"
	"gorm.io/gorm"
)

func CreateAdminUsers(db *gorm.DB) {
	repository := repositories.NewUserRepository(db)
	ctx := context.Background()
	users := []entities.User{
		{
			UserId:   0,
			UserName: "admin01",
			Email:    "admin01@gmail.com",
			Password: encryption.HashPassword("admin123"),
			IsActive: true,
			Image: "",
			Token: encryption.GenerateRandomToken(),
			UniqueId: encryption.GenerateUUID(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			UserId:   0,
			UserName: "admin01",
			Email:    "admin01@gmail.com",
			Password: encryption.HashPassword("admin123"),
			IsActive: true,
			Image: "",
			Token: encryption.GenerateRandomToken(),
			UniqueId: encryption.GenerateUUID(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	err := repository.CreateBatch(ctx, users)
	if err != nil {
		log.Printf("Error while creating users: %v", err.Error())
	} else {
		log.Println("Users Created with roles")
	}
}