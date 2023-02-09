package route_test

import (
	"my-arch/gateway"
	"my-arch/graph"
	"my-arch/mydb"
	"my-arch/server"
	"my-arch/usecase"

	"github.com/gin-gonic/gin"
)

func NewServer() *gin.Engine {
	db := mydb.New()
	repos := gateway.NewRepositories(db)
	usecase := usecase.New(repos)
	gqlResolver := graph.NewResolver(repos, usecase)
	server := server.New(usecase, gqlResolver)

	return server
}
