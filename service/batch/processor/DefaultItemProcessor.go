package processor

import (
	"encoding/json"
	"fmt"
)

type DefaultItemProcessor struct{}

func (DefaultItemProcessor) Process(item interface{}) interface{} {
	jsonData, err := json.Marshal(item)
	if err != nil {
		fmt.Println("Error al convertir a JSON:", err)
	}
	return jsonData
}
