// アプリケーションレイヤーに渡すコンテキストを定義
package route

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Context gin.Context

func (c Context) SetCookie(name, value string, maxAge int, path, domain string, secure, httpOnly bool) {
	ginCtx := (gin.Context)(c)
	ginCtx.SetCookie(name, value, maxAge, path, domain, secure, httpOnly)
}

func (c Context) GetCookie(name string) (string, error) {
	ginCtx := (gin.Context)(c)
	return ginCtx.Cookie(name)
}

// Implement the context.Context interface

func (c Context) Deadline() (deadline time.Time, ok bool) {
	ginCtx := (gin.Context)(c)
	return ginCtx.Deadline()
}

func (c Context) Done() <-chan struct{} {
	ginCtx := (gin.Context)(c)
	return ginCtx.Done()
}

func (c Context) Err() error {
	ginCtx := (gin.Context)(c)
	return ginCtx.Err()
}

func (c Context) Value(key any) any {
	ginCtx := (gin.Context)(c)
	return ginCtx.Value(key)
}
