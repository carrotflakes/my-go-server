package middleware

import (
	"my-arch/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionState := &usecase.SessionState{
			UserID: -1,
		}
		c.Set(usecase.CtxSessionState, sessionState)

		userIDStr, err := c.Cookie("userID")
		if err == http.ErrNoCookie {
			goto Next
		}
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		if userIDStr == "" {
			goto Next
		}

		{
			userID, err := strconv.Atoi(userIDStr)
			if err != nil {
				c.AbortWithError(http.StatusInternalServerError, err)
				return
			}
			sessionState.UserID = userID
		}

	Next:
		c.Next()

		// Update cookie
		if sessionState.UserID != -1 {
			c.SetCookie("userID", strconv.Itoa(sessionState.UserID), 60*60*24, "/", "", false, false)
		} else {
			c.SetCookie("userID", "", 0, "/", "", false, false)
		}
	}
}
