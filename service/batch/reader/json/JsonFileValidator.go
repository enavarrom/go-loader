package json

import "errors"

type JsonFileValidator struct{}

func (JsonFileValidator) ValidateFileExtension(extension string) error {
	if extension != ".JSON" {
		return errors.New("formato de archivo no soportado")
	}
	return nil
}
