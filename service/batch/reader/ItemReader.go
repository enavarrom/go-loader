package reader

type ItemReader interface {
	Read() (interface{}, error)
	SetFilePath(filePath string) error
}
