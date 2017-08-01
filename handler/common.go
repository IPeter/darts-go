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
		x, _ := strconv.Atoi(c.Query("x"))
		y, _ := strconv.Atoi(c.Query("y"))
		cam1Img := c.Query("cam2img")
		cam2Img := c.Query("cam2img")


		c.Set("command", &model.CamCommand{Score: commandNum, Modifier: commandMod, X:x, Y:y, Cam1Img:cam1Img, Cam2Img:cam2Img })

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

func NoCache() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
		c.Header("Pragma", "no-cache")
		c.Header("Expires", "0")

		c.Next()
	}
}
