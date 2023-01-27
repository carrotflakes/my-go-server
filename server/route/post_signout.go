package route

import (
	"my-arch/usecase"

	"github.com/gin-gonic/gin"
)

func init() {
	handlers = append(handlers,
		Handler{
			"POST",
			"/signout",
			func(usecase *usecase.Usecase) gin.HandlerFunc {
				return func(ctx *gin.Context) {
					ctx.SetCookie("userID", "", 0, "/", "", false, false) // TODO: use sessionState

					ctx.JSON(200, gin.H{
						"status": "ok",
					})
				}
			},
		},
	)
}
