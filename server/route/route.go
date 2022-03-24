package route

import (
	"my-arch/usecase"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	method     string
	path       string
	handlerGen func(*usecase.Usecase) gin.HandlerFunc
}

var handlers = []Handler{}

func Setup(r *gin.Engine, usecase *usecase.Usecase) {
	r.LoadHTMLFiles("public/index.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{})
	})

	for _, handler := range handlers {
		r.Handle(handler.method, handler.path, handler.handlerGen(usecase))
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

	r.GET("/healthz", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"status": "ok",
		})
	})
}
