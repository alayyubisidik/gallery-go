package helper

import (
	"gallery_go/exception"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"fmt"

	"github.com/gin-gonic/gin"
)

func ValidateImageFile(ctx *gin.Context) (*multipart.FileHeader, error) {
	fileHeader, err := ctx.FormFile("image")
	if err != nil {
		return nil, exception.NewBadRequestError("Image is requiredss")
	}

	const maxFileSize = 5 * 1024 * 1024 
	if fileHeader.Size > maxFileSize {
		return nil, exception.NewBadRequestError("Image file is too large")
	}

	allowedExtensions := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
	}
	ext := filepath.Ext(fileHeader.Filename)
	if !allowedExtensions[ext] {
		return nil, exception.NewBadRequestError("Invalid file extension")
	}

	return fileHeader, nil
}

// SaveImage saves the uploaded image file
func SaveImage(ctx *gin.Context, fileHeader *multipart.FileHeader) (string, error) {
	ext := filepath.Ext(fileHeader.Filename)
	currentTime := time.Now().Unix()
	newFileName := fmt.Sprintf("%d%s", currentTime, ext)

	err := ctx.SaveUploadedFile(fileHeader, fmt.Sprintf("./public/images/%s", newFileName))
	if err != nil {
		return "", err
	}

	return newFileName, nil
}

func DeleteImage(fileName string) error {
    filePath := fmt.Sprintf("./public/images/%s", fileName)
    if err := os.Remove(filePath); err != nil {
        return err
    }
    return nil
}
