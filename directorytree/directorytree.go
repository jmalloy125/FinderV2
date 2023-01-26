package directorytree

import (
	"io/fs"
	"log"
	"path/filepath"
)

type DirectoryTree struct {
	root  string
	files []FileData
}

func NewDirectoryTree(path string) DirectoryTree {
	return DirectoryTree{root: path}
}

func (directory *DirectoryTree) Walk() {
	err := filepath.WalkDir(directory.root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal("Failed to walk at path: ", path)
		}
		if !d.IsDir() {
			directory.files = append(directory.files, FileData{path: path})
		}
		return nil
	})
	if err != nil {
		log.Fatal("Failed to walk directory: ", directory.root)
	}
}

func (directory *DirectoryTree) GetFiles() []FileData {
	return directory.files
}
