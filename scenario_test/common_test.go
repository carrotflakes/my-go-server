package scenario_test

import (
	"bytes"
	"encoding/json"
	"my-arch/gateway"
	"my-arch/graph"
	"my-arch/mydb"
	"my-arch/pubsub"
	"my-arch/server"
	"my-arch/usecase"
	"net/http"
	"net/http/httptest"

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

func GetRequest(server *gin.Engine, path string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", path, nil)
	server.ServeHTTP(w, req)
	return w
}

func PostRequest(server *gin.Engine, path string, body map[string]any) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	j, _ := json.Marshal(body)
	reader := bytes.NewReader(j)
	req, _ := http.NewRequest("POST", path, reader)
	server.ServeHTTP(w, req)
	return w
}
