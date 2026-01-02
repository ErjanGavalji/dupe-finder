package analyzer

import (
	imagereader "dupe-finder/image-reader"
	"path/filepath"
)

// TODO: This might be totally unnecessary, as we will have the path duplicated,
// as it is used as a key of the map

type DirInfo struct {
	Path       string
	ImageInfos []imagereader.ImageInfo
}

func (dirInfo *DirInfo) add(imageInfo imagereader.ImageInfo) {
	dirInfo.ImageInfos = append(dirInfo.ImageInfos, imageInfo)
}

func Drill(infos []imagereader.ImageInfo) map[string]*DirInfo {
	dirInfos := make(map[string]*DirInfo, 0)

	for _, info := range infos {
		dir := filepath.Dir(info.Path)
		if _, ok := dirInfos[dir]; !ok {
			dirInfos[dir] = &DirInfo{dir, make([]imagereader.ImageInfo, 0)}
		}
		dirInfos[dir].add(info)
	}

	return dirInfos
}
