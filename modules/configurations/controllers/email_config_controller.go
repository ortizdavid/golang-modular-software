package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/config"
	authentication "github.com/ortizdavid/golang-modular-software/modules/authentication/models"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/models"
	"github.com/ortizdavid/go-nopain/conversion"
)


type EmailConfigController struct {
}


func (configuration EmailConfigController) Routes(router *fiber.App) {
	group := router.Group("/email-configurations")
	group.Get("", configuration.index)
	group.Get("/edit", configuration.editForm)
	group.Post("/edit", configuration.edit)
}


func (EmailConfigController) index(ctx *fiber.Ctx) error {
	configuration, _ := models.GetBasicConfiguration()
	return ctx.Render("configurations/email/index", fiber.Map{
		"Title": "Email Configurations",
		"BasicConfiguration": configuration,
		"LoggedUser": authentication.GetLoggedUser(ctx),
	})
}


func (EmailConfigController) editForm(ctx *fiber.Ctx) error {
	configurationEmail, _ := models.GetEmailConfiguration()
	configurationBasica, _ := models.GetBasicConfiguration()
	return ctx.Render("configurations/email/edit", fiber.Map{
		"Title": "Edita EmailConfig de Email",
		"EmailConfiguration": configurationEmail,
		"BasicConfiguration": configurationBasica,
		"LoggedUser": authentication.GetLoggedUser(ctx),
	})
}


func (EmailConfigController) edit(ctx *fiber.Ctx) error {
	smtpServer := ctx.FormValue("SMTPServer")
	smtpPort := ctx.FormValue("SMTPPort")
	senderEmail := ctx.FormValue("SenderEmail")
	senderPassword := ctx.FormValue("SenderPassword")
	loggedUser := authentication.GetLoggedUser(ctx)

	var configurationModel models.EmailConfigurationModel
	configuration, _ := models.GetEmailConfiguration()
	configuration.SMTPServer = smtpServer
	configuration.SMTPPort = conversion.StringToInt(smtpPort)
	configuration.SenderEmail = senderEmail
	configuration.SenderPassword = senderPassword
	_, err := configurationModel.Update(configuration)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	loggerConfiguration.Info(fmt.Sprintf("User '%s' updated email configurations!", loggedUser.UserName), config.LogRequestPath(ctx))
	return ctx.Redirect("/email-configurations")
}
