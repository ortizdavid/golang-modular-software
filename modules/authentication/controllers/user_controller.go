package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
	"github.com/ortizdavid/golang-modular-software/database"
	entities "github.com/ortizdavid/golang-modular-software/modules/authentication/entities"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/services"
	configurations "github.com/ortizdavid/golang-modular-software/modules/configurations/services"
)

type UserController struct {
	service *services.UserService
	roleService *services.RoleService
	authService *services.AuthService
	configService *configurations.BasicConfigurationService
	infoLogger *helpers.Logger
	errorLogger *helpers.Logger
}

func NewUserController(db *database.Database) *UserController {
	return &UserController{
		service:       services.NewUserService(db),
		roleService:   services.NewRoleService(db),
		authService:   services.NewAuthService(db),
		configService: configurations.NewBasicConfigurationService(db),
		infoLogger:    helpers.NewInfoLogger("users-info.log"),
		errorLogger:   helpers.NewInfoLogger("users-error.log"),
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
	router.Get("/active-users", ctrl.getAllActiveUsers)
	router.Post("/inactive-users", ctrl.getAllInactiveUsers)
}

func (ctrl *UserController) index(c *fiber.Ctx) error {
	loggedUser, err := ctrl.authService.GetLoggedUser(c.Context(), c)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	basicConfig, err := ctrl.configService.GetBasicConfiguration(c.Context())
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	params := helpers.GetPaginationParams(c)
	pagination, err := ctrl.service.GetAllUsers(c.Context(), c, params)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	return c.Render("authentication/user/index", fiber.Map{
		"Title": "Users",
		"Pagination": pagination,
		"CurrentPage": pagination.MetaData.CurrentPage + 1,
		"TotalPages": pagination.MetaData.TotalPages + 1,
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
	return c.Render("authentication/user/details", fiber.Map{
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
	return c.Render("authentication/user/create", fiber.Map{
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
		ctrl.errorLogger.Error(c, err.Error())
		return helpers.HandleHttpErrors(c, err)
	}
	ctrl.infoLogger.Info(c, "User '"+request.UserName+"' added successfully")
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
	return c.Render("authentication/user/edit", fiber.Map{
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
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	var request entities.AssignUserRoleRequest
	if err := c.BodyParser(&request); err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	err = ctrl.service.AssignUserRole(c.Context(), user.UserId, request)
	if err != nil {
		ctrl.errorLogger.Error(c, err.Error())
		return helpers.HandleHttpErrors(c, err)
	}
	ctrl.infoLogger.Info(c, "User '"+user.UserName+"' assigned role")
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
	searchParam := c.FormValue("search_param")
	loggedUser, err := ctrl.authService.GetLoggedUser(c.Context(), c)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	basicConfig, err := ctrl.configService.GetBasicConfiguration(c.Context())
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	count, _ := ctrl.service.CountUsersByParam(c.Context(), searchParam)
	params := helpers.GetPaginationParams(c)
	pagination, _ := ctrl.service.SearchUsers(c.Context(), c, searchParam, params)
	ctrl.infoLogger.Info(c, fmt.Sprintf("User '%s' searched for '%v' and found %d results", loggedUser.UserName, searchParam, count))
	return c.Render("authentication/user/search-results", fiber.Map{
		"Title": "Results",
		"Pagination": pagination,
		"Param": searchParam,
		"Count": count,
		"CurrentPage": pagination.MetaData.CurrentPage + 1,
		"TotalPages": pagination.MetaData.TotalPages + 1,
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
		ctrl.errorLogger.Error(c, err.Error())
		return helpers.HandleHttpErrors(c, err)
	}
	ctrl.infoLogger.Info(c, fmt.Sprintf("User '%s' deactivated successfully!", user.UserName))
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
		ctrl.errorLogger.Error(c, err.Error())
		return helpers.HandleHttpErrors(c, err)
	}
	ctrl.infoLogger.Info(c, fmt.Sprintf("User '%s' activated successfully!", user.UserName))
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
		ctrl.errorLogger.Error(c, err.Error())
		return helpers.HandleHttpErrors(c, err)
	}
	ctrl.infoLogger.Info(c, fmt.Sprintf("User '%s' added image", user.UserName))
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
		ctrl.errorLogger.Error(c, err.Error())
		return helpers.HandleHttpErrors(c, err)
	}
	ctrl.infoLogger.Info(c, fmt.Sprintf("User '%s' updated password", loggedUser.UserName))
	return c.Redirect("/auth/login")
}


func (ctrl *UserController) getAllActiveUsers(c *fiber.Ctx) error {
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
	users, err := ctrl.service.GetAllActiveUsers(c.Context(), c, params)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	return c.Render("authentication/user/active-users", fiber.Map{
		"Title": "Active Users",
		"Users": users,
		"Count": count,
		"LoggedUser": loggedUser,
		"BasicConfig": basicConfig,
	})
}

func (ctrl *UserController) getAllInactiveUsers(c *fiber.Ctx) error {
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
	users, err := ctrl.service.GetAllInactiveUsers(c.Context(), c, params)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	return c.Render("authentication/user/inactive-users", fiber.Map{
		"Title": "Inactive Users",
		"Users": users,
		"Count": count,
		"LoggedUser": loggedUser,
		"BasicConfig": basicConfig,
	})
}