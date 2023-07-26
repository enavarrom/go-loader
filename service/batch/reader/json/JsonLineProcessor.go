package json

import (
	"encoding/json"
	"fileLoader/config"
	"fmt"
)

type JsonLineProcessor struct{}

func (JsonLineProcessor) ProcessLine(line string, configFields []*config.Field) interface{} {
	data := make(map[string]interface{})
	filteredData := make(map[string]interface{})

	err := json.Unmarshal([]byte(line), &data)
	if err != nil {
		fmt.Println("Error al convertir JSON a mapa:", err)
		return nil
	}

	for key, value := range data {
		for _, field := range configFields {
			if key == field.FieldName {
				filteredData[key] = value
			}
		}
	}

	return filteredData
}
