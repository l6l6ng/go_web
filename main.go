package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func someJson(c *gin.Context) {
	data := map[string]interface{}{
		"lang": "  语言",
		"tag":  "  <br>",
	}
	c.JSON(http.StatusOK, data)
	c.AsciiJSON(http.StatusOK, data)
}

func index(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}

func main() {
	r := gin.Default()

	r.GET("/", index)
	r.GET("/someJson", someJson)

	r.Run(":8080")
}
