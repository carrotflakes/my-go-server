package route_test

import (
	"my-arch/mydb"
	"my-arch/registry"
	"my-arch/server"
	"my-arch/usecase"

	"github.com/gin-gonic/gin"
)

func NewServer() *gin.Engine {
	db := mydb.New()
	repos := registry.NewRepositories(db)
	usecase := usecase.New(repos)
	server := server.New(usecase)

	return server
}
