package main

import (
	"fileLoader/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	route := gin.Default()
	//route.SetTrustedProxies(nil)
	controller.NewUploadFileController().Configure(route)
	route.GET("/health/status", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"status": "UP",
		})
	})
	route.Run()
}
