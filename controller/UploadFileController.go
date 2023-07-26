package controller

import (
	"fileLoader/config"
	"fileLoader/service/batch/job"
	"fileLoader/service/batch/processor"
	"fileLoader/service/batch/reader"
	"fileLoader/service/batch/writer"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
)

type UploadFileController struct {
}

func NewUploadFileController() *UploadFileController {
	return &UploadFileController{}
}

func (c *UploadFileController) Configure(engine *gin.Engine) {
	engine.POST("/upload", c.uploadFile)
}

func (c *UploadFileController) uploadFile(context *gin.Context) {
	// Verificar que el método de la solicitud sea POST
	if context.Request.Method != http.MethodPost {
		context.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Método no permitido"})
		return
	}

	// Obtener el archivo de la solicitud
	file, err := context.FormFile("file")
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el archivo"})
		return
	}

	// Crear un nuevo archivo en el servidor para guardar el archivo subido
	uploadFolder := os.Getenv("LOADER_UPLOAD_FOLDER")
	if uploadFolder == "" {
		uploadFolder, _ = os.UserHomeDir()
	}

	serverFilePath := filepath.Join(uploadFolder, file.Filename)
	nuevoArchivo, err := os.Create(serverFilePath)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear el archivo en el servidor"})
		return
	}
	defer nuevoArchivo.Close()

	// Copiar el contenido del archivo subido al nuevo archivo en el servidor
	if err := context.SaveUploadedFile(file, serverFilePath); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar el archivo en el servidor"})
		return
	}

	if err := c.launchLoader(serverFilePath); err != nil {
		context.JSON(http.StatusUnsupportedMediaType, gin.H{"error": "Formato de Archivo no soportado"})
		return
	}

	if err := os.Remove(serverFilePath); err != nil {
		fmt.Printf("Error al borrar el archivo: %v\n", err)
		return
	}

	// Respuesta exitosa
	context.JSON(http.StatusOK, gin.H{"message": "Archivo subido exitosamente"})
}

func (c *UploadFileController) launchLoader(pathFile string) error {
	itemWriter := writer.GetItemWriter()
	itemReader := reader.GetItemReader()
	itemProcessor := processor.DefaultItemProcessor{}
	loader := job.NewLoader(config.GetInstance().LoaderName, itemReader, itemProcessor, itemWriter)
	loader.SetChunkSize(config.GetInstance().ChunkSize)
	loader.SetFetchSize(config.GetInstance().FetchSize)
	err := loader.SetFilePath(pathFile)
	if err != nil {
		return err
	}
	loader.Launch()
	return nil
}
