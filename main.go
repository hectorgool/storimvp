package main

import (
	"storimvp/controller"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"api": "Hola Stori",
		})
	})
	router.GET("/sendmail/:userEmail", controller.SendMail)
	router.DELETE("/reset", controller.Reset)
	router.Run(":8080")
}
