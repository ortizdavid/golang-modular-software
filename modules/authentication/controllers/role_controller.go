package controllers

import (
	"github.com/gofiber/fiber/v2"
	authentication "github.com/ortizdavid/golang-modular-software/modules/authentication/repositories"
	configurations "github.com/ortizdavid/golang-modular-software/modules/configurations/repositores"
)

type RoleController struct {
}


func (role RoleController) Routes(router *fiber.App) {

}


func (RoleController) index(c *fiber.Ctx) error {
	basicConfig, _ := configurations.GetBasicConfiguration()
	return c.Render("users/roles/index", fiber.Map{
		"Title":       "User Details",
		"LoggedUser":  authentication.GetLoggedUser(c),
		"BasicConfig": basicConfig,
	})
}


func (RoleController) addForm(c *fiber.Ctx) error {
	roles, _ := models.RoleModel{}.FindAll()
	basicConfig, _ := configurations.GetBasicConfiguration()
	return c.Render("users/roles/add", fiber.Map{
		"Title":       "Add Role",
		"Roles":       roles,
		"LoggedUser":  authentication.GetLoggedUser(c),
		"BasicConfig": basicConfig,
	})
}