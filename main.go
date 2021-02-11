package main

import (
	"net/http"

	"github.com/connectwithub/go-urlshortner/urlshortner"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/:path", func(c *gin.Context) {
		path := c.Param("path")
		resolvedPath, err := urlshortner.ResolvePath(path)
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
		}
		c.Redirect(http.StatusMovedPermanently, resolvedPath)
	})
	r.Run(":8080")
}
