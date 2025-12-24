package analyzer

import (
	imagereader "dupe-finder/image-reader"
	"fmt"
)

// TODO: This might be totally unnecessary, as we will have the path
//  duplicated, as it is used as a key of the map

type DirInfo struct {
	Path       string
	ImageInfos []imagereader.ImageInfo
}

func Drill(infos []imagereader.ImageInfo) map[string]DirInfo {
	dirInfos := make(map[string]DirInfo, 0)

	fmt.Printf("Heeeey, this is someting from the analyzer package")
	return dirInfos
}
