package controllers

import (
	"fmt"
	"time"
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/go-nopain/encryption"
	"github.com/ortizdavid/golang-modular-software/common/config"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
	entities "github.com/ortizdavid/golang-modular-software/modules/authentication/entities"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/services"
	authentication "github.com/ortizdavid/golang-modular-software/modules/authentication/services"
	configurations "github.com/ortizdavid/golang-modular-software/modules/configurations/services"
	"gorm.io/gorm"
)

type UserController struct {
	service *services.UserService
	roleService *services.RoleService
	authService *services.AuthService
	configService *configurations.BasicConfigurationService
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{
		service: authentication.NewUserService(db),
	}
}

func (ctrl *UserController) Routes(router *fiber.App) {
	group := router.Group("/users")
	group.Get("/", ctrl.index)
	group.Get("/create", ctrl.createForm)
	group.Post("/create", ctrl.create)
	group.Get("/:id/details", ctrl.details)
	group.Get("/:id/assign-role", ctrl.assignRoleForm)
	group.Post("/:id/assign-role", ctrl.assignRole)
	group.Get("/:id/deactivate", ctrl.deactivateForm)
	group.Post("/:id/deactivate", ctrl.deactivate)
	group.Get("/:id/activate", ctrl.activateForm)
	group.Post("/:id/activate", ctrl.activate)
	group.Get("/change-image", ctrl.changeUserImageForm)
	group.Post("/change-image", ctrl.changeUserImage)
	group.Get("/search", ctrl.searchForm)
	group.Get("/search-results", ctrl.search)
	router.Get("/change-password", ctrl.changePasswordForm)
	router.Post("/change-password", ctrl.changePassword)
}

func (ctrl *UserController) index(c *fiber.Ctx) error {
	var params helpers.PaginationParam
	if err := c.QueryParser(&params); err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	loggedUser, err := ctrl.authService.GetLoggedUser(c.Context(), c)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	basicConfig, err := ctrl.configService.GetBasicConfiguration(c.Context())
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	count, err := ctrl.service.CountUsers(c.Context())
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	users, err := ctrl.service.GetAllUsers(c.Context(), c, params)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	return c.Render("authentication/user/index", fiber.Map{
		"Title": "Users",
		"Users": users,
		"Count": count,
		"LoggedUser": loggedUser,
		"BasicConfig": basicConfig,
	})
}

func (ctrl *UserController) details(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, err := ctrl.authService.GetLoggedUser(c.Context(), c)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	basicConfig, err := ctrl.configService.GetBasicConfiguration(c.Context())
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	user, err := ctrl.service.GetUserByUniqueId(c.Context(), id)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	return c.Render("user/details", fiber.Map{
		"Title": "User Details",
		"User": user,
		"LoggedUser": loggedUser,
		"BasicConfig": basicConfig,
	})
}

func (ctrl *UserController) createForm(c *fiber.Ctx) error {
	loggedUser, err := ctrl.authService.GetLoggedUser(c.Context(), c)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	basicConfig, err := ctrl.configService.GetBasicConfiguration(c.Context())
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	roles, err := ctrl.roleService.GetAllRoles(c.Context())
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	return c.Render("users/user/add", fiber.Map{
		"Title": "Add User",
		"Roles": roles,
		"LoggedUser": loggedUser,
		"BasicConfig": basicConfig,
	})
}

func (ctrl *UserController) create(c *fiber.Ctx) error {
	var request entities.CreateUserRequest
	if err := c.BodyParser(&request); err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	err := ctrl.service.CreateUser(c.Context(), request)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	userLogger.Info("User '"+request.UserName+"' added successfully", config.LogRequestPath(c))
	return c.Redirect("/users")
}


func (ctrl *UserController) assignRoleForm(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := ctrl.service.GetUserByUniqueId(c.Context(), id)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	loggedUser, err := ctrl.authService.GetLoggedUser(c.Context(), c)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	basicConfig, err := ctrl.configService.GetBasicConfiguration(c.Context())
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	roles, err := ctrl.roleService.GetAllRoles(c.Context())
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	return c.Render("users/user/edit", fiber.Map{
		"Title": "Assign Role",
		"Roles": roles,
		"User": user,
		"LoggedUser": loggedUser,
		"BasicConfig": basicConfig,
	})
}


func (ctrl *UserController) assignRole(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := ctrl.service.GetUserByUniqueId(c.Context(), id)
	var request entities.AssignUserRoleRequest
	if err := c.BodyParser(&request); err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	err = ctrl.service.AssignUserRole(c.Context(), user.UserId, request)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	userLogger.Info("User '"+user.UserName+"' assigned role", config.LogRequestPath(c))
	return c.Redirect("/users")
}


