package analyzer_test

import (
	"dupe-finder/analyzer"
	imagereader "dupe-finder/image-reader"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestDirAnalysis(t *testing.T) {
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
				{Path: "/dir1/image1.png", HashCode: "hash1"},
			},
			wantDirs: 1,
		},
		{
			name: "two images in same dir - one dir",
			input: []imagereader.ImageInfo{
				{Path: "/dir1/image1.png", HashCode: "hash1"},
				{Path: "/dir1/image2.png", HashCode: "hash2"},
			},
			wantDirs: 1,
		},
		{
			name: "two images in different dirs - two dirs",
			input: []imagereader.ImageInfo{
				{Path: "/dir1/image1.png", HashCode: "hash1"},
				{Path: "/dir2/image2.png", HashCode: "hash2"},
			},
			wantDirs: 2,
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

func TestDupeIdentification(t *testing.T) {
	tests := []struct {
		name  string
		input []imagereader.ImageInfo
		// Note, for simplicity, for the time being we will give fake duplucates
		// here to avoid having preliminarily assigned imageInfos. Though that might
		// be the proper thing in the future.
		expectedOutput map[string]*analyzer.DirInfo
	}{
		{
			name:           "empty",
			input:          []imagereader.ImageInfo{},
			expectedOutput: make(map[string]*analyzer.DirInfo, 0),
		},
		{
			name: "single",
			input: []imagereader.ImageInfo{
				{Path: "d1/d1f1", HashCode: "Code1"},
			},
			expectedOutput: map[string]*analyzer.DirInfo{
				"d1": {
					Path:       "d1",
					ImageInfos: []imagereader.ImageInfo{{Path: "d1/d1f1", HashCode: "Code1"}},
				},
			},
		},
		{
			name: "two different files in dirs",
			input: []imagereader.ImageInfo{
				{Path: "d1/d1f1", HashCode: "Code1"},
				{Path: "d2/d2f1", HashCode: "Code2"},
			},
			expectedOutput: map[string]*analyzer.DirInfo{
				"d1": {
					Path:       "d1",
					ImageInfos: []imagereader.ImageInfo{{Path: "d1/d1f1", HashCode: "Code1"}},
				},
				"d2": {
					Path:       "d2",
					ImageInfos: []imagereader.ImageInfo{{Path: "d2/d2f1", HashCode: "Code2"}},
				},
			},
		},
		{
			name: "four different files in two dirs",
			input: []imagereader.ImageInfo{
				{Path: "d1/d1f1", HashCode: "Code11"},
				{Path: "d1/d1f2", HashCode: "Code12"},
				{Path: "d2/d2f1", HashCode: "Code21"},
				{Path: "d2/d2f2", HashCode: "Code22"},
			},
			expectedOutput: map[string]*analyzer.DirInfo{
				"d1": {
					Path: "d1",
					ImageInfos: []imagereader.ImageInfo{
						{Path: "d1/d1f1", HashCode: "Code11"},
						{Path: "d1/d1f2", HashCode: "Code12"},
					},
				},
				"d2": {
					Path: "d2",
					ImageInfos: []imagereader.ImageInfo{
						{Path: "d2/d2f1", HashCode: "Code21"},
						{Path: "d2/d2f2", HashCode: "Code22"},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		result := analyzer.Drill(tt.input)

		if len(result) != len(tt.expectedOutput) {
			t.Errorf("result output length does not match the one of the expected output")
		}

		for key := range tt.expectedOutput {
			if _, exists := result[key]; !exists {
				t.Errorf("expected key %q not found in result", key)
			}
		}

		for key := range result {
			if _, exists := tt.expectedOutput[key]; !exists {
				t.Errorf("unexpected key %q found in result", key)
			}
		}

		opts := []cmp.Option{
			cmpopts.SortSlices(func(a, b imagereader.ImageInfo) bool {
				if a.Path != b.Path {
					return a.Path < b.Path
				}
				return a.HashCode < b.HashCode
			}),
		}

		if diff := cmp.Diff(tt.expectedOutput, result, opts...); diff != "" {
			t.Errorf("Drill() mismatch (-want +got):\n%s", diff)
		}
	}
}
