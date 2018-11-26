package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	authorized := router.Group("/", gin.BasicAuth(gin.Accounts{
		"aaa": "bbb",
	}))

	authorized.GET("/me", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})
	 return router
}

func main() {
	router := setupRouter()

	router.Run(":8998")
}
