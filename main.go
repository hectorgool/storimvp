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
	router.GET("/sendmail", controller.SendMail)
	router.Run(":8080")
}
