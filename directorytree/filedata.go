package directorytree

type FileData struct {
	Path string
	Data string
}

func (file *FileData) SetBytes(bytes string) {
	file.Data = bytes
}

func (file *FileData) GetPath() string {
	return file.Path
}
