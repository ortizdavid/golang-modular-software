package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/ortizdavid/go-nopain/conversion"
)

func ConfigStaticFiles(app *fiber.App) {
	app.Static("/", "./public/static")
	app.Static("/uploads", "./public/uploads")
}

func GetTemplateEngine() *html.Engine {
	engine := html.New("./public/templates", ".html")
	return engine
}

func ListenAddr() string {
	return GetEnv("APP_HOST") +":"+ GetEnv("APP_PORT")
}

func ShutdownTimeout() int {
	return conversion.StringToInt(GetEnv("SHUTDOWN_TIMEOUT"))
}

//------- Uploads
func UploadImagePath() string {
	return GetEnv("UPLOAD_IMAGE_PATH")
}

func UploadDocumentPath() string {
	return GetEnv("UPLOAD_DOCUMENT_PATH")
}

func MaxUploadImageSize() int {
	return conversion.StringToInt(GetEnv("MAX_UPLOAD_IMAGE_SIZE"))
}

func MaxUploadDocumentSize() int {
	return conversion.StringToInt(GetEnv("MAX_UPLOAD_DOCUMENT_SIZE"))
}

// ---- Requests
func RequestsPerSecond() int {
	return conversion.StringToInt(GetEnv("REQUESTS_PER_SECONDS"))
}

func RequestsExpiration() int {
	return conversion.StringToInt(GetEnv("REQUESTS_EXPIRATION"))
}

