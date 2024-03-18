package utils

import (
	"net/url"
	"strings"

	"github.com/go-playground/validator/v10"
)

func MustValidImageUrl(fl validator.FieldLevel) bool {
	urlString := fl.Field().String()

	// Parse the URL
	u, err := url.Parse(urlString)
	if err != nil {
		return false
	}

	// Get the file extension
	parts := strings.Split(u.Path, ".")
	extension := parts[len(parts)-1]

	// Check if the extension is jpg or jpeg
	if extension != "jpg" && extension != "jpeg" {
		return false
	}

	return true
}
