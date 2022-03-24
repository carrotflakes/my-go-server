package route_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PostNotes(t *testing.T) {
	server := NewServer()

	w := httptest.NewRecorder()
	body := strings.NewReader(`{"text":"test"}`)
	req, _ := http.NewRequest("POST", "/notes", body)
	server.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"note\":{\"id\":0,\"text\":\"test\",\"created_at\":\"2019-01-01T00:00:00Z\"}}", w.Body.String())
}
