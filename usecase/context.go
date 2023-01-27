// アプリケーションレイヤー用のコンテキストを定義
// キーは公開し、外側のレイヤーからsetすることができるようにする
package usecase

import (
	"context"
	"fmt"
)

type SessionState struct {
	// -1 ならログインしていない
	UserID int
}

const CtxSessionState = "sessionState"

func getSessionState(ctx context.Context) (*SessionState, error) {
	sessionState := ctx.Value(CtxSessionState)
	if sessionState == nil {
		return nil, fmt.Errorf("sessionState is not set")
	}
	return sessionState.(*SessionState), nil
}
