package main

import (
	"log"
	"net/http"

	"github.com/flosch/pongo2/v6"
	"github.com/gin-gonic/gin"
	"gitlab.com/go-box/pongo2gin/v6"
)

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	}
}

func createGinApp() *gin.Engine {
	app := gin.Default()
	app.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return app
}

func main() {
	app := createGinApp()

	app.Use(CORS())

	app.Use(gin.Recovery())

	app.HTMLRender = pongo2gin.Default()
	app.Static("/static", "./static")

	v1 := app.Group("/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})

		v1.GET("/ping2", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong2",
			})
		})
	}

	app.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html",
			pongo2.Context{
				"name": "hello pongo",
				// "posts": posts,
			},
		)
	})

	log.Fatal(app.Run(":8080"))
}
