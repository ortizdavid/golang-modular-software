package controllers

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/go-nopain/conversion"
	"github.com/ortizdavid/go-nopain/encryption"
	"github.com/ortizdavid/golang-modular-software/config"
	"github.com/ortizdavid/golang-modular-software/helpers"
	entities "github.com/ortizdavid/golang-modular-software/modules/users/entities"
	models "github.com/ortizdavid/golang-modular-software/modules/users/models"
	configurations "github.com/ortizdavid/golang-modular-software/modules/configurations/models"
	authentication "github.com/ortizdavid/golang-modular-software/modules/authentication/models"
)

type UserController struct {
}


func (user UserController) Routes(router *fiber.App) {
	group := router.Group("/users")
	group.Get("/", user.index)
	group.Get("/add", user.addForm)
	group.Post("/add", user.add)
	group.Get("/:id/details", user.details)
	group.Get("/:id/edit", user.editForm)
	group.Post("/:id/edit", user.edit)
	group.Get("/:id/deactivate", user.deactivateForm)
	group.Post("/:id/deactivate", user.deactivate)
	group.Get("/:id/activate", user.activateForm)
	group.Post("/:id/activate", user.activate)
	group.Get("/add-image", user.addImageForm)
	group.Post("/add-image", user.addImage)
	group.Get("/search", user.searchForm)
	group.Get("/search-results", user.search)
	router.Get("/change-password", user.changePasswordForm)
	router.Post("/change-password", user.changePassword)
}


func (UserController) index(ctx *fiber.Ctx) error {
	var pagination helpers.Pagination
	var userModel models.UserModel
	
	loggedUser := authentication.GetLoggedUser(ctx)
	basicConfig, _ := configurations.GetBasicConfiguration()
	itemsPerPage := basicConfig.NumOfRecordsPerPage
	pageNumber := pagination.GetPageNumber(ctx, "page")
	startIndex := pagination.CalculateStartIndex(pageNumber, itemsPerPage)
	users, _ := userModel.FindAllDataLimit(startIndex, itemsPerPage)
	countUsers, _ := userModel.Count()
	count := int(countUsers)
	totalPages := pagination.CalculateTotalPages(count, itemsPerPage)

	if pageNumber>totalPages && count!=0 {
		return ctx.Status(fiber.StatusInternalServerError).Render("errors/pagination", fiber.Map{
			"Title": "Users",
			"TotalPages": totalPages, 
			"LoggedUser": loggedUser,
			"BasicConfig": basicConfig,
		})
	}
	return ctx.Render("user/index", fiber.Map{
		"Title": "Users",
		"Users": users,
		"Pagination": helpers.NewPaginationRender(pageNumber),
		"Count": count,
		"PageNumber": pageNumber,
		"TotalPages": totalPages,
		"LoggedUser": loggedUser,
		"BasicConfig": basicConfig,
	})
}


func (UserController) details(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	user, _ := models.UserModel{}.GetDataByUniqueId(id)
	basicConfig, _ := configurations.GetBasicConfiguration()
	return ctx.Render("user/details", fiber.Map{
		"Title": "User Details",
		"User": user,
		"LoggedUser": authentication.GetLoggedUser(ctx),
		"BasicConfig": basicConfig,
	})
}


func (UserController) addForm(ctx *fiber.Ctx) error {
	roles, _ := models.RoleModel{}.FindAll()
	basicConfig, _ := configurations.GetBasicConfiguration()
	return ctx.Render("users/user/add", fiber.Map{
		"Title": "Add User",
		"Roles": roles,
		"LoggedUser": authentication.GetLoggedUser(ctx),
		"BasicConfig": basicConfig,
	})
}


