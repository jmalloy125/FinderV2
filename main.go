package main

import (
	directorytree "finder/FinderV2/DirectoryTree"
	"fmt"
)

func main() {
	dirs := directorytree.NewDirectoryTree("./testDir/")
	dirs.Walk()
	dirs.ReadFilesToBytesAsync()
	for _, duplicate := range dirs.FindDuplicates() {
		if len(duplicate) > 1 {
			fmt.Printf("Found a duplicate set: %v", duplicate)
		}
	}
}
