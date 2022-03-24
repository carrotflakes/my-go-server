package route_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetNotes(t *testing.T) {
	server := NewServer()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/notes", nil)
	server.ServeHTTP(w, req)

	assert.Equal(t, w.Code, 200)
}
