package router

import (
	"chat/lib"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Use(lib.PanicJson{}.Handle)
	r.GET("/t", func(c *gin.Context) {
		panic(lib.PanicJson{10, "xxx", map[string]interface{}{"count": 4, "sw": "dddd"}})
		c.JSON(200, gin.H{
			"message": "pong",
		})
		return
	})
	r.GET("/t1", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/t2", func(c *gin.Context) {
		num1 := 10
		num2 := 0
		num3 := num1 / num2
		c.JSON(200, gin.H{
			"message": num3,
		})
	})
	return r
}
