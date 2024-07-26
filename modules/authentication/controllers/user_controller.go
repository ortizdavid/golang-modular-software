package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/config"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
	entities "github.com/ortizdavid/golang-modular-software/modules/authentication/entities"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/services"
	authentication "github.com/ortizdavid/golang-modular-software/modules/authentication/services"
	configurations "github.com/ortizdavid/golang-modular-software/modules/configurations/services"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserController struct {
	service *services.UserService
	roleService *services.RoleService
	authService *services.AuthService
	configService *configurations.BasicConfigurationService
	infoLogger *zap.Logger
	errorLogger *zap.Logger
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{
		service:       authentication.NewUserService(db),
		roleService:   authentication.NewRoleService(db),
		authService:   authentication.NewAuthService(db),
		configService: configurations.NewBasicConfigurationService(db),
		infoLogger:    userInfoLogger,
		errorLogger:   userErrorLogger,
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
	if err := c.QueryParser(&params); err != nil {
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
		ctrl.errorLogger.Error(err.Error())
		return helpers.HandleHttpErrors(c, err)
	}
	ctrl.infoLogger.Info("User '"+request.UserName+"' added successfully", config.LogRequestPath(c))
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
		ctrl.errorLogger.Error(err.Error())
		return helpers.HandleHttpErrors(c, err)
	}
	ctrl.infoLogger.Info("User '"+user.UserName+"' assigned role", config.LogRequestPath(c))
	return c.Redirect("/users")
}


func (ctrl *UserController) searchForm(c *fiber.Ctx) error {
	loggedUser, err := ctrl.authService.GetLoggedUser(c.Context(), c)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	basicConfig, err := ctrl.configService.GetBasicConfiguration(c.Context())
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	return c.Render("authentication/user/search", fiber.Map{
		"Title": "Search Users",
		"LoggedUser": loggedUser,
		"BasicConfig": basicConfig,
	})
}


func (ctrl *UserController) search(c *fiber.Ctx) error {
	var paginationParams helpers.PaginationParam
	searchParam := c.FormValue("search_param")
	loggedUser, err := ctrl.authService.GetLoggedUser(c.Context(), c)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	basicConfig, err := ctrl.configService.GetBasicConfiguration(c.Context())
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	count, err := ctrl.service.CountUsersByParam(c.Context(), searchParam)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	if err := c.QueryParser(&paginationParams); err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	results, err := ctrl.service.SearchUsers(c.Context(), c, searchParam, paginationParams)
	if err != nil {
		ctrl.errorLogger.Error(err.Error())
		return helpers.HandleHttpErrors(c, err)
	}
	ctrl.infoLogger.Info(fmt.Sprintf("User '%s' searched for '%v' and found %d results", loggedUser.UserName, searchParam, count), config.LogRequestPath(c))
	return c.Render("authentication/user/search-results", fiber.Map{
		"Title": "Results",
		"Results": results,
		"Param": searchParam,
		"Count": count,
		"LoggedUser": loggedUser,
		"BasicConfig": basicConfig,
	})
}


func (ctrl *UserController) deactivateForm(c *fiber.Ctx) error {
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
	return c.Render("authentication/user/deactivate", fiber.Map{
		"Title": "Deactivate User",
		"User": user,
		"LoggedUser": loggedUser,
		"BasicConfig": basicConfig,
	})
}


func (ctrl *UserController) deactivate(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := ctrl.service.GetUserByUniqueId(c.Context(), id)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	err = ctrl.service.DeactivateUser(c.Context(), user.UserId)
	if err != nil {
		ctrl.errorLogger.Error(err.Error())
		return helpers.HandleHttpErrors(c, err)
	}
	ctrl.infoLogger.Info(fmt.Sprintf("User '%s' deactivated successfully!", user.UserName), config.LogRequestPath(c))
	return c.Redirect("/users/"+id+"/details")
}


func (ctrl *UserController) activateForm(c *fiber.Ctx) error {
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
	return c.Render("authentication/activate", fiber.Map{
		"Title": "Activate User",
		"User": user,
		"LoggedUser": loggedUser,
		"BasicConfig": basicConfig,
	})
}


func (ctrl *UserController) activate(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := ctrl.service.GetUserByUniqueId(c.Context(), id)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	err = ctrl.service.ActivateUser(c.Context(), user.UserId)
	if err != nil {
		ctrl.errorLogger.Error(err.Error())
		return helpers.HandleHttpErrors(c, err)
	}
	ctrl.infoLogger.Info(fmt.Sprintf("User '%s' activated successfully!", user.UserName), config.LogRequestPath(c))
	return c.Redirect("/users/"+id+"/details")
}


func (ctrl *UserController) changeUserImageForm(c *fiber.Ctx) error {
	loggedUser, err := ctrl.authService.GetLoggedUser(c.Context(), c)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	basicConfig, err := ctrl.configService.GetBasicConfiguration(c.Context())
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	return c.Render("authentication/user/change-image", fiber.Map{
		"Title": "Add Image",
		"LoggedUser": loggedUser,
		"BasicConfig": basicConfig,
	})
}


func (ctrl *UserController) changeUserImage(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := ctrl.service.GetUserByUniqueId(c.Context(), id)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	err = ctrl.service.ChangeUserImage(c.Context(), c, user.UserId)
	if err != nil {
		ctrl.errorLogger.Error(err.Error())
		return helpers.HandleHttpErrors(c, err)
	}
	ctrl.infoLogger.Info(fmt.Sprintf("User '%s' added image", user.UserName), config.LogRequestPath(c))
	return c.Redirect("/user-data")
}


func (ctrl *UserController) changePasswordForm(c *fiber.Ctx) error {
	loggedUser, err := ctrl.authService.GetLoggedUser(c.Context(), c)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	basicConfig, err := ctrl.configService.GetBasicConfiguration(c.Context())
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	return c.Render("authentication/user/change-password", fiber.Map{
		"Title": "Updated Password",
		"LoggedUser": loggedUser,
		"BasicConfig": basicConfig,
	})
}


func (ctrl *UserController) changePassword(c *fiber.Ctx) error {
	var request entities.ChangePasswordRequest
	loggedUser, err := ctrl.authService.GetLoggedUser(c.Context(), c)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	err = ctrl.service.ChangeUserPassword(c.Context(), loggedUser.UserId, request)
	if err != nil {
		ctrl.errorLogger.Error(err.Error())
		return helpers.HandleHttpErrors(c, err)
	}
	ctrl.infoLogger.Info(fmt.Sprintf("User '%s' updated password", loggedUser.UserName), config.LogRequestPath(c))
	return c.Redirect("/auth/login")
}
