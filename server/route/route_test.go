package route_test

import (
	"my-arch/gateway"
	"my-arch/graph"
	"my-arch/mydb"
	"my-arch/pubsub"
	"my-arch/server"
	"my-arch/usecase"

	"github.com/gin-gonic/gin"
)

func NewServer() *gin.Engine {
	db := mydb.New()
	repos := gateway.NewRepositories(db)
	pubsub := pubsub.NewPubsub()
	usecase := usecase.New(repos)
	gqlResolver := graph.NewResolver(repos, pubsub, usecase)
	server := server.New(usecase, gqlResolver)

	return server
}
