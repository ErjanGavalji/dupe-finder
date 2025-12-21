package imagereader

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

type ImageInfo struct {
	Path     string
	HashCode string
}

func isImage(path string) bool {
	validExts := []string{".png", ".jpg", ".jpeg", ".gif", ".bmp", ".webp"}
	ext := strings.ToLower(filepath.Ext(path))
	return slices.Contains(validExts, ext)
}

func ReadImages(rootDirs []string) (images []string, err error) {
	var allImagePaths []string
	for _, rootDir := range rootDirs {
		imagePaths, err := readDirImages(rootDir)
		if err != nil {
			fmt.Printf("Error reading images in dir %q\n", rootDir)
			return nil, err
		}
		allImagePaths = append(allImagePaths, imagePaths...)
	}
	return allImagePaths, nil
}

func readDirImages(rootDir string) (images []string, err error) {
	var allImages []string
	walkErr := filepath.Walk(rootDir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing %q: %v\n", path, err)
			return err
		}

		if info.IsDir() {
			return nil
		}
		if isImage(path) {
			allImages = append(allImages, path)
		}

		return nil
	})
	if walkErr != nil {
		return nil, walkErr
	}
	return allImages, nil
}

func CalculateHash(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}

func ComputeHashes(imagePaths []string, compute func(path string) (string, error)) (infos []ImageInfo, err error) {
	imageInfos := make([]ImageInfo, 0, len(imagePaths))
	for _, path := range imagePaths {
		hash, err := compute(path)
		if err != nil {
			return nil, err
		}
		newImageInfo := ImageInfo{Path: path, HashCode: hash}
		imageInfos = append(imageInfos, newImageInfo)
	}

	return imageInfos, nil
}
