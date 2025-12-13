package main

import (
	"flag"
	"fmt"
)

func main() {
	var rootDir string
	flag.StringVar(&rootDir, "root-dir", ".", "Root directory. Defaults to .")
	flag.Parse()
	fmt.Println(rootDir)

}
