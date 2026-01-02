package main

import (
	"dupe-finder/analyzer"
	imagereader "dupe-finder/image-reader"
	"flag"
	"fmt"
	"os"
)

// First of all, we need to be able to have multiple directories as inputs, as I
// have backups of my photos all around. Some exist at locationA and locationB,
// while other duplicate directories exist at locations B and C and D. A third
// set exist in all the locations.
type Dupe struct {
	info  imagereader.ImageInfo
	dupes []imagereader.ImageInfo
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

// Prints the ImageMap based on a specified analysis pattern. For example, there
// might be entire directories duplicated and we'd rather print them instead of
// every single image path. Additionally, we'd print the level of folder
// duplicacy if it is above a certain treshold.
func printMap(infos map[string][]imagereader.ImageInfo) {
}

func parseArgs() (rootDirs []string) {
	var rootDirsArg stringArrayFlags
	flag.Var(&rootDirsArg, "root-dirs", "Root directories, multiple instances. If none specified, reads the current directory only")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "\n\nAnalyzes the provided directories for duplicate files and folders.\n")
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	return rootDirsArg
}

func main() {
	rootDirs := parseArgs()
	if len(rootDirs) == 0 {
		rootDirs = append(rootDirs, ".")
	}

	infos, err := imagereader.ReadImages(rootDirs)
	if err != nil {
		fmt.Printf("Error walking the dirs %q: %v\n", rootDirs, err)
		return
	}

	for _, info := range infos {
		fmt.Printf("Found image under %q; HashCode: %q\n", info.Path, info.HashCode)
	}

	dupes := analyzer.Drill(infos)
	fmt.Printf("\n\nDupes: %v\n\n\n", dupes)

	var zeMap = getDupeMap(infos)
	fmt.Printf("There are %v duplicated items\n", len(zeMap))
	fmt.Println("Here they are:")
	fmt.Println("================================================================================")
	fmt.Printf("%v\n", zeMap)

}
