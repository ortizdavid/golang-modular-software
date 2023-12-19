package config

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/ortizdavid/go-nopain/conversion"
)

func ConfigStaticFiles(app *fiber.App) {
	app.Static("/", "./static")
}

func GetTemplateEngine() *html.Engine {
	engine := html.New("./templates", ".html")
	return engine
}

func ListenAddr() string {
	LoadDotEnv()
	return os.Getenv("APP_HOST")+":"+os.Getenv("APP_PORT")
}

func ShutdownTimeout() int {
	LoadDotEnv()
	return conversion.StringToInt(os.Getenv("SHUTDOWN_TIMEOUT"))
}

//------- Uploads
func UploadImagePath() string {
	LoadDotEnv()
	return os.Getenv("UPLOAD_IMAGE_PATH")
}

func UploadDocumentPath() string {
	LoadDotEnv()
	return os.Getenv("UPLOAD_DOCUMENT_PATH")
}

// ---- Requests
func RequestsPerSecond() int {
	LoadDotEnv()
	return conversion.StringToInt(os.Getenv("REQUESTS_PER_SECONDS"))
}

func RequestsExpiration() int {
	LoadDotEnv()
	return conversion.StringToInt(os.Getenv("REQUESTS_EXPIRATION"))
}

