package csv

import "errors"

type CsvFileValidator struct{}

func (CsvFileValidator) ValidateFileExtension(extension string) error {
	if !(extension == ".CSV" || extension == ".TXT") {
		return errors.New("formato de archivo no soportado")
	}
	return nil
}
