package contenttype

import (
	"errors"
	"mime"
	"path/filepath"
	"strings"
)

func GetContentType(path string) (string, error) {
	var ct string

	path = filepath.Base(path)
	ext := filepath.Ext(path)
	fn := strings.TrimSuffix(path, ext)

	if ext == "" || fn == "" {
		return "", errors.New("Invalid file path")
	}

	ct = mime.TypeByExtension(ext)

	if ct == "" {
		return "", errors.New("Could not get content type")
	}

	return ct, nil
}
