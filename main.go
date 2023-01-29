package main

import (
	directorytree "finder/FinderV2/DirectoryTree"
	"fmt"
)

func main() {
	dirs := directorytree.NewDirectoryTree("./testDir/")
	dirs.Walk()
	dirs.ReadFilesToBytesAsync()
	for key, duplicate := range dirs.FindDuplicates() {
		fmt.Printf("%v: %v\n", key, duplicate)
	}
}
