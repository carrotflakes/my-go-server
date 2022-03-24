package test

import (
	"my-arch/mydb"
	"my-arch/registry"
	"my-arch/server"
	"my-arch/usecase"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func NewServer() *gin.Engine {
	db := mydb.New()
	repos := registry.NewRepositories(db)
	usecase := usecase.New(repos)
	server := server.New(usecase)

	return server
}

func Test_Ping(t *testing.T) {
	server := NewServer()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	server.ServeHTTP(w, req)

	assert.Equal(t, w.Code, 200)
}
