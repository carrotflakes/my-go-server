// アプリケーションレイヤー用のコンテキストを定義
// キーは公開し、外側のレイヤーからsetすることができるようにする
package usecase

import (
	"context"
	"fmt"
)

type Context interface {
	context.Context

	SetCookie(name, value string, maxAge int, path, domain string, secure, httpOnly bool)
	GetCookie(name string) (string, error)
}

const CtxUserID = "userID"

func getUserID(ctx Context) (int, error) {
	userID := ctx.Value(CtxUserID)
	if userID == nil {
		return 0, fmt.Errorf("userID is not set")
	}
	return userID.(int), nil
}
