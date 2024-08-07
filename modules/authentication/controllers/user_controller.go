package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/apperrors"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
	"github.com/ortizdavid/golang-modular-software/common/middlewares"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/entities"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/services"
	configurations "github.com/ortizdavid/golang-modular-software/modules/configurations/services"
)

type UserController struct {
	service *services.UserService
	authService *services.AuthService
	roleService *services.RoleService
	appConfig *configurations.AppConfiguration
	infoLogger *helpers.Logger
	errorLogger *helpers.Logger
}

func NewUserController(db *database.Database) *UserController {
	return &UserController{
		service:       services.NewUserService(db),
		authService:   services.NewAuthService(db),
		roleService:   services.NewRoleService(db),
		appConfig: 		configurations.LoadAppConfigurations(db),
		infoLogger:    helpers.NewInfoLogger("users-info.log"),
		errorLogger:   helpers.NewInfoLogger("users-error.log"),
	}
}

func (ctrl *UserController) Routes(router *fiber.App, db *database.Database) {
	authMiddleware := middlewares.NewAuthenticationMiddleware(db)

	group := router.Group("/users", authMiddleware.CheckLoggedUser)
	group.Get("/", ctrl.index)
	group.Get("/create", ctrl.createForm)
	group.Post("/create", ctrl.create)
	group.Get("/:id/edit", ctrl.editForm)
	group.Post("/:id/edit", ctrl.edit)
	group.Get("/:id/details", ctrl.details)
	group.Get("/:id/assign-role", ctrl.assignRoleForm)
	group.Post("/:id/assign-role", ctrl.assignRole)
	group.Get("/:id/deactivate", ctrl.deactivateForm)
	group.Post("/:id/deactivate", ctrl.deactivate)
	group.Get("/:id/activate", ctrl.activateForm)
	group.Post("/:id/activate", ctrl.activate)
	group.Get("/:id/reset-password", ctrl.resetPasswordForm)
	group.Post("/:id/reset-password", ctrl.resetPassword)
	group.Get("/search", ctrl.searchForm)
	group.Get("/search-results", ctrl.search)
	group.Get("/active-users", ctrl.getAllActiveUsers)
	group.Get("/inactive-users", ctrl.getAllInactiveUsers)
	group.Get("/online-users", ctrl.getAllOnlineUsers)
	group.Get("/offline-users", ctrl.getAllOfflineUsers)
}

func (ctrl *UserController) index(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
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
		"AppConfig": ctrl.appConfig,
	})
}

