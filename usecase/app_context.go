// コンテキストの値の読み書きを行う関数を定義
// キーは公開し、外側のレイヤーからsetすることができるようにする
package usecase

import (
	"context"
	"fmt"
)

const CtxUserID = "userID"

func getUserID(ctx context.Context) (int, error) {
	userID := ctx.Value(CtxUserID)
	if userID == nil {
		return 0, fmt.Errorf("userID is not set")
	}
	return userID.(int), nil
}
