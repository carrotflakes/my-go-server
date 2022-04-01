package scenario_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SignupSignout(t *testing.T) {
	server := NewServer()

	w := PostRequest(server, "/signup", map[string]any{
		"name":     "hoge",
		"email":    "hoge@example.com",
		"password": "hogehoge",
	})
	assert.Equal(t, w.Code, 200)

	w = PostRequest(server, "/signout", map[string]any{})
	assert.Equal(t, w.Code, 200)

	w = PostRequest(server, "/signin", map[string]any{
		"email":    "hoge@example.com",
		"password": "hogehoge",
	})
	assert.Equal(t, w.Code, 200)

	w = PostRequest(server, "/signout", map[string]any{})
	assert.Equal(t, w.Code, 200)
}
