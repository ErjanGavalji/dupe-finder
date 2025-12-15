package main

import (
	"crypto/md5"
	"encoding/hex"
	"flag"
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

type Dupe struct {
	info  ImageInfo
	dupes []ImageInfo
}

func isImage(path string) bool {
	validExts := []string{".png", ".jpg", ".jpeg", ".gif", ".bmp", ".webp"}
	ext := strings.ToLower(filepath.Ext(path))
	return slices.Contains(validExts, ext)
}

func readImages(rootDir string) (images []string, err error) {
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

func calculateHash(path string) (string, error) {
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

func computeHashes(imagePaths []string, compute func(path string) (string, error)) (infos []ImageInfo, err error) {
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

func findDupes(infos []string) []Dupe {
	var dupes []Dupe
	return dupes
}

func main() {
	var rootDir string
	flag.StringVar(&rootDir, "root-dir", ".", "Root directory. Defaults to .")
	flag.Parse()

	allImages, err := readImages(rootDir)

	if err != nil {
		fmt.Printf("Error walking the path %q: %v\n", rootDir, err)
		return
	}

	infos, err := computeHashes(allImages, calculateHash)

	for _, info := range infos {
		fmt.Printf("Found image under %q; HashCode: %q\n", info.Path, info.HashCode)
	}

}
