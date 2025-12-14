package main

import (
	"flag"
	"fmt"
	"io/fs"
	"path/filepath"
)

func main() {
	var rootDir string
	flag.StringVar(&rootDir, "root-dir", ".", "Root directory. Defaults to .")
	flag.Parse()
	err := filepath.Walk(rootDir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing %q: %v\n", path, err)
			return err
		}
		fmt.Printf("visited file or dir: %q\n", path)
		return nil
	})
	if err != nil {
		fmt.Printf("Error walking the path %q: %v\n", rootDir, err)
	}
	fmt.Println(rootDir)

}
