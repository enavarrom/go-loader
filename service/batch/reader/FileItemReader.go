package reader

import (
	"bufio"
	"fileLoader/config"
	"fmt"
	"os"
	"path"
	"strings"
)

type ILineProcessor interface {
	ProcessLine(line string, configFields []*config.Field) interface{}
}

type IFileValidator interface {
	ValidateFileExtension(extension string) error
}

type FileItemReader struct {
	filePath      string
	file          *os.File
	scanner       *bufio.Scanner
	isFileOpen    bool
	fields        []*config.Field
	lineProcessor ILineProcessor
	fileValidator IFileValidator
}

func NewFileItemReader(fields []*config.Field, lineProcessor ILineProcessor, fileValidator IFileValidator) *FileItemReader {
	return &FileItemReader{fields: fields, lineProcessor: lineProcessor, fileValidator: fileValidator}
}

func (c *FileItemReader) Read() (interface{}, error) {
	if !c.isFileOpen {
		c.open()
		c.isFileOpen = true
	}

	var line = c.getNextLine()

	if line != "" {
		return c.lineProcessor.ProcessLine(c.scanner.Text(), c.fields), nil
	}

	if err := c.scanner.Err(); err != nil {
		return nil, err
	}

	defer c.file.Close()

	return nil, nil
}

func (c *FileItemReader) SetFilePath(filePath string) error {
	c.filePath = filePath
	extension := strings.ToUpper(path.Ext(filePath))
	return c.fileValidator.ValidateFileExtension(extension)
}

func (c *FileItemReader) open() error {
	file, err := os.Open(c.filePath)
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return err
	}
	c.file = file
	c.scanner = bufio.NewScanner(c.file)
	return nil
}

func (c *FileItemReader) getNextLine() string {
	if c.scanner.Scan() {
		return c.scanner.Text()
	}
	return ""
}
