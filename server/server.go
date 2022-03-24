package server

import (
	"my-arch/server/middleware"
	"my-arch/server/route"
	"my-arch/usecase"

	"github.com/gin-gonic/gin"
)

func New(usecase *usecase.Usecase) *gin.Engine {
	r := gin.New()

	r.GET("/healthz", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"status": "ok",
		})
	})

	r.Use(middleware.Cors())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Authorize())

	route.Setup(r, usecase)

	return r
}
