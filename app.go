package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"name":     "hedui",
			"age":      23,
			"describe": "fron Fu jian sheng",
			"hobby":    "simple is better",
		})
	})
	r.Run(":8085")
}
