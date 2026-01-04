package analyzer_test

import (
	"dupe-finder/analyzer"
	imagereader "dupe-finder/image-reader"
	"testing"
)

func TestDrill(t *testing.T) {
	tests := []struct {
		name     string
		input    []imagereader.ImageInfo
		wantDirs int
		checkFn  func(t *testing.T, result map[string]*analyzer.DirInfo)
	}{
		{
			name:     "empty input",
			input:    []imagereader.ImageInfo{},
			wantDirs: 0,
		},
		{
			name: "one image - one dir",
			input: []imagereader.ImageInfo{
				imagereader.ImageInfo{Path: "/dir1/image1.png", HashCode: "hash1"},
			},
			wantDirs: 1,
		},
		{
			name: "two images in same dir - one dir",
			input: []imagereader.ImageInfo{
				imagereader.ImageInfo{Path: "/dir1/image1.png", HashCode: "hash1"},
				imagereader.ImageInfo{Path: "/dir1/image2.png", HashCode: "hash2"},
			},
			wantDirs: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := analyzer.Drill(tt.input)

			if len(result) != tt.wantDirs {
				t.Errorf("got %d directories, want %d", len(result), tt.wantDirs)
			}

			if tt.checkFn != nil {
				tt.checkFn(t, result)
			}
		})
	}

}
