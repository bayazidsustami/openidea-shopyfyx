package image_upload_controller

import (
	"openidea-shopyfyx/service/auth_service"
	"openidea-shopyfyx/service/image_service"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

const (
	MaxUploadSize = 2 << 20  // 2 MB
	MinUploadSize = 10 << 10 // 10 KB
)

type ImageUploadController struct {
	AuthService  auth_service.AuthService
	ImageService image_service.ImageService
}

func New(
	authService auth_service.AuthService,
	imageService image_service.ImageService,
) ImageUploadController {
	return ImageUploadController{
		AuthService:  authService,
		ImageService: imageService,
	}
}

func (controller *ImageUploadController) UploadImage(ctx *fiber.Ctx) error {
	form, err := ctx.MultipartForm()
	if err != nil {
		return err
	}

	files := form.File["file"]

	if len(files) == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "no image found")
	}

	file := files[0]

	if file.Size > MaxUploadSize || file.Size < MinUploadSize {
		return fiber.NewError(fiber.StatusBadRequest, "image more than 2MB or less than 10KB")
	}

	ext := filepath.Ext(file.Filename)
	if ext != ".jpg" && ext != ".jpeg" {
		return fiber.NewError(fiber.StatusBadRequest, "not *.jpg | *.jpeg")
	}

	src, err := file.Open()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	url, err := controller.ImageService.UploadImage(ctx.UserContext(), src, file.Filename)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "something error")
	}

	return ctx.JSON(map[string]string{
		"imageUrl": url,
	})
}
