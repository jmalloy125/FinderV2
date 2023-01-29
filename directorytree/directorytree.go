package directorytree

import (
	"finder/FinderV2/hashing"
	"io/fs"
	"log"
	"path/filepath"
	"sync"
)

type DirectoryTree struct {
	asyncFileGroup sync.WaitGroup
	root           string
	files          map[string]*FileData
	//files          []FileData

}

func NewDirectoryTree(path string) DirectoryTree {
	return DirectoryTree{root: path, files: make(map[string]*FileData)}
}

func (directory *DirectoryTree) GetFiles() map[string]*FileData {
	return directory.files
}

func (directory *DirectoryTree) ReadFilesToBytesAsync() {
	receiveDataChan := make(chan hashing.ByteData)
	for path, _ := range directory.files {
		directory.asyncFileGroup.Add(1)
		go hashing.GetBytesHashed(path, &directory.asyncFileGroup, receiveDataChan)
	}
	go func() {
		directory.asyncFileGroup.Wait()
		close(receiveDataChan)
	}()
	for data := range receiveDataChan {
		directory.files[data.Path].Data = data.Bytes
	}
}

func (directory *DirectoryTree) Walk() {
	err := filepath.WalkDir(directory.root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal("Failed to walk at path: ", path)
		}
		if !d.IsDir() {
			directory.files[path] = &FileData{Path: path}
		}
		return nil
	})
	if err != nil {
		log.Fatal("Failed to walk directory: ", directory.root)
	}
}

func (directory *DirectoryTree) FindDuplicates() map[string]int {
	duplicates := make(map[string]int)
	for _, file := range directory.files {
		if key, ok := duplicates[file.Path]; !ok {

		}
	}
	return duplicates
}

// todo fix this shit
