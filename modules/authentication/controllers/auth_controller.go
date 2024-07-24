package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/config"
	configurations "github.com/ortizdavid/golang-modular-software/modules/configurations/services"
	services "github.com/ortizdavid/golang-modular-software/modules/authentication/services"
	"github.com/ortizdavid/go-nopain/encryption"
)

type AuthController struct {
	authService *services.AuthService
}


func (auth AuthController) Routes(router *fiber.App) {
	router.Get("/", auth.index)

	group := router.Group("/auth")
	group.Get("/login", auth.loginForm)
	group.Post("/login", auth.login)
	group.Get("/logout", auth.logout)
	group.Get("/recover-password/:token", auth.recoverPasswordForm)
	group.Post("/recover-password/:token", auth.recoverPassword)
	group.Get("/get-recover-link", auth.getRecoverLinkForm)
	group.Post("/get-recover-link", auth.getRecoverLink)
}


func (AuthController) index(c *fiber.Ctx) error {
	return c.Redirect("/auth/login")
}



func (AuthController) loginForm(c *fiber.Ctx) error {
	return c.Render("authentication/login", fiber.Map{
		"Title": "Authentication",
	})
}


func (AuthController) login(c *fiber.Ctx) error {
	var userModel models.UserModel
	userName := c.FormValue("username")
	password := c.FormValue("password")
	user, _ := userModel.FindByUserName(userName)
	exists, _ := userModel.ExistsActiveUser(userName)
	hashedPassword := user.Password
	
	if exists && encryption.CheckPassword(hashedPassword, password) {
		store := config.GetSessionStore()
		session, _ := store.Get(c)
		session.Set("username", userName)
		session.Set("password", hashedPassword)
		session.Set("authenticated", true)
		session.Save()
		//Update Token
		user.Token = encryption.GenerateRandomToken()
		_, err := userModel.Update(user)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		authLogger.Info(fmt.Sprintf("User '%s' authenticated sucessful!", userName), config.LogRequestPath(c))
		return c.Status(fiber.StatusOK).Redirect("/home")
	} else {
		authLogger.Error(fmt.Sprintf("User '%s' failed to login", userName), config.LogRequestPath(c))
		return c.Status(fiber.StatusUnauthorized).Redirect("/auth/login")
	}
}


func (AuthController) logout(c *fiber.Ctx) error {
	loggedUser := models.GetLoggedUser(c)
	userName := loggedUser.UserName
	store := config.GetSessionStore()
	session, _ := store.Get(c)
	session.Destroy()
	authLogger.Info(fmt.Sprintf("User '%s' logged out", userName), config.LogRequestPath(c))
	return c.Redirect("/auth/login")
}


func (AuthController) recoverPasswordForm(c *fiber.Ctx) error {
	token := c.Params("token")
	user, _ := models.UserModel{}.FindByToken(token)
	return c.Render("authentication/recover-password", fiber.Map{
		"Title": "Recuperação da Senha",
		"User": user,
	})
}


func (AuthController) recoverPassword(c *fiber.Ctx) error {
	password := c.FormValue("password")
	//passwordConf := c.FormValue("password_conf")
	token := c.Params("token")
	var userModel models.UserModel
	
	user, _ := userModel.FindByToken(token)
	user.Password = encryption.HashPassword(password)
	_, err := userModel.Update(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
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
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	authLogger.Info(fmt.Sprintf("User '%s' recovered password", user.UserName), config.LogRequestPath(c))
	return c.Redirect("/auth/login")
}


func (AuthController) getRecoverLinkForm(c *fiber.Ctx) error {
	return c.Render("authentication/get-recover-link", fiber.Map{
		"Title": "Get Recovery Link",
	})
}


func (AuthController) getRecoverLink(c *fiber.Ctx) error {
	email := c.FormValue("email")
	var userModel models.UserModel
	user, _ := userModel.FindByUserName(email)
	
	//enviar os credenciais por email
	emailService := configurations.DefaultEmailService()
	recoverLink := fmt.Sprintf("%s/auth/recover-password/%s", c.BaseURL(), user.Token)
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
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	authLogger.Info(fmt.Sprintf("User '%s' recovered password", email), config.LogRequestPath(c))
	return c.Redirect("/auth/get-recover-link")
}