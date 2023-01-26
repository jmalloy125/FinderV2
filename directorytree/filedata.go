package directorytree

type FileData struct {
	path string
	data []byte
}

func (file *FileData) ReadBytes(filepath string) {
	return
}

func (file *FileData) GetPath() string {
	return file.path
}
