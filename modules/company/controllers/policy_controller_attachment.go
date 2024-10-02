package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/go-nopain/filemanager"
	"github.com/ortizdavid/golang-modular-software/common/config"
	"github.com/ortizdavid/golang-modular-software/modules/company/entities"
)

func (ctrl *PolicyController) addAttachmentForm(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.moduleFlagStatusService.LoadModuleFlagStatus(c.Context())
	policy, err := ctrl.service.GetByUniqueId(c.Context(), id)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	return c.Render("company/policy/add-attachment", fiber.Map{
		"Title":            "Add Attachment",
		"Policy":           policy,
		"LoggedUser":       loggedUser,
		"AppConfig":        ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": moduleFlagStatus,
	})
}

func (ctrl *PolicyController) addAttachment(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	attachment, err := ctrl.service.GetByUniqueId(c.Context(), id)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	var request entities.CreatePolicyAttachmentRequest
	if err := c.BodyParser(&request); err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	err = ctrl.attachmentService.Create(c.Context(), c, request)
	if err != nil {
		ctrl.errorLogger.Error(c, err.Error())
		return ctrl.HandleErrorsWeb(c, err)
	}
	ctrl.infoLogger.Info(c, fmt.Sprintf("User '%s' added policy attachment for policy '%s'", loggedUser.UserName, attachment.PolicyName))
	return c.Redirect("/company/policies/" + id + "/details")
}

func (ctrl *PolicyController) displayAttachment(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	attachment, err := ctrl.attachmentService.GetByUniqueId(c.Context(), id)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	documentPath := config.UploadDocumentPath() + "/company"
	ctrl.infoLogger.Info(c, fmt.Sprintf("User '%s' displayed policy attachment '%s'", loggedUser.UserName, attachment.AttachmentName))
	return c.SendFile(documentPath + "/" + attachment.FileName)
}

func (ctrl *PolicyController) deleteAttachmentForm(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	moduleFlagStatus, _ := ctrl.moduleFlagStatusService.LoadModuleFlagStatus(c.Context())
	attachment, err := ctrl.attachmentService.GetByUniqueId(c.Context(), id)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	policy, err := ctrl.service.GetPolicyById(c.Context(), attachment.PolicyId)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	return c.Render("company/policy/delete-attachment", fiber.Map{
		"Title":            "Delete Attachment",
		"Policy":           policy,
		"LoggedUser":       loggedUser,
		"AppConfig":        ctrl.configService.LoadAppConfigurations(c.Context()),
		"ModuleFlagStatus": moduleFlagStatus,
		"PolicyAttachment": attachment,
	})
}

func (ctrl *PolicyController) deleteAttachment(c *fiber.Ctx) error {
	id := c.Params("id")
	loggedUser, _ := ctrl.authService.GetLoggedUser(c.Context(), c)
	attachment, err := ctrl.attachmentService.GetByUniqueId(c.Context(), id)
	if err != nil {
		return ctrl.HandleErrorsWeb(c, err)
	}
	policy, err := ctrl.service.GetPolicyById(c.Context(), attachment.PolicyId)
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
	ctrl.infoLogger.Info(c, fmt.Sprintf("User '%s' deleted policy attachment '%s'", loggedUser.UserName, attachment.AttachmentName))
	return c.Redirect("/company/policies/" + policy.UniqueId + "/details")
}