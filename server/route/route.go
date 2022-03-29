package route

import (
	"my-arch/usecase"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	method      string
	path        string
	makeHandler func(*usecase.Usecase) gin.HandlerFunc
}

var handlers = []Handler{}

func Setup(r *gin.Engine, usecase *usecase.Usecase) {
	// r.LoadHTMLFiles("public/index.html")
	// r.GET("/", func(c *gin.Context) {
	// 	c.HTML(200, "index.html", gin.H{})
	// })
	r.StaticFile("/", "public/index.html")

	for _, handler := range handlers {
		r.Handle(handler.method, handler.path, handler.makeHandler(usecase))
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/me", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"userId": ctx.Request.Context().Value("userID").(int),
		})
	})
}
