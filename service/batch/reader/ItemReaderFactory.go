package reader

import (
	"fileLoader/config"
	"fileLoader/service/batch/reader/csv"
	"fileLoader/service/batch/reader/json"
)

func GetItemReader() ItemReader {
	switch config.GetInstance().FormatFile {
	case "CSV":
		return NewFileItemReader(config.GetInstance().Fields, csv.NewCsvLineProcessor(config.GetInstance().Delimiter), csv.CsvFileValidator{})
	case "JSON":
		return NewFileItemReader(config.GetInstance().Fields, json.JsonLineProcessor{}, json.JsonFileValidator{})
	default:
		return nil
	}
}
