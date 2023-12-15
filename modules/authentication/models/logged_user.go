package models

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/config"
	entities "github.com/ortizdavid/golang-modular-software/modules/users/entities"
	models "github.com/ortizdavid/golang-modular-software/modules/users/models"
	"github.com/ortizdavid/go-nopain/conversion"
)


func GetLoggedUser(ctx *fiber.Ctx) entities.UserData {
	store := config.GetSessionStore()
	session, _ := store.Get(ctx)
	userName := conversion.ConvertAnyToString(session.Get("username"))
	password := conversion.ConvertAnyToString(session.Get("password"))
	user, err := models.UserModel{}.GetByUserNameAndPassword(userName, password)
	if err != nil {
		log.Fatal(err)
		return entities.UserData{}
	}
	return user
}


func IsUserAuthenticated(ctx *fiber.Ctx) bool {
	loggedUser := GetLoggedUser(ctx)
	if loggedUser.UserId == 0 && loggedUser.RoleId == 0 {
		return false
	}
	return true
}



func IsUserAdmin(ctx *fiber.Ctx) bool {
	loggedUser := GetLoggedUser(ctx)
	return loggedUser.RoleCode == "admin"
}
