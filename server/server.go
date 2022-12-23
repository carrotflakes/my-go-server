package server

import (
	"my-arch/graph"
	"my-arch/server/middleware"
	"my-arch/server/route"
	"my-arch/usecase"

	gqlhandler "github.com/99designs/gqlgen/graphql/handler"
	gqlplayground "github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

func New(usecase *usecase.Usecase, gqlResolver *graph.Resolver) *gin.Engine {
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

	gqlSrv := gqlhandler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: gqlResolver}))
	playgroundSrv := gqlplayground.Handler("GraphQL playground", "/gql/query")
	r.GET("/gql/", func(c *gin.Context) {
		playgroundSrv.ServeHTTP(c.Writer, c.Request)
	})
	r.POST("/gql/query", func(c *gin.Context) {
		gqlSrv.ServeHTTP(c.Writer, c.Request)
	})

	return r
}
