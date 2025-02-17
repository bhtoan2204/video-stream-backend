package initialize

import (
	"net"

	"github.com/bhtoan2204/user/global"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	MainGroup := r.Group("/api/v1")
	{
		MainGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "OK",
			})
		})
		MainGroup.GET("/users", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "OK",
				"port":    global.Listener.Addr().(*net.TCPAddr).Port,
			})
		})
	}
	return r
}
