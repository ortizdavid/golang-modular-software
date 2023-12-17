package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/config"
	configurations "github.com/ortizdavid/golang-modular-software/modules/configurations/models"
	models "github.com/ortizdavid/golang-modular-software/modules/authentication/models"
	"github.com/ortizdavid/go-nopain/encryption"
)

type AuthController struct {
}


func (auth AuthController) Routes(router *fiber.App) {
	group := router.Group("/authentication")
	group.Get("/login", auth.loginForm)
	group.Post("/login", auth.login)
	group.Get("/logout", auth.logout)
	group.Get("/recover-password/:token", auth.recoverPasswordForm)
	group.Post("/recover-password/:token", auth.recoverPassword)
	group.Get("/get-recover-link", auth.getRecoverLinkForm)
	group.Post("/get-recover-link", auth.getRecoverLink)
}


func (AuthController) loginForm(ctx *fiber.Ctx) error {
	return ctx.Render("authentication/login", fiber.Map{
		"Title": "Authentication",
	})
}


func (AuthController) login(ctx *fiber.Ctx) error {
	var userModel models.UserModel
	userName := ctx.FormValue("username")
	password := ctx.FormValue("password")
	user, _ := userModel.FindByUserName(userName)
	exists, _ := userModel.ExistsActiveUser(userName)
	hashedPassword := user.Password
	
	if exists && encryption.CheckPassword(hashedPassword, password) {
		store := config.GetSessionStore()
		session, _ := store.Get(ctx)
		session.Set("username", userName)
		session.Set("password", hashedPassword)
		session.Set("authenticated", true)
		session.Save()
		//Update Token
		user.Token = encryption.GenerateRandomToken()
		_, err := userModel.Update(user)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		loggerAuth.Info(fmt.Sprintf("User '%s' authenticated sucessfully!", userName), config.LogRequestPath(ctx))
		return ctx.Status(fiber.StatusOK).Redirect("/inicio")
	} else {
		loggerAuth.Error(fmt.Sprintf("User '%s' failed to login", userName), config.LogRequestPath(ctx))
		return ctx.Status(fiber.StatusUnauthorized).Redirect("/authentication/login")
	}
}


func (AuthController) logout(ctx *fiber.Ctx) error {
	loggedUser := models.GetLoggedUser(ctx)
	userName := loggedUser.UserName
	store := config.GetSessionStore()
	session, _ := store.Get(ctx)
	session.Destroy()
	loggerAuth.Info(fmt.Sprintf("User '%s' logged out", userName), config.LogRequestPath(ctx))
	return ctx.Redirect("/authentication/login")
}


func (AuthController) recoverPasswordForm(ctx *fiber.Ctx) error {
	token := ctx.Params("token")
	user, _ := models.UserModel{}.FindByToken(token)
	return ctx.Render("authentication/recover-password", fiber.Map{
		"Title": "Recuperação da Senha",
		"User": user,
	})
}


func (AuthController) recoverPassword(ctx *fiber.Ctx) error {
	password := ctx.FormValue("password")
	//passwordConf := ctx.FormValue("password_conf")
	token := ctx.Params("token")
	var userModel models.UserModel
	
	user, _ := userModel.FindByToken(token)
	user.Password = encryption.HashPassword(password)
	_, err := userModel.Update(user)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	//enviar os credenciais por email
	emailService := configurations.DefaultEmailService()
	htmlBody := `
		<html>
			<body>
				<h1>Password Changed!</h1>
				<p>Hello, `+user.UserName+`!</p>
				<p>Your new password: <b>`+password+`</b></p>
			</body>
		</html>`
	err = emailService.SendHTMLEmail(user.UserName, "New Password", htmlBody)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	loggerAuth.Info(fmt.Sprintf("User '%s' recovered password", user.UserName), config.LogRequestPath(ctx))
	return ctx.Redirect("/authentication/login")
}


func (AuthController) getRecoverLinkForm(ctx *fiber.Ctx) error {
	return ctx.Render("authentication/get-recover-link", fiber.Map{
		"Title": "Get Recovery Link",
	})
}


func (AuthController) getRecoverLink(ctx *fiber.Ctx) error {
	email := ctx.FormValue("email")
	var userModel models.UserModel
	user, _ := userModel.FindByUserName(email)
	
	//enviar os credenciais por email
	emailService := configurations.DefaultEmailService()
	recoverLink := fmt.Sprintf("%s/authentication/recover-password/%s", ctx.BaseURL(), user.Token)
	htmlBody := `
		<html>
			<body>
				<h1>Password Recovery!</h1>
				<p>Hello, `+user.UserName+`!</p>
				<p>To recover password Click <a href="`+recoverLink+`">Here</a></p>
			</body>
		</html>`
	err := emailService.SendHTMLEmail(user.UserName, "Password Recovery", htmlBody)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	loggerAuth.Info(fmt.Sprintf("User '%s' recovered password", email), config.LogRequestPath(ctx))
	return ctx.Redirect("/authentication/get-recover-link")
}