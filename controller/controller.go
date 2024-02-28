package controller

import (
	"github.com/gin-gonic/gin"
)

func SendMail(c *gin.Context) {
	c.JSON(200, gin.H{
		"api": "Send Mail",
	})
}
