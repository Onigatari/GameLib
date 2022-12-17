package server

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func StartServer() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.Static("/assets", "./assets")

	router.GET("/", MainPage)

	log.Println("HTTP server is running! Connect - http://localhost:8080")
	router.Run(":8080")
}

func MainPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}