func (UserController) add(ctx *fiber.Ctx) error {
	userName := ctx.FormValue("username")
	password := ctx.FormValue("password")
	roleId := ctx.FormValue("role_id")

	var userModel models.UserModel
	user := entities.User{
		UserId:    0,
		RoleId:    conversion.StringToInt(roleId),
		UserName:  userName,
		Password:  encryption.HashPassword(password),
		Active:    "Yes",
		Image:     "",
		UniqueId:  encryption.GenerateUUID(),
		Token:     encryption.GenerateRandomToken(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	_, err := userModel.Create(user)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	loggerUser.Info("User '"+userName+"' added successfully", config.LogRequestPath(ctx))
	return ctx.Redirect("/users")
}


func (UserController) editForm(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	user, _ := models.UserModel{}.GetDataByUniqueId(id)
	roles, _ := models.RoleModel{}.FindAll()
	basicConfig, _ := configurations.GetBasicConfiguration()
	return ctx.Render("users/user/edit", fiber.Map{
		"Title": "Editar User",
		"Roles": roles,
		"User": user,
		"LoggedUser": authentication.GetLoggedUser(ctx),
		"BasicConfig": basicConfig,
	})
}


func (UserController) edit(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	roleId := ctx.FormValue("role_id")
	var userModel models.UserModel
	user, _ := userModel.FindByUniqueId(id)
	user.RoleId = conversion.StringToInt(roleId)
	user.UpdatedAt = time.Now()
	_, err := userModel.Update(user)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	loggerUser.Info("User '"+user.UserName+"' updated successfully", config.LogRequestPath(ctx))
	return ctx.Redirect("/users")
}


func (UserController) searchForm(ctx *fiber.Ctx) error {
	basicConfig, _ := configurations.GetBasicConfiguration()
	return ctx.Render("user/user/search", fiber.Map{
		"Title": "Search Users",
		"LoggedUser": authentication.GetLoggedUser(ctx),
		"BasicConfig": basicConfig,
	})
}


func (UserController) search(ctx *fiber.Ctx) error {
	param := ctx.FormValue("search_param")
	results, _ := models.UserModel{}.Search(param)
	count := len(results)
	loggedUser := authentication.GetLoggedUser(ctx)
	basicConfig, _ := configurations.GetBasicConfiguration()
	loggerUser.Info(fmt.Sprintf("User '%s' searched for '%v' and found %d results", loggedUser.UserName, param, count), config.LogRequestPath(ctx))
	return ctx.Render("user/user/search-results", fiber.Map{
		"Title": "Results",
		"Results": results,
		"Param": param,
		"Count": count,
		"LoggedUser": loggedUser,
		"BasicConfig": basicConfig,
	})
}


func (UserController) deactivateForm(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	user, _ := models.UserModel{}.GetDataByUniqueId(id)
	basicConfig, _ := configurations.GetBasicConfiguration()
	return ctx.Render("user/deactivate", fiber.Map{
		"Title": "Desactivar User",
		"User": user,
		"LoggedUser": authentication.GetLoggedUser(ctx),
		"BasicConfig": basicConfig,
	})
}


func (UserController) deactivate(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var userModel models.UserModel
	user, _ := userModel.FindByUniqueId(id)
	user.Active = "No"
	user.UpdatedAt = time.Now()
	userModel.Update(user)
	loggerUser.Info(fmt.Sprintf("User '%s' deactivated successfully!", user.UserName), config.LogRequestPath(ctx))
	return ctx.Redirect("/users/"+id+"/details")
}


func (UserController) activateForm(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	user, _ := models.UserModel{}.GetDataByUniqueId(id)
	basicConfig, _ := configurations.GetBasicConfiguration()
	return ctx.Render("user/user/activate", fiber.Map{
		"Title": "Activar User",
		"User": user,
		"LoggedUser": authentication.GetLoggedUser(ctx),
		"BasicConfig": basicConfig,
	})
}


func (UserController) activate(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var userModel models.UserModel
	user, _ := userModel.FindByUniqueId(id)
	user.Active = "Yes"
	user.UpdatedAt = time.Now()
	user.Token = encryption.GenerateRandomToken()
	_, err := userModel.Update(user)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	loggerUser.Info(fmt.Sprintf("User '%s' activated successfully!", user.UserName), config.LogRequestPath(ctx))
	return ctx.Redirect("/users/"+id+"/details")
}


func (UserController) addImageForm(ctx *fiber.Ctx) error {
	basicConfig, _ := configurations.GetBasicConfiguration()
	return ctx.Render("user/user/add-image", fiber.Map{
		"Title": "Add Image",
		"LoggedUser": authentication.GetLoggedUser(ctx),
		"BasicConfig": basicConfig,
	})
}


func (UserController) addImage(ctx *fiber.Ctx) error {
	userImage, _ := helpers.UploadFile(ctx, "user_image", "image", config.UploadImagePath())
	loggedUser := authentication.GetLoggedUser(ctx)
	var userModel models.UserModel
	user, _ := userModel.FindById(loggedUser.UserId)
	user.Image = userImage
	user.UpdatedAt = time.Now()
	_, err := userModel.Update(user)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	loggerUser.Info(fmt.Sprintf("User '%s' added image", user.UserName), config.LogRequestPath(ctx))
	return ctx.Redirect("/user-data")
}


func (UserController) changePasswordForm(ctx *fiber.Ctx) error {
	basicConfig, _ := configurations.GetBasicConfiguration()
	return ctx.Render("user/user/change-password", fiber.Map{
		"Title": "Updated Password",
		"LoggedUser": authentication.GetLoggedUser(ctx),
		"BasicConfig": basicConfig,
	})
}


func (UserController) changePassword(ctx *fiber.Ctx) error {
	password := ctx.FormValue("password")
	//passwordConf := ctx.FormValue("password_conf")
	loggedUser := authentication.GetLoggedUser(ctx)

	var userModel models.UserModel
	user, _ := userModel.FindById(loggedUser.UserId)
	user.Password = encryption.HashPassword(password)
	_, err := userModel.Update(user)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	loggerUser.Info(fmt.Sprintf("User '%s' updated password", user.UserName), config.LogRequestPath(ctx))
	return ctx.Redirect("/auth/login")
}
