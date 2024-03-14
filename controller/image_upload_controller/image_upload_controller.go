package image_upload_controller

import (
	"openidea-shopyfyx/service/auth_service"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

const (
	MaxUploadSize = 2 << 20  // 2 MB
	MinUploadSize = 10 << 10 // 10 KB
)

type ImageUploadController struct {
	AuthService auth_service.AuthService
}

func New(
	authService auth_service.AuthService,
) ImageUploadController {
	return ImageUploadController{
		AuthService: authService,
	}
}

func (controller *ImageUploadController) UploadImage(ctx *fiber.Ctx) error {
	form, err := ctx.MultipartForm()
	if err != nil {
		return err
	}

	files := form.File["file"]

	for _, file := range files {
		if file.Size > MaxUploadSize || file.Size < MinUploadSize {
			return fiber.NewError(fiber.StatusBadRequest, "image more than 2MB or less than 10KB")
		}

		ext := filepath.Ext(file.Filename)
		if ext != ".jpg" && ext != ".jpeg" {
			return fiber.NewError(fiber.StatusBadRequest, "not *.jpg | *.jpeg")
		}
	}

	return ctx.SendString("image uploaded successfully")
}
