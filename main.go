package main

import (
	directorytree "finder/FinderV2/DirectoryTree"
)

func main() {
	dirs := directorytree.NewDirectoryTree("./testDir/")
	dirs.Walk()
	//for _, file := range dirs.GetFiles() {
	//	fmt.Println(file.GetPath())
	//}
	dirs.ReadFilesToBytesAsync()
}
