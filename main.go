package main

import (
	directorytree "finder/FinderV2/DirectoryTree"
	"fmt"
)

func main() {
	dirs := directorytree.NewDirectoryTree("./testDir/")
	dirs.Walk()
	for _, file := range dirs.GetFiles() {
		fmt.Printf("%v\n", file.GetPath())
	}
}
