package scenario_test

import (
	"bytes"
	"encoding/json"
	"my-arch/gateway"
	"my-arch/mydb"
	"my-arch/server"
	"my-arch/usecase"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func NewServer() *gin.Engine {
	db := mydb.New()
	repos := gateway.NewRepositories(db)
	usecase := usecase.New(repos)
	server := server.New(usecase)

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
