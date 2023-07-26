package config

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

type Field struct {
	FieldName string `json:"fieldName"`
	Index     int    `json:"index"`
}

type FileConfig struct {
	LoaderName    string   `json:"loaderName"`
	Delimiter     string   `json:"delimiter"`
	FetchSize     int      `json:"fetchSize"`
	ChunkSize     int      `json:"chunkSize"`
	SkipFirstLine bool     `json:"skipFirstLine"`
	FormatFile    string   `json:"formatFile"`
	StreamName    string   `json:"streamName"`
	Fields        []*Field `json:"fields"`
}

var instance *FileConfig
var once sync.Once

func GetInstance() *FileConfig {
	once.Do(func() {
		loadConfig()
	})
	return instance
}

func loadConfig() {
	configPath := os.Getenv("LOADER_CONFIG_PATH")
	if configPath != "" {
		fileContent, err := os.ReadFile(configPath)
		if err != nil {
			fmt.Println("Error al leer el archivo de configuración:", err)
		}

		// Decodifica el contenido del archivo JSON en la estructura Config
		err = json.Unmarshal(fileContent, &instance)
		if err != nil {
			fmt.Println("Error al decodificar el archivo JSON:", err)
		}
	} else {
		fmt.Println("Error en la ruta del Archivo de configuración")
	}

}
