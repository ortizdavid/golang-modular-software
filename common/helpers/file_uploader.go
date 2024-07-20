package helpers

import (
	"fmt"
	"time"
	"strconv"
	"strings"
	"path/filepath"
	"github.com/gofiber/fiber/v2"
)


// Upload a file
func UploadFile(ctx *fiber.Ctx, formFile string, fileType string, savePath string) (string, error) {
    file, err := ctx.FormFile(formFile)
    if err != nil {
        return "", ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
    }
    fileName := file.Filename
    fileExtension := getFileExtension(fileName)
    extensions := extensionsByFileType(fileType)
    if !isValidExtension(fileExtension, extensions) {
        return "", ctx.Status(fiber.StatusBadRequest).SendString("Invalid file extension. Only " +fileType)
    }
    newFileName := generateUniqueFileName(fileName)
    saveFilePath := filepath.Join(savePath, newFileName)
    err = ctx.SaveFile(file, saveFilePath)
    if err != nil {
        return "", ctx.Status(fiber.StatusInternalServerError).SendString("Error saving the file")
    }
    return newFileName, nil
}


func extensionsByFileType(fileType string) []string {
    var extensions []string
    switch fileType {
    case "document":
        extensions = []string{".pdf"}
    case "image":
        extensions = []string{".jpg", ".jpeg", ".png", ".gif"}
    }
    return extensions
}


func generateUniqueFileName(originalName string) string {
    extension := getFileExtension(originalName)
    timestamp := time.Now().Unix()
    random := strconv.FormatInt(time.Now().UnixNano(), 10)[10:]
    return fmt.Sprintf("%d%s%s", timestamp, random, extension)
}


func getFileExtension(originalFileName string) string {
    return strings.ToLower(filepath.Ext(originalFileName))
}


func isValidExtension(fileExtension string, allowedExtensions []string) bool {
    valid := false
    for _, ext := range allowedExtensions {
        if ext == fileExtension {
            valid = true
            break
        }
    }
    return valid
}