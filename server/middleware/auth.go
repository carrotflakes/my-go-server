package middleware

import (
	"my-arch/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		userIDStr, err := c.Cookie("userID")
		if err == http.ErrNoCookie {
			c.Next()
			return
		}
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		if userIDStr == "" {
			c.Next()
			return
		}

		userID, err := strconv.Atoi(userIDStr)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.Set(usecase.CtxUserID, userID)
		c.Next()
	}
}
