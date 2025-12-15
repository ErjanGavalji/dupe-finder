package main

import (
	"flag"
	"fmt"
	"io/fs"
	"path/filepath"
	"slices"
	"strings"
)

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

func main() {
	var rootDir string
	flag.StringVar(&rootDir, "root-dir", ".", "Root directory. Defaults to .")
	flag.Parse()

	allImages, err := readImages(rootDir)

	if err != nil {
		fmt.Printf("Error walking the path %q: %v\n", rootDir, err)
		return
	}

	for _, imagePath := range allImages {
		fmt.Printf("Found imaage under %q\n", imagePath)
	}

	fmt.Println(rootDir)

}
