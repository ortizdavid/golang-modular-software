package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/go-nopain/encryption"
	"github.com/ortizdavid/golang-modular-software/common/config"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/entities"
	services "github.com/ortizdavid/golang-modular-software/modules/authentication/services"
	configurations "github.com/ortizdavid/golang-modular-software/modules/configurations/services"
	"gorm.io/gorm"
)

type AuthController struct {
	service *services.AuthService
	configService *configurations.BasicConfigurationService
	infoLogger *helpers.Logger
	errorLogger *helpers.Logger
}

func NewAuthController(db *gorm.DB) *AuthController {
	return &AuthController{
		service:       services.NewAuthService(db),
		configService: configurations.NewBasicConfigurationService(db),
		infoLogger:    helpers.NewInfoLogger("auth-info.log"),
		errorLogger:   helpers.NewInfoLogger("auth-error.log"),
	}
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


func (ctrl *AuthController) index(c *fiber.Ctx) error {
	return c.Redirect("/auth/login")
}


func (ctrl *AuthController) loginForm(c *fiber.Ctx) error {
	return c.Render("authentication/login", fiber.Map{
		"Title": "Authentication",
	})
}

func (ctrl *AuthController) login(c *fiber.Ctx) error {
	var request entities.LoginRequest
	if err := c.BodyParser(&request); err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	err := ctrl.service.Authenticate(c.Context(), c, request)
	if err != nil {
		ctrl.errorLogger.Error(c, fmt.Sprintf("User '%s' failed to login", request.UserName))
		return c.Status(fiber.StatusUnauthorized).Redirect("/auth/login")
	}
	ctrl.infoLogger.Info(c, fmt.Sprintf("User '%s' authenticated sucessful!", request.UserName))
	return c.Status(fiber.StatusOK).Redirect("/home")
}


func (ctrl *AuthController) logout(c *fiber.Ctx) error {
	loggedUser := models.GetLoggedUser(c)
	userName := loggedUser.UserName
	store := config.GetSessionStore()
	session, _ := store.Get(c)
	session.Destroy()
	authLogger.Info(fmt.Sprintf("User '%s' logged out", userName))
	return c.Redirect("/auth/login")
}


func (ctrl *AuthController) recoverPasswordForm(c *fiber.Ctx) error {
	token := c.Params("token")
	user, _ := models.UserModel{}.FindByToken(token)
	return c.Render("authentication/recover-password", fiber.Map{
		"Title": "Recuperação da Senha",
		"User": user,
	})
}


func (ctrl *AuthController) recoverPassword(c *fiber.Ctx) error {
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
	authLogger.Info(fmt.Sprintf("User '%s' recovered password", user.UserName))
	return c.Redirect("/auth/login")
}


func (ctrl *AuthController) getRecoverLinkForm(c *fiber.Ctx) error {
	return c.Render("authentication/get-recover-link", fiber.Map{
		"Title": "Get Recovery Link",
	})
}


func (ctrl *AuthController) getRecoverLink(c *fiber.Ctx) error {
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
	authLogger.Info(fmt.Sprintf("User '%s' recovered password", email))
	return c.Redirect("/auth/get-recover-link")
}