func (ctrl *UserController) createForm(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	roles, err := ctrl.roleService.GetAllRoles(c.Context())
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	return c.Render("authentication/user/create", fiber.Map{
		"Title": "Add User",
		"Roles": roles,
		"LoggedUser": loggedUser,
		"AppConfig": ctrl.appConfig,
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

func (ctrl *UserController) editForm(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := ctrl.service.GetUserByUniqueId(c.Context(), id)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	return c.Render("authentication/user/edit", fiber.Map{
		"Title": "Edit User",
		"User": user,
		"LoggedUser": loggedUser,
		"AppConfig": ctrl.appConfig,
	})
}

func (ctrl *UserController) edit(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := ctrl.service.GetUserByUniqueId(c.Context(), id)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	var request entities.EditUserRequest
	if err := c.BodyParser(&request); err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	err = ctrl.service.EditUser(c.Context(), user.UserId, request)
	if err != nil {
		ctrl.errorLogger.Error(c, err.Error())
		return helpers.HandleHttpErrors(c, err)
	}
	ctrl.infoLogger.Info(c, "User '"+user.UserName+"' edited successfuly")
	return c.Redirect("/users")
}

func (ctrl *UserController) details(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	user, err := ctrl.service.GetUserByUniqueId(c.Context(), id)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	return c.Render("authentication/user/details", fiber.Map{
		"Title": "User Details",
		"User": user,
		"LoggedUser": loggedUser,
		"AppConfig": ctrl.appConfig,
	})
}

func (ctrl *UserController) assignRoleForm(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := ctrl.service.GetUserByUniqueId(c.Context(), id)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	roles, err := ctrl.roleService.GetAllRoles(c.Context())
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	return c.Render("authentication/user/assign-role", fiber.Map{
		"Title": "Assign Role",
		"Roles": roles,
		"User": user,
		"LoggedUser": loggedUser,
		"AppConfig": ctrl.appConfig,
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
	ctrl.infoLogger.Info(c, fmt.Sprintf("User '%s' assigned role %d", user.UserName, request.RoleId))
	return c.Redirect("/users")
}

func (ctrl *UserController) searchForm(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	return c.Render("authentication/user/search", fiber.Map{
		"Title": "Search Users",
		"LoggedUser": loggedUser,
		"AppConfig": ctrl.appConfig,
	})
}

func (ctrl *UserController) search(c *fiber.Ctx) error {
	searcParam := c.FormValue("search_param")
	request := entities.SearchUserRequest{SearchParam: searcParam}
    loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
    params := helpers.GetPaginationParams(c)
    pagination, err := ctrl.service.SearchUsers(c.Context(), c, request, params)
    if err != nil {
        return helpers.HandleHttpErrors(c, err)
    }
    ctrl.infoLogger.Info(c, fmt.Sprintf("User '%s' searched for '%v' and found %d results", loggedUser.UserName, request.SearchParam, pagination.MetaData.TotalItems))
    return c.Render("authentication/user/search-results", fiber.Map{
        "Title":        "Search Results",
        "Pagination":   pagination,
        "Param":        request.SearchParam,
        "CurrentPage":  pagination.MetaData.CurrentPage + 1,
        "TotalPages":   pagination.MetaData.TotalPages + 1,
        "LoggedUser":   loggedUser,
        "AppConfig":  ctrl.appConfig,
    })
}

func (ctrl *UserController) deactivateForm(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := ctrl.service.GetUserByUniqueId(c.Context(), id)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	if loggedUser.UserId == user.UserId {
		return helpers.HandleHttpErrors(c, apperrors.NewConflictError("You cannot deactivate your own account"))
	}
	return c.Render("authentication/user/deactivate", fiber.Map{
		"Title": "Deactivate User",
		"User": user,
		"LoggedUser": loggedUser,
		"AppConfig": ctrl.appConfig,
	})
}

func (ctrl *UserController) deactivate(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := ctrl.service.GetUserByUniqueId(c.Context(), id)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	if loggedUser.UserId == user.UserId {
		return helpers.HandleHttpErrors(c, apperrors.NewConflictError("You cannot deactivate your own account"))
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
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	if loggedUser.UserId == user.UserId {
		return helpers.HandleHttpErrors(c, apperrors.NewConflictError("You cannot activate your own account"))
	}
	return c.Render("authentication/user/activate", fiber.Map{
		"Title": "Activate User",
		"User": user,
		"LoggedUser": loggedUser,
		"AppConfig": ctrl.appConfig,
	})
}

func (ctrl *UserController) activate(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := ctrl.service.GetUserByUniqueId(c.Context(), id)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	if loggedUser.UserId == user.UserId {
		return helpers.HandleHttpErrors(c, apperrors.NewConflictError("You cannot activate your own account"))
	}
	err = ctrl.service.ActivateUser(c.Context(), user.UserId)
	if err != nil {
		ctrl.errorLogger.Error(c, err.Error())
		return helpers.HandleHttpErrors(c, err)
	}
	ctrl.infoLogger.Info(c, fmt.Sprintf("User '%s' activated successfully!", user.UserName))
	return c.Redirect("/users/"+id+"/details")
}

func (ctrl *UserController) getAllActiveUsers(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	params := helpers.GetPaginationParams(c)
	pagination, err := ctrl.service.GetAllActiveUsers(c.Context(), c, params)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	return c.Render("authentication/user/active-users", fiber.Map{
		"Title": "Active Users",
		"Pagination": pagination,
		"LoggedUser": loggedUser,
		"AppConfig": ctrl.appConfig,
	})
}

func (ctrl *UserController) getAllInactiveUsers(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	params := helpers.GetPaginationParams(c)
	pagination, err := ctrl.service.GetAllInactiveUsers(c.Context(), c, params)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	return c.Render("authentication/user/inactive-users", fiber.Map{
		"Title": "Inactive Users",
		"Pagination": pagination,
		"LoggedUser": loggedUser,
		"AppConfig": ctrl.appConfig,
	})
}

func (ctrl *UserController) getAllOnlineUsers(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	params := helpers.GetPaginationParams(c)
	pagination, err := ctrl.service.GetAllOnlineUsers(c.Context(), c, params)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	return c.Render("authentication/user/online-users", fiber.Map{
		"Title": "Online Users",
		"Pagination": pagination,
		"LoggedUser": loggedUser,
		"AppConfig": ctrl.appConfig,
	})
}

func (ctrl *UserController) getAllOfflineUsers(c *fiber.Ctx) error {
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	params := helpers.GetPaginationParams(c)
	pagination, err := ctrl.service.GetAllOfflineUsers(c.Context(), c, params)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	return c.Render("authentication/user/offline-users", fiber.Map{
		"Title": "Offline Users",
		"Pagination": pagination,
		"LoggedUser": loggedUser,
		"AppConfig": ctrl.appConfig,
	})
}

func (ctrl *UserController) resetPasswordForm(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := ctrl.service.GetUserByUniqueId(c.Context(), id)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	if loggedUser.UserId == user.UserId {
		return helpers.HandleHttpErrors(c, apperrors.NewConflictError("You cannot reset your own password"))
	}
	return c.Render("authentication/user/reset-password", fiber.Map{
		"Title":  "Reset Password",
		"LoggedUser": loggedUser,
		"AppConfig":  ctrl.appConfig,
		"User": user,
	})
}

func (ctrl *UserController) resetPassword(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := ctrl.service.GetUserByUniqueId(c.Context(), id)
	if err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	var request entities.ResetPasswordRequest
	if err := c.BodyParser(&request); err != nil {
		return helpers.HandleHttpErrors(c, err)
	}
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	if loggedUser.UserId == user.UserId {
		return helpers.HandleHttpErrors(c, apperrors.NewConflictError("You cannot reset your own password"))
	}
	fmt.Printf("\nRESET PASSWORD\nNew Password: %s\nPassword Conf: %s", request.NewPassword, request.PasswordConf)
	err = ctrl.service.ResetUserPassword(c.Context(), user.UserId, request)
	if err != nil {
		ctrl.errorLogger.Error(c, err.Error())
		return helpers.HandleHttpErrors(c, err)
	}
	ctrl.infoLogger.Info(c, fmt.Sprintf("User '%s' password reseted", loggedUser.UserName))
	return c.Redirect("/users/"+user.UniqueId+"/details")
}