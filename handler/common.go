package handler

import (
	"darts-go/model"
	"strconv"

	"gopkg.in/gin-gonic/gin.v1"
)

func ParseCommand() gin.HandlerFunc {
	return func(c *gin.Context) {
		commandNum, _ := strconv.Atoi(c.Query("num"))
		commandMod, _ := strconv.Atoi(c.Query("modifier"))

		c.Set("command", &model.CamCommand{Score: commandNum, Modifier: commandMod})

		c.Next()
	}
}

func JsonRecover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.JSON(200, gin.H{
					"message": err,
					"success": false,
				})
			}
		}()

		c.Next()
	}
}
