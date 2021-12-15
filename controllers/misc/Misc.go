package misc

import "github.com/gin-gonic/gin"

// GetPing godoc
// @Tags misc
// @Summary Get ping
// @Description Get ping

// @Success 200 {object} string
// @Failure 400 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /misc/ping [get]
func GetPing(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
