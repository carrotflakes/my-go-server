// time.Now() のモック化
package domain

import (
	"my-arch/config"
	"time"
)

// Use this instead of time.Now()
func TimeNow() time.Time {
	if config.Get().Env == "test" {
		return time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)
	}
	return time.Now()
}
