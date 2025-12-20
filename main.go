package main

import (
	"crypto/md5"
	imagereader "dupe-finder/image-reader"
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

// First of all, we need to be able to have multiple directories as inputs, as
//
//	I have backups of my photos all around. Some exist at locationA and
//	locationB, while other duplicate directories exist at locations B and C and
//	D. A third set exist in all the locations.
type Dupe struct {
	info  imagereader.ImageInfo
	dupes []imagereader.ImageInfo
}

func isImage(path string) bool {
	validExts := []string{".png", ".jpg", ".jpeg", ".gif", ".bmp", ".webp"}
	ext := strings.ToLower(filepath.Ext(path))
	return slices.Contains(validExts, ext)
}

func readImages(rootDirs []string) (images []string, err error) {
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

func computeHashes(imagePaths []string, compute func(path string) (string, error)) (infos []imagereader.ImageInfo, err error) {
	imageInfos := make([]imagereader.ImageInfo, 0, len(imagePaths))
	for _, path := range imagePaths {
		hash, err := compute(path)
		if err != nil {
			return nil, err
		}
		newImageInfo := imagereader.ImageInfo{Path: path, HashCode: hash}
		imageInfos = append(imageInfos, newImageInfo)
	}

	return imageInfos, nil
}

func getDupeMap(infos []imagereader.ImageInfo) map[string][]imagereader.ImageInfo {
	var dupeMap map[string][]imagereader.ImageInfo = make(map[string][]imagereader.ImageInfo)

	for _, info := range infos {

		for _, ptDupe := range infos {
			if info.HashCode == ptDupe.HashCode && info.Path != ptDupe.Path {
				if dupeMap[info.HashCode] == nil {
					dupeMap[info.HashCode] = make([]imagereader.ImageInfo, 0)
				}
				dupeMap[info.HashCode] = append(dupeMap[info.HashCode], ptDupe)
			}
		}
	}

	return dupeMap
}

//func getDuplicateDirs(infos map[string][]imagereader.ImageInfo) map[string][]string {
//	dirs := make(map[string]string, 0)
//	for _, info := range infos {
//		infoDir := filepath.Dir(info.Path)
//		dirs[infoDir]
//	}
//	return make(map[string][]string)
//}

// Prints the ImageMap based on a specified analysis pattern.
// For example, there might be entire directories duplicated and we'd rather
// print them instead of every single image path. Additionally, we'd print the
// level of folder duplicacy if it is above a certain treshold.
func printMap(infos map[string][]imagereader.ImageInfo) {
}

func parseArgs() (rootDirs []string) {
	var rootDirsArg stringArrayFlags
	flag.Var(&rootDirsArg, "root-dirs", "Root directories, multiple instances. If none specified, reads the current directory only")
	flag.Parse()

	return rootDirsArg
}

func main() {
	rootDirs := parseArgs()
	if len(rootDirs) == 0 {
		rootDirs = append(rootDirs, ".")
	}

	allImages, err := readImages(rootDirs)

	if err != nil {
		fmt.Printf("Error walking the path %q: %v\n", rootDirs, err)
		return
	}

	infos, err := computeHashes(allImages, calculateHash)

	for _, info := range infos {
		fmt.Printf("Found image under %q; HashCode: %q\n", info.Path, info.HashCode)
	}

	var zeMap = getDupeMap(infos)
	fmt.Printf("There are %v duplicated items\n", len(zeMap))
	fmt.Println("Here they are:")
	fmt.Println("================================================================================")
	fmt.Printf("%v\n", zeMap)

}
