package server

import (
	"my-arch/server/middleware"
	"my-arch/server/route"
	"my-arch/usecase"

	"github.com/gin-gonic/gin"
)

func New(usecase *usecase.Usecase) *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Authorize())

	route.Setup(r, usecase)

	return r
}