func (ctrl *UserController) searchForm(c *fiber.Ctx) error {
	basicConfig, _ := configurations.GetBasicConfiguration()
	return c.Render("user/user/search", fiber.Map{
		"Title": "Search Users",
		"LoggedUser": authentication.GetLoggedUser(c),
		"BasicConfig": basicConfig,
	})
}


func (ctrl *UserController) search(c *fiber.Ctx) error {
	param := c.FormValue("search_param")
	results, _ := models.UserModel{}.Search(param)
	count := len(results)
	loggedUser := authentication.GetLoggedUser(c)
	basicConfig, _ := configurations.GetBasicConfiguration()
	userLogger.Info(fmt.Sprintf("User '%s' searched for '%v' and found %d results", loggedUser.UserName, param, count), config.LogRequestPath(c))
	return c.Render("user/user/search-results", fiber.Map{
		"Title": "Results",
		"Results": results,
		"Param": param,
		"Count": count,
		"LoggedUser": loggedUser,
		"BasicConfig": basicConfig,
	})
}


func (ctrl *UserController) deactivateForm(c *fiber.Ctx) error {
	id := c.Params("id")
	user, _ := models.UserModel{}.GetDataByUniqueId(id)
	basicConfig, _ := configurations.GetBasicConfiguration()
	return c.Render("user/deactivate", fiber.Map{
		"Title": "Desactivar User",
		"User": user,
		"LoggedUser": authentication.GetLoggedUser(c),
		"BasicConfig": basicConfig,
	})
}


func (ctrl *UserController) deactivate(c *fiber.Ctx) error {
	id := c.Params("id")
	var userModel models.UserModel
	user, _ := userModel.FindByUniqueId(id)
	ctrl.Active = "No"
	ctrl.UpdatedAt = time.Now()
	userModel.Update(user)
	userLogger.Info(fmt.Sprintf("User '%s' deactivated successfully!", ctrl.UserName), config.LogRequestPath(c))
	return c.Redirect("/users/"+id+"/details")
}


func (ctrl *UserController) activateForm(c *fiber.Ctx) error {
	id := c.Params("id")
	user, _ := models.UserModel{}.GetDataByUniqueId(id)
	basicConfig, _ := configurations.GetBasicConfiguration()
	return c.Render("user/user/activate", fiber.Map{
		"Title": "Activar User",
		"User": user,
		"LoggedUser": authentication.GetLoggedUser(c),
		"BasicConfig": basicConfig,
	})
}


func (ctrl *UserController) activate(c *fiber.Ctx) error {
	id := c.Params("id")
	var userModel models.UserModel
	user, _ := userModel.FindByUniqueId(id)
	ctrl.Active = "Yes"
	ctrl.UpdatedAt = time.Now()
	ctrl.Token = encryption.GenerateRandomToken()
	_, err := userModel.Update(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	userLogger.Info(fmt.Sprintf("User '%s' activated successfully!", ctrl.UserName), config.LogRequestPath(c))
	return c.Redirect("/users/"+id+"/details")
}


func (ctrl *UserController) changeUserImageForm(c *fiber.Ctx) error {
	basicConfig, _ := configurations.GetBasicConfiguration()
	return c.Render("user/user/add-image", fiber.Map{
		"Title": "Add Image",
		"LoggedUser": authentication.GetLoggedUser(c),
		"BasicConfig": basicConfig,
	})
}


func (ctrl *UserController) changeUserImage(c *fiber.Ctx) error {
	userImage, _ := helpers.UploadFile(c, "user_image", "image", config.UploadImagePath())
	loggedUser := authentication.GetLoggedUser(c)
	var userModel models.UserModel
	user, _ := userModel.FindById(loggedUser.UserId)
	ctrl.Image = userImage
	ctrl.UpdatedAt = time.Now()
	_, err := userModel.Update(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	userLogger.Info(fmt.Sprintf("User '%s' added image", ctrl.UserName), config.LogRequestPath(c))
	return c.Redirect("/user-data")
}


func (ctrl *UserController) changePasswordForm(c *fiber.Ctx) error {
	basicConfig, _ := configurations.GetBasicConfiguration()
	return c.Render("user/user/change-password", fiber.Map{
		"Title": "Updated Password",
		"LoggedUser": authentication.GetLoggedUser(c),
		"BasicConfig": basicConfig,
	})
}


func (ctrl *UserController) changePassword(c *fiber.Ctx) error {
	password := c.FormValue("password")
	//passwordConf := c.FormValue("password_conf")
	loggedUser := authentication.GetLoggedUser(c)

	var userModel models.UserModel
	user, _ := userModel.FindById(loggedUser.UserId)
	ctrl.Password = encryption.HashPassword(password)
	_, err := userModel.Update(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	userLogger.Info(fmt.Sprintf("User '%s' updated password", ctrl.UserName), config.LogRequestPath(c))
	return c.Redirect("/auth/login")
}
