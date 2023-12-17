package controllers

import (
	"github.com/gofiber/fiber/v2"
	authentication "github.com/ortizdavid/golang-modular-software/modules/authentication/models"
	configurations "github.com/ortizdavid/golang-modular-software/modules/configurations/models"
	models "github.com/ortizdavid/golang-modular-software/modules/authentication/models"
)

type RoleController struct {
}


func (role RoleController) Routes(router *fiber.App) {

}


func (RoleController) index(ctx *fiber.Ctx) error {
	basicConfig, _ := configurations.GetBasicConfiguration()
	return ctx.Render("users/roles/index", fiber.Map{
		"Title":       "User Details",
		"LoggedUser":  authentication.GetLoggedUser(ctx),
		"BasicConfig": basicConfig,
	})
}


func (RoleController) addForm(ctx *fiber.Ctx) error {
	roles, _ := models.RoleModel{}.FindAll()
	basicConfig, _ := configurations.GetBasicConfiguration()
	return ctx.Render("users/roles/add", fiber.Map{
		"Title":       "Add Role",
		"Roles":       roles,
		"LoggedUser":  authentication.GetLoggedUser(ctx),
		"BasicConfig": basicConfig,
	})
}