package csv

import (
	"fileLoader/config"
	"strings"
)

type CsvLineProcessor struct {
	delimiter string
}

func NewCsvLineProcessor(delimiter string) *CsvLineProcessor {
	return &CsvLineProcessor{delimiter: delimiter}
}

func (c *CsvLineProcessor) ProcessLine(line string, configFields []*config.Field) interface{} {
	data := strings.Split(line, c.delimiter)
	itemMap := make(map[string]string)
	for _, field := range configFields {
		itemMap[field.FieldName] = data[field.Index-1]
	}
	return itemMap
}
