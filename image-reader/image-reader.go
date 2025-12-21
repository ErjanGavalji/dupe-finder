package imagereader

import (
	"path/filepath"
	"slices"
	"strings"
)

type ImageInfo struct {
	Path     string
	HashCode string
}

func IsImage(path string) bool {
	validExts := []string{".png", ".jpg", ".jpeg", ".gif", ".bmp", ".webp"}
	ext := strings.ToLower(filepath.Ext(path))
	return slices.Contains(validExts, ext)
}
