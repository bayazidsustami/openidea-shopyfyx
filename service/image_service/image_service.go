package image_service

import (
	"context"
	"mime/multipart"
)

type ImageService interface {
	UploadImage(ctx context.Context, file multipart.File, fileName string) (string, error)
}
