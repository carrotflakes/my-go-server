package middleware

import (
	"my-arch/usecase"

	"github.com/gin-gonic/gin"
)

func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO

		c.Set(usecase.CtxUserID, 1)
		c.Next()
	}
}
