package main

import (
	"flag"
	"fmt"
	"io/fs"
	"path/filepath"
	"slices"
)

func isImage(path string) bool {
	validExts := []string{".png", ".jpg", ".jpeg", ".gif", ".bmp", ".webp"}
	ext := filepath.Ext(path)
	return slices.Contains(validExts, ext)
}

func main() {
	var rootDir string
	flag.StringVar(&rootDir, "root-dir", ".", "Root directory. Defaults to .")
	flag.Parse()
	err := filepath.Walk(rootDir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing %q: %v\n", path, err)
			return err
		}

		if info.IsDir() {
			return nil
		}
		if isImage(path) {
			fmt.Printf("image found: %s\n", path)
		}

		return nil
	})
	if err != nil {
		fmt.Printf("Error walking the path %q: %v\n", rootDir, err)
	}
	fmt.Println(rootDir)

}
