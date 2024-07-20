package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/config"
	authentication "github.com/ortizdavid/golang-modular-software/modules/authentication/models"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/models"
	"github.com/ortizdavid/go-nopain/conversion"
)


type CompanyConfigController struct {
}


func (configuration CompanyConfigController) Routes(router *fiber.App) {
	group := router.Group("/company-configurations")
	group.Get("", configuration.index)
	group.Get("/edit", configuration.editForm)
	group.Post("/edit", configuration.edit)
}


func (CompanyConfigController) index(ctx *fiber.Ctx) error {
	configuration, _ := models.GetBasicConfiguration()
	return ctx.Render("configurations/company/index", fiber.Map{
		"Title": "Company Configurations",
		"BasicConfiguration": configuration,
		"LoggedUser": authentication.GetLoggedUser(ctx),
	})
}


func (CompanyConfigController) editForm(ctx *fiber.Ctx) error {
	configurationEmail, _ := models.GetCompanyConfiguration()
	configurationBasica, _ := models.GetBasicConfiguration()
	return ctx.Render("configuration/company/edit", fiber.Map{
		"Title": "Edit Company Configuarions",
		"CompanyConfiguration": configurationEmail,
		"BasicConfiguration": configurationBasica,
		"LoggedUser": authentication.GetLoggedUser(ctx),
	})
}


func (CompanyConfigController) edit(ctx *fiber.Ctx) error {
	companyName := ctx.FormValue("company_name")
	companyAcronym := ctx.FormValue("company_name")
	companyPhone := ctx.FormValue("company_phone")
	companyEmail := ctx.FormValue("company_email")
	companyMainColor := ctx.FormValue("company_main_color")
	loggedUser := authentication.GetLoggedUser(ctx)

	var configurationModel models.CompanyConfigurationModel
	configuration, _ := models.GetCompanyConfiguration()
	configuration.CompanyName= companyName
	configuration.CompanyAcronym = companyAcronym
	configuration.CompanyPhone = conversion.StringToInt(companyPhone)
	configuration.CompanyEmail = companyEmail
	configuration.CompanyMainColor = companyMainColor

	_, err := configurationModel.Update(configuration)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	loggerConfiguration.Info(fmt.Sprintf("User '%s' updated company configurations!", loggedUser.UserName), config.LogRequestPath(ctx))
	return ctx.Redirect("/email-configurations")
}
