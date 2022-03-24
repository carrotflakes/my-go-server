package route

import (
	"my-arch/presenter"
	"my-arch/usecase"

	"github.com/gin-gonic/gin"
)

func init() {
	handlers = append(handlers,
		Handler{
			"GET",
			"/users",
			func(usecase *usecase.Usecase) gin.HandlerFunc {
				return func(ctx *gin.Context) {
					users, err := usecase.UserGetAll(ctx.Request.Context())
					if err != nil {
						ctx.JSON(500, gin.H{
							"message": err.Error(),
						})
						return
					}

					ctx.JSON(200, gin.H{
						"users": presenter.Users(users),
					})
				}
			},
		},
	)
}
