package helpers

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"time"
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/go-nopain/encryption"
	"github.com/ortizdavid/go-nopain/filemanager"
	"github.com/ortizdavid/go-nopain/serialization"
	"github.com/ortizdavid/go-nopain/units"
)

// Uploader manages file uploads.
type Uploader struct {
	DestinationPath   string          // DestinationPath specifies the path to upload files.
	AllowedExtensions FileExtensions  // AllowedExtensions specifies the allowed file extensions.
	MaxSize           int64           // MaxSize specifies the maximum file size in MB.
}

// UploadInfo stores the information of the uploaded file.
type UploadInfo struct {
	OriginalFileName string
	FinalName        string    // FinalName stores the final name of the uploaded file after processing or saving.
	Size             int64     // Size stores the size of the uploaded file in bytes.
	ContentType      string    // ContentType stores the MIME type of the uploaded file.
	Extension        string    // Extension stores the file extension of the uploaded file.
	UploadTime       time.Time // UploadTime stores the time when the file was uploaded.
}

// FileExtensions represents a list of file extensions.
type FileExtensions []string

// Common file extensions.
var (
	ExtImages    = FileExtensions{".jpg", ".jpeg", ".png", ".gif"}
	ExtAudios    = FileExtensions{".mp3", ".aac", ".wav", ".flac"}
	ExtDocuments = FileExtensions{".txt", ".pdf", ".docx", ".ppt", ".pptx", ".xls", ".xlsx"}
	ExtVideos    = FileExtensions{".mp4", ".mkv", ".avi", ".flv"}
	ExtArchives  = FileExtensions{".zip", ".rar", ".7z", ".tar", ".gz"}
)

// file i
var upFileInfo filemanager.FileInfo
var upFileManager filemanager.FileManager

// NewUploader creates a new file uploader.
func NewUploader(destinationPath string, maxSize int, allowedExts FileExtensions) Uploader {
	if destinationPath == "" {
		panic("destinationPath cannot be empty")
	}
	if !upFileInfo.ExistsDir(destinationPath) {
		upFileManager.CreateSingleFolder(destinationPath)
	}
	return Uploader{
		DestinationPath:   destinationPath,
		AllowedExtensions: allowedExts,
		MaxSize:           int64(maxSize) * int64(units.MEGA_BYTE),
	}
}

// UploadSingleFile uploads a single file using Fiber.
func (upl Uploader) UploadSingleFile(c *fiber.Ctx, formFile string) (UploadInfo, error) {
	fileHeader, err := c.FormFile(formFile)
	if err != nil {
		return UploadInfo{}, err
	}
	file, err := fileHeader.Open()
	if err != nil {
		return UploadInfo{}, err
	}
	defer file.Close()

	// Validate file size
	if fileHeader.Size > upl.MaxSize {
		return UploadInfo{}, fmt.Errorf("file too large. must contain %dMB", upl.MaxSize)
	}

	// Get file extension and validate
	fileExt := upl.getFileExtension(fileHeader.Filename)
	if !upl.isValidExtension(fileExt) {
		return UploadInfo{}, fmt.Errorf("invalid file extension. allowed: %v", upl.AllowedExtensions)
	}
	fileName := upl.generateUniqueFileName(fileHeader.Filename)

	// Save the file
	err = upl.saveUploadedFile(file, fileName)
	if err != nil {
		return UploadInfo{}, err
	}

	// Build UploadInfo
	return UploadInfo{
		OriginalFileName: fileHeader.Filename,
		FinalName:        fileName,
		Size:             fileHeader.Size,
		ContentType:      fileHeader.Header.Get("Content-Type"),
		Extension:        fileExt,
		UploadTime:       time.Now(),
	}, nil
}

// UploadMultipleFiles uploads multiple files using Fiber.
func (upl Uploader) UploadMultipleFiles(c *fiber.Ctx, formFile string) ([]UploadInfo, error) {
	form, err := c.MultipartForm()
	if err != nil {
		return nil, err
	}
	files := form.File[formFile]
	var uploadInfos []UploadInfo

	// Iterate over each file
	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			return nil, err
		}
		defer file.Close()

		// Validate file size
		if fileHeader.Size > upl.MaxSize {
			return nil, fmt.Errorf("file '%s' is too large. Must be smaller than %d MB", fileHeader.Filename, upl.MaxSize)
		}

		// Get file extension and validate
		fileExt := upl.getFileExtension(fileHeader.Filename)
		if !upl.isValidExtension(fileExt) {
			return nil, fmt.Errorf("file '%s' has an invalid extension. Allowed: %v", fileHeader.Filename, upl.AllowedExtensions)
		}

		// Generate unique filename
		fileName := upl.generateUniqueFileName(fileHeader.Filename)

		// Save the file
		err = upl.saveUploadedFile(file, fileName)
		if err != nil {
			return nil, err
		}

		// Build UploadInfo
		uploadInfo := UploadInfo{
			OriginalFileName: fileHeader.Filename,
			FinalName:        fileName,
			Size:             fileHeader.Size,
			ContentType:      fileHeader.Header.Get("Content-Type"),
			Extension:        fileExt,
			UploadTime:       time.Now(),
		}
		uploadInfos = append(uploadInfos, uploadInfo)
	}
	return uploadInfos, nil
}

func (upl Uploader) UploadSingleToEndpoint(url string, formFile string) ([]UploadInfo, error) {
	// Open the file
	file, err := os.Open(formFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Get file info
	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	// Create a new HTTP request
	req, err := http.NewRequest("POST", url, file)
	if err != nil {
		return nil, err
	}

	// Set headers
	req.Header.Set("Content-Type", "application/octet-stream")
	req.Header.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`, formFile, fileInfo.Name()))

	// Perform the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("upload failed with status code: %d", resp.StatusCode)
	}

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Unmarshal response body to get upload info
	var uploadInfos []UploadInfo
	err = serialization.UnserializeJson(body, &uploadInfos)
	if err != nil {
		return nil, err
	}
	return uploadInfos, nil
}

// saveUploadedFile saves a file to the specified path.
func (upl Uploader) saveUploadedFile(file multipart.File, fileName string) error {
	defer file.Close()

	destination := filepath.Join(upl.DestinationPath, fileName)
	dst, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	return err
}

// generateUniqueFileName generates a unique file name by appending a UUID to the original file name's extension.
func (upl Uploader) generateUniqueFileName(originalName string) string {
	return encryption.GenerateUUID() + filepath.Ext(originalName)
}

// getFileExtension retrieves the file extension from the original file name and converts it to lowercase.
func (upl Uploader) getFileExtension(originalFileName string) string {
	return strings.ToLower(filepath.Ext(originalFileName))
}

// isValidExtension checks if the provided file extension is valid based on the allowed extensions.
func (upl Uploader) isValidExtension(fileExt string) bool {
	return fileExt != "" && slices.Contains(upl.AllowedExtensions, fileExt)
}
