package server

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.Static("/assets", "./assets")

	router.GET("/", MainPage)
	router.GET("/wheel", WhellPage)

	log.Println("HTTP server is running! Connect - http://localhost:8080")
	router.Run(":8080")
}

func MainPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func WhellPage(c *gin.Context) {
	c.HTML(http.StatusOK, "wheel.html", gin.H{})
}
