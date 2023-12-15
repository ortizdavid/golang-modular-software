package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/config"
	authentication "github.com/ortizdavid/golang-modular-software/modules/authentication/models"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/models"
	"github.com/ortizdavid/go-nopain/conversion"
)

type BasicConfigController struct {
}


func (configuration BasicConfigController) Routes(router *fiber.App) {
	group := router.Group("/basic-configurations")
	group.Get("", configuration.index)
	group.Get("/edit", configuration.editForm)
	group.Post("/edit", configuration.edit)
}


func (BasicConfigController) index(ctx *fiber.Ctx) error {
	configuration, _ := models.GetBasicConfiguration()
	return ctx.Render("configurations/email/index", fiber.Map{
		"Title": "Email Configurations",
		"BasicConfiguration": configuration,
		"LoggedUser": authentication.GetLoggedUser(ctx),
	})
}


func (BasicConfigController) editForm(ctx *fiber.Ctx) error {
	configuration, _ := models.GetBasicConfiguration()
	return ctx.Render("configurations/email/edit", fiber.Map{
		"Title": "Edit Basic Configuration",
		"BasicConfiguaration": configuration,
		"LoggedUser": authentication.GetLoggedUser(ctx),
	})
}


func (BasicConfigController) edit(ctx *fiber.Ctx) error {
	companyName := ctx.FormValue("companyName")
	companyAcronym := ctx.FormValue("companyAcronym")
	numOfRecordPerPage := ctx.FormValue("numOfRecordPerPage")
	loggedUser := authentication.GetLoggedUser(ctx)

	var configurationModel models.BasicConfigurationModel
	configuration, _ := models.GetBasicConfiguration()
	configuration.CompanyName = companyName
	configuration.CompanyAcronym = companyAcronym
	configuration.NumOfRecordsPerPage = conversion.StringToInt(numOfRecordPerPage)
	_, err := configurationModel.Update(configuration)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	loggerConfiguration.Info(fmt.Sprintf("User '%s' updated basic configurations!", loggedUser.UserName), config.LogRequestPath(ctx))
	return ctx.Redirect("/basic-configurations")
}