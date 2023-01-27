package main

import (
	directorytree "finder/FinderV2/DirectoryTree"
)

func main() {
	dirs := directorytree.NewDirectoryTree("./testDir/")
	dirs.Walk()
	dirs.ReadFilesToBytesAsync()
}
