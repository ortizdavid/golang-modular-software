package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/go-nopain/filemanager"
	"github.com/ortizdavid/golang-modular-software/common/config"
	"github.com/ortizdavid/golang-modular-software/modules/company/entities"
)

func (ctrl *ProjectController) addAttachmentForm(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.moduleFlagStatusService.LoadModuleFlagStatus(c.Context())
	project, err := ctrl.service.GetByUniqueId(c.Context(), id)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	return c.Render("company/project/add-attachment", fiber.Map{
		"Title":            "Add Attachment",
		"Project":           project,
		"LoggedUser":       loggedUser,
		"AppConfig":        ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": moduleFlagStatus,
	})
}

func (ctrl *ProjectController) addAttachment(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	attachment, err := ctrl.service.GetByUniqueId(c.Context(), id)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	var request entities.CreateProjectAttachmentRequest
	if err := c.BodyParser(&request); err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	err = ctrl.attachmentService.Create(c.Context(), c, request)
	if err != nil {
		ctrl.errorLogger.Error(c, err.Error())
		return ctrl.HandleErrorsWeb(c, err)
	}
	ctrl.infoLogger.Info(c, fmt.Sprintf("User '%s' added project attachment for project '%s'", loggedUser.UserName, attachment.ProjectName))
	return c.Redirect("/company/projects/" + id + "/details")
}

func (ctrl *ProjectController) displayAttachment(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	attachment, err := ctrl.attachmentService.GetByUniqueId(c.Context(), id)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	documentPath := config.UploadDocumentPath() + "/company"
	ctrl.infoLogger.Info(c, fmt.Sprintf("User '%s' displayed project attachment '%s'", loggedUser.UserName, attachment.AttachmentName))
	return c.SendFile(documentPath + "/" + attachment.FileName)
}

func (ctrl *ProjectController) deleteAttachmentForm(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.moduleFlagStatusService.LoadModuleFlagStatus(c.Context())
	attachment, err := ctrl.attachmentService.GetByUniqueId(c.Context(), id)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	project, err := ctrl.service.GetById(c.Context(), attachment.ProjectId)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	return c.Render("company/project/delete-attachment", fiber.Map{
		"Title":            "Delete Attachment",
		"Project":           project,
		"LoggedUser":       loggedUser,
		"AppConfig":        ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": moduleFlagStatus,
		"ProjectAttachment": attachment,
	})
}

func (ctrl *ProjectController) deleteAttachment(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	attachment, err := ctrl.attachmentService.GetByUniqueId(c.Context(), id)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	project, err := ctrl.service.GetById(c.Context(), attachment.ProjectId)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	err = ctrl.attachmentService.Remove(c.Context(), attachment.UniqueId)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	documentPath := config.UploadDocumentPath() + "/company"
	var fileManager filemanager.FileManager
	fileManager.RemoveFile(documentPath,attachment.FileName)
	ctrl.infoLogger.Info(c, fmt.Sprintf("User '%s' deleted project attachment '%s'", loggedUser.UserName, attachment.AttachmentName))
	return c.Redirect("/company/projects/" + project.UniqueId + "/details")
